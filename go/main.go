package main

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"sync"
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
	StartTime   uint64  `json:"startTime"`
	PUUID       string  `json:"puuid"`
	Role        string  `json:"role"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toPlayerStats(entry RankedEntry, acc Account, tag string) PlayerStats {
	totalGames := entry.Wins + entry.Losses
	var winrate float64 = 0
	if totalGames != 0 {
		winrate = math.Round(10000*float64(entry.Wins)/(float64(entry.Wins+entry.Losses))) / 100
	}
	return PlayerStats{
		Name:        acc.GameName,
		PUUID:       acc.PUUID,
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

type RoleCount struct {
	Role  string
	Count int
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

	playerStats := toPlayerStats(soloQueueEntry, acc, inputPlayer.Tag)

	log.Printf("PLAYERSTATS %+v\n", playerStats)

	activeGame, found := getActiveGamesByPuuid(acc.PUUID)

	if found {
		playerStats.InGame = true
		playerStats.StartTime = uint64(activeGame.GameStartTime)
		log.Println(activeGame.GameStartTime)
	}

	matchHistory, found := getMatchHistoryByPuuid(acc.PUUID)

	log.Printf("THE ATCHES AFTER SUBSLICE : %+v", matchHistory)

	playerStats.Role = getMostPlayedRole(matchHistory, acc.PUUID)

	return playerStats, true
}

func getMostPlayedRole(matchHistoryIds []string, puuid string) (role string) {
	rolesMap := make(map[string]int)

	for _, matchId := range matchHistoryIds {
		match, found := getMatchByMatchId(matchId)

		if !found {
			return ""
		}
		for _, player := range match.Info.Participants {
			if player.PUUID == puuid {
				rolesMap[player.TeamPosition]++
			}
		}
	}

	var roleCounts []RoleCount
	for role, count := range rolesMap {
		roleCounts = append(roleCounts, RoleCount{role, count})
	}

	sort.Slice(roleCounts, func(i, j int) bool {
		return roleCounts[i].Count > roleCounts[j].Count
	})

	log.Printf("ROLES PLAYED: %+v", roleCounts)

	if len(roleCounts) > 0 {
		return roleCounts[0].Role
	}
	log.Printf("NO ROLES? %+v %+v", roleCounts, rolesMap)
	return ""
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
	{Name: "GediGrandMaster", Tag: "EUW"},
	{Name: "Yellow Lada", Tag: "EUW"},
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
var cachedPlayers []PlayerStats
var mu sync.RWMutex
var lastUpdated time.Time
var cacheTTL = 60 * time.Second

func refreshCache() {
	log.Println("Refreshing Riot cache...")

	fresh := loadPlayerData()

	mu.Lock()
	cachedPlayers = fresh
	mu.Unlock()

	log.Println("Cache refreshed")
}

func startCacheRefresher(interval time.Duration) {
	// Initial load
	refreshCache()

	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			refreshCache()
		}
	}()
}

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
	startCacheRefresher(3 * time.Minute)

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
		mu.RLock()
		data := cachedPlayers
		mu.RUnlock()
		c.JSON(200, data)
	})
	router.GET("/activeGame/:PUUID", func(c *gin.Context) {
		PUUID := c.Param("PUUID")
		game, found := getActiveGamesByPuuid(PUUID)
		if !found {
			c.String(http.StatusNotFound, "NO ACTIVE GAME FOR USER")
		} else {
			c.JSON(200, game)
		}
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
