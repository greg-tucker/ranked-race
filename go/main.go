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
	Image       string  `json:"image"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toPlayerStats(entry RankedEntry, acc Account, tag string, image string) PlayerStats {
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
		Image:       image,
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
		if entries[i].QueueType == queueType {
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

	log.Printf("SOLO QUEUE %+v\n", soloQueueEntry)

	playerStats := toPlayerStats(soloQueueEntry, acc, inputPlayer.Tag, inputPlayer.Image)

	log.Printf("PLAYERSTATS %+v\n", playerStats)
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
	{Name: "Impala", Tag: "KAZ", Image: "mayalover3.jpg"},
	{Name: "Bjerkingfan", Tag: "EUW", Image: "haudyerwheesht.jpg"},
	{Name: "HaudYerWheesht", Tag: "EUW", Image: "haudyerwheesht.jpg"},
	{Name: "ctrl alt cute", Tag: "xoxo", Image: "oystericetea.jpg"},
	{Name: "oystericetea", Tag: "EUW", Image: "oystericetea.jpg"},
	{Name: "Pissinglnthewind", Tag: "EUW", Image: "obe.png"},
	{Name: "mayalover3", Tag: "EUW", Image: "mayalover3.jpg"},
	{Name: "SnazzyG", Tag: "EUW", Image: "ali.jpg"},
	{Name: "crochecha", Tag: "EUW", Image: "joe.JPG"},
	{Name: "jigoa", Tag: "XDD", Image: "jigoa.jpg"},
	{Name: "nisenna", Tag: "EUW", Image: "nichy.jpg"},
	{Name: "OneLargeBoi", Tag: "EUW", Image: "duncan.jpg"},
	{Name: "gemgeffery", Tag: "EUW", Image: "mayalover3.jpg"},
	{Name: "Purple Volvo", Tag: "EUW", Image: "volvo.jpg"},
	{Name: "soap tastes ok", Tag: "EUW", Image: "obe.png"},
	{Name: "GediGrandMaster", Tag: "EUW", Image: "geddes.jpg"},
	{Name: "Yellow Lada", Tag: "EUW", Image: "nichy.jpg"},
	{Name: "JGeddes", Tag: "EUW", Image: "geddes.jpg"},
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
var cachedActiveGames map[string]CurrentGameInfo
var cachedRoles map[string]string
var playerMu sync.RWMutex
var activeGamesMu sync.RWMutex
var rolesMu sync.RWMutex
var lastUpdated time.Time

func startPlayerCacheRefresher(interval time.Duration) {
	// Initial load
	refreshPlayersCache()

	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			refreshPlayersCache()
		}
	}()
}

func refreshPlayersCache() {
	log.Println("Refreshing players cache...")

	fresh := loadPlayerData()

	playerMu.Lock()
	cachedPlayers = fresh
	playerMu.Unlock()

	log.Println("Player Cache refreshed")
}

func startActiveGamesCacheRefresher(interval time.Duration) {
	// Initial load
	refreshActiveGamesCache()

	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			refreshActiveGamesCache()
		}
	}()
}

func refreshActiveGamesCache() {
	log.Println("Refreshing active games cache...")

	fresh := loadActiveGamesData()

	activeGamesMu.Lock()
	cachedActiveGames = fresh
	activeGamesMu.Unlock()

	log.Println("Active Games Cache refreshed")
}

func startRolesCacheRefresher(interval time.Duration) {
	// Initial load
	refreshRolesCache()

	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			refreshRolesCache()
		}
	}()
}

func refreshRolesCache() {
	log.Println("Refreshing roles cache...")

	fresh := loadRolesData()

	rolesMu.Lock()
	cachedRoles = fresh
	rolesMu.Unlock()

	log.Println("Roles Cache refreshed")
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

func getPuuidsFromCachedPlayers() []string {
	var puuids []string
	playerMu.RLock()
	for _, player := range cachedPlayers {
		puuids = append(puuids, player.PUUID)
	}
	playerMu.RUnlock()
	return puuids
}

func loadActiveGamesData() map[string]CurrentGameInfo {
	puuids := getPuuidsFromCachedPlayers()
	activeGamesMap := make(map[string]CurrentGameInfo, len(puuids))
	for _, puuid := range puuids {
		matches, found := getActiveGamesByPuuid(puuid)
		if found {
			activeGamesMap[puuid] = matches
		}
	}
	return activeGamesMap
}

func loadRolesData() map[string]string {
	puuids := getPuuidsFromCachedPlayers()
	rolesMap := make(map[string]string, len(puuids))
	for _, puuid := range puuids {
		matches, found := getMatchHistoryByPuuid(puuid)
		if found {
			rolesMap[puuid] = getMostPlayedRole(matches, puuid)
		}
	}
	return rolesMap
}

func combineCachesForRankEndpoint() []PlayerStats {
	playerMu.RLock()
	activeGamesMu.RLock()
	rolesMu.RLock()
	defer playerMu.RUnlock()
	defer activeGamesMu.RUnlock()
	defer rolesMu.RUnlock()
	var combinedPlayers []PlayerStats
	for _, player := range cachedPlayers {
		player.Role = cachedRoles[player.PUUID]
		activeGame, ok := cachedActiveGames[player.PUUID]
		if ok {
			player.InGame = true
			player.StartTime = uint64(activeGame.GameStartTime)
		}
		combinedPlayers = append(combinedPlayers, player)
	}
	return combinedPlayers
}

func main() {
	startPlayerCacheRefresher(3 * time.Minute)
	startActiveGamesCacheRefresher(30 * time.Second)
	startRolesCacheRefresher(30 * time.Minute)

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
		c.JSON(200, combineCachesForRankEndpoint())
	})
	router.GET("/activeGame/:PUUID", func(c *gin.Context) {
		PUUID := c.Param("PUUID")
		activeGamesMu.RLock()
		activeGame, ok := cachedActiveGames[PUUID]
		activeGamesMu.RUnlock()
		if !ok {
			c.String(http.StatusNotFound, "NO ACTIVE GAME FOR USER")
		} else {
			c.JSON(200, activeGame)
		}
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
