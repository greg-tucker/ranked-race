package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
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
	Name        string `json:"name"`
	Peak        int    `json:"peak"`
	Gained      int    `json:"gained"`
	Current     int    `json:"current"`
	Wins        int    `json:"wins"`
	Losses      int    `json:"losses"`
	Played      int    `json:"played"`
	Tier        string `json:"tier"`
	Rank        string `json:"rank"`
	LP          int    `json:"lp"`
	Date        string `json:"date"`
	WinRate     string `json:"winRate"`
	DisplayRank string `json:"displayRank"`
	Tag         string `json:"tag"`
}

type InputPlayer struct {
	Name string
	Tag  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isSoloQueue(entry rankedEntry) bool {
	return entry.QueueType == "RANKED_SOLO_5x5"
}

func toPlayerStats(entry rankedEntry, name string, tag string) PlayerStats {
	return PlayerStats{
		Name:        name,
		Tier:        entry.Tier,
		Rank:        entry.Rank,
		Wins:        entry.Wins,
		Losses:      entry.Losses,
		Played:      entry.Wins + entry.Losses,
		LP:          entry.LeaguePoints,
		Current:     calculateLp(entry),
		Date:        time.Now().Format("2006-01-02"),
		WinRate:     fmt.Sprint(math.Round(10000*float64(entry.Wins)/(float64(entry.Wins+entry.Losses)))/100) + "%",
		DisplayRank: entry.Tier + " " + entry.Rank + " " + fmt.Sprint(entry.LeaguePoints) + " LP",
		Tag:         tag,
	}
}

func calculateLp(entry rankedEntry) int {
	metalIndex := metals[entry.Tier]
	return metalIndex*400 + ranksMap[entry.Rank]*100 + entry.LeaguePoints
}

func getPlayerStats(inputPlayer InputPlayer) PlayerStats {
	resp, err := http.Get(baseUrl + accountsPath + inputPlayer.Name + "/" + inputPlayer.Tag + apiKeyParam)
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

	//We Read the response body on the line below.
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

	playerStats := toPlayerStats(soloQueue, acc.GameName, inputPlayer.Tag)

	log.Printf("PLAYERSTATS %+v\n", playerStats)

	return playerStats
}

var apiKey = os.Getenv("RIOT_API_KEY")
var baseUrl = "https://europe.api.riotgames.com/"
var serverBaseUrl = "https://euw1.api.riotgames.com/"
var accountsPath = "riot/account/v1/accounts/by-riot-id/"
var entriesPath = "lol/league/v4/entries/by-puuid/"
var apiKeyParam = "?api_key=" + apiKey
var users = []InputPlayer{
	{Name: "Impala", Tag: "KAZ"},
	{Name: "Bjerkingfan", Tag: "EUW"},
	{Name: "ctrl alt cute", Tag: "xoxo"},
	{Name: "oystericetea", Tag: "EUW"},
	{Name: "Pissinglnthewind", Tag: "EUW"},
	{Name: "mayalover3", Tag: "EUW"},
	{Name: "SnazzyG", Tag: "EUW"},
	{Name: "crochecha", Tag: "EUW"},
	{Name: "jigoa", Tag: "XDD"},
	{Name: "nisenna", Tag: "EUW"},
}

var metals = map[string]int{
	"IRON":     0,
	"BRONZE":   1,
	"SILVER":   2,
	"GOLD":     3,
	"PLATINUM": 4,
	"EMERALD":  5,
	"DIAMOND":  6,
	"MASTER":   7,
}

var ranksMap = map[string]int{
	"I":   3,
	"II":  2,
	"III": 1,
	"IV":  0,
}

func main() {

	var playerStatsList []PlayerStats
	for _, user := range users {
		playerStats := getPlayerStats(user)
		playerStatsList = append(playerStatsList, playerStats)
	}

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
		c.JSON(200, playerStatsList)
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
