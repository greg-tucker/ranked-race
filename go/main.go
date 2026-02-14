package main

import (
	"fmt"
	"math"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"log"

	_ "github.com/joho/godotenv/autoload"
)

type PlayerStats struct {
	Name        string  `json:"name"`
	Peak        int     `json:"peak"`
	Gained      int     `json:"gained"`
	Current     int     `json:"current"`
	Wins        int     `json:"wins"`
	Losses      int     `json:"losses"`
	Played      int     `json:"played"`
	Tier        string  `json:"tier"`
	Rank        string  `json:"rank"`
	LP          int     `json:"lp"`
	Date        string  `json:"date"`
	WinRate     float64 `json:"winRate"`
	DisplayRank string  `json:"displayRank"`
	Tag         string  `json:"tag"`
	InGame      bool    `json:"inGame"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toPlayerStats(entry RankedEntry, name string, tag string) PlayerStats {
	totalGames := entry.Wins + entry.Losses
	var winrate float64 = 0
	if totalGames != 0 {
		winrate = math.Round(10000*float64(entry.Wins)/(float64(entry.Wins+entry.Losses))) / 100
	}
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
		WinRate:     winrate,
		DisplayRank: entry.Tier + " " + entry.Rank + " " + fmt.Sprint(entry.LeaguePoints) + " LP",
		Tag:         tag,
		InGame:      false,
	}
}

func calculateLp(entry RankedEntry) int {
	metalIndex := metals[entry.Tier]
	return metalIndex*400 + ranksMap[entry.Rank]*100 + entry.LeaguePoints
}

func filterRankedEntriesForGameType(entries []RankedEntry, queueType string) (entry RankedEntry, found bool) {
	for i := range entries {
		if entries[i].QueueType == SOLO_QUEUE {
			return entries[i], true
		}
	}
	return RankedEntry{}, false
}

func getPlayerStats(inputPlayer InputPlayer) (player PlayerStats, found bool) {
	acc, err := getAccountFromNameAndTag(inputPlayer)
	if err != nil {
		log.Fatalln(err)
		return PlayerStats{}, false
	}

	log.Printf("%+v\n", acc)
	log.Printf("PUUID LIT NASTY %s", acc.PUUID)

	rankedEntries, err := getRankedEntriesByAcc(acc)

	soloQueueEntry, found := filterRankedEntriesForGameType(rankedEntries, SOLO_QUEUE)

	if found == false {
		return PlayerStats{}, false
	}

	log.Printf("SOLO QUEUE %+v\n", soloQueueEntry)

	playerStats := toPlayerStats(soloQueueEntry, acc.GameName, inputPlayer.Tag)

	log.Printf("PLAYERSTATS %+v\n", playerStats)

	_, found = getActiveGamesByPuuid(acc.PUUID)

	if found {
		playerStats.InGame = true
	}

	return playerStats, true
}

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
	{Name: "OneLargeBoi", Tag: "EUW"},
	{Name: "gemgeffery", Tag: "EUW"},
	{Name: "Purple Volvo", Tag: "EUW"},
	{Name: "soap tastes ok", Tag: "EUW"},
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

var SOLO_QUEUE = "RANKED_SOLO_5x5"

func loadPlayerData() []PlayerStats {
	var playerStatsList []PlayerStats
	for _, user := range users {
		playerStats, found := getPlayerStats(user)
		if found {
			playerStatsList = append(playerStatsList, playerStats)
		}
	}
	return playerStatsList
}

func main() {
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
		c.JSON(200, loadPlayerData())
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
