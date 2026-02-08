package main

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

type account struct {
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type rankedEntry struct {
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
}

type PlayerStats struct {
	Name    string `json:"name"`
	Peak    int    `json:"peak"`
	Gained  int    `json:"gained"`
	Current int    `json:"current"`
	Wins    int    `json:"wins"`
	Losses  int    `json:"losses"`
	Played  int    `json:"played"`
	Tier    string `json:"tier"`
	Rank    string `json:"rank"`
	LP      int    `json:"lp"`
	Date    string `json:"date"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isSoloQueue(entry rankedEntry) bool {
	return entry.QueueType == "RANKED_SOLO_5x5"
}

func toPlayerStats(entry rankedEntry, name string) PlayerStats {
	return PlayerStats{
		Name:   name,
		Tier:   entry.Tier,
		Rank:   entry.Rank,
		Wins:   entry.Wins,
		Losses: entry.Losses,
		Played: entry.Wins + entry.Losses,
		LP:     entry.LeaguePoints,
		Date:   time.Now().Format("2006-01-02"),
	}
}

func main() {
	apiKey := os.Getenv("RIOT_API_KEY")
	baseUrl := "https://europe.api.riotgames.com/"
	serverBaseUrl := "https://euw1.api.riotgames.com/"
	accountsPath := "riot/account/v1/accounts/by-riot-id/"
	entriesPath := "lol/league/v4/entries/by-puuid/"
	apiKeyParam := "?api_key=" + apiKey

	resp, err := http.Get(baseUrl + accountsPath + "Impala/KAZ" + apiKeyParam)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var acc account
	err = json.Unmarshal(body, &acc)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", acc)

	log.Printf("PUUID LIT NASTY %s", acc.PUUID)

	resp, err = http.Get(serverBaseUrl + entriesPath + acc.PUUID + apiKeyParam)
	if err != nil {
		log.Fatalln(err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("BODY: %s", string(body))

	var rankedEntries []rankedEntry
	err = json.Unmarshal(body, &rankedEntries)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", rankedEntries)

	soloQueue := rankedEntry{}

	for i := range rankedEntries {
		if isSoloQueue(rankedEntries[i]) {
			soloQueue = rankedEntries[i]
		}
	}

	log.Printf("SOLO QUEUE %+v\n", soloQueue)

	playerStats := toPlayerStats(soloQueue, acc.GameName)

	log.Printf("PLAYERSTATS %+v\n", playerStats)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",

			"https://rankedrace.win",
			"https://www.rankedrace.win",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true, // only if using cookies
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/rank", func(c *gin.Context) {
		c.JSON(200, playerStats)
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
