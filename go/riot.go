package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"context"

	"golang.org/x/time/rate"
)

type Account struct {
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type RankedEntry struct {
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
}

type CurrentGameInfo struct {
	GameID            int64                    `json:"gameId"`
	GameType          string                   `json:"gameType"`
	GameStartTime     int64                    `json:"gameStartTime"`
	MapID             int64                    `json:"mapId"`
	GameLength        float64                  `json:"gameLength"` // double in Riot API
	PlatformID        string                   `json:"platformId"`
	GameMode          string                   `json:"gameMode"`
	BannedChampions   []BannedChampion         `json:"bannedChampions"`
	GameQueueConfigID int64                    `json:"gameQueueConfigId"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
}

type CurrentGameParticipant struct {
	ChampionID               int64                     `json:"championId"`
	Perks                    Perks                     `json:"perks"`
	ProfileIconID            int64                     `json:"profileIconId"`
	Bot                      bool                      `json:"bot"`
	TeamID                   int64                     `json:"teamId"`
	RiotID                   string                    `json:"riotId"`
	SummonerID               string                    `json:"summonerId"`
	Puuid                    string                    `json:"puuid"`
	Spell1ID                 int64                     `json:"spell1Id"`
	Spell2ID                 int64                     `json:"spell2Id"`
	GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
}

type Perks struct {
	PerkIDs      []int64 `json:"perkIds"`
	PerkStyle    int64   `json:"perkStyle"`
	PerkSubStyle int64   `json:"perkSubStyle"`
}

type GameCustomizationObject struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}

type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

type BannedChampion struct {
	PickTurn   int64 `json:"pickTurn"`
	ChampionID int64 `json:"championId"`
	TeamID     int64 `json:"teamId"`
}

type Match struct {
	Info MatchInfo `json:"info"`
}

type MatchInfo struct {
	Participants []Participants `json:"participants"`
}

type Participants struct {
	TeamPosition string `json:"teamPosition"`
	PUUID        string `json:"puuid"`
}

type InputPlayer struct {
	Name  string
	Tag   string
	Image string
}

var shortLimiter = rate.NewLimiter(20, 20)
var longLimiter = rate.NewLimiter(rate.Every(time.Minute/50), 50)

var apiKey = os.Getenv("RIOT_API_KEY")
var riotBaseUrl = "https://europe.api.riotgames.com/"
var leagueBaseUrl = "https://euw1.api.riotgames.com/"
var accountsPath = "riot/account/v1/accounts/by-riot-id/"
var entriesPath = "lol/league/v4/entries/by-puuid/"
var specatatorPath = "lol/spectator/v5/active-games/by-summoner/"
var matchHistoryPath = "lol/match/v5/matches/by-puuid/%s/ids"
var matchPath = "lol/match/v5/matches/%s"
var apiKeyParam = "?api_key=" + apiKey

func callRiot(url string) ([]byte, error) {
	return callRiotWithQueryParams(url, "")
}

func callRiotWithQueryParams(url string, queryParams string) ([]byte, error) {
	err := shortLimiter.Wait(context.Background())

	if err != nil {
		return nil, err
	}

	err = longLimiter.Wait(context.Background())

	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url + apiKeyParam + queryParams)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return body, nil
}

func getAccountFromNameAndTag(inputPlayer InputPlayer) (Account, error) {
	body, err := callRiot(riotBaseUrl + accountsPath + inputPlayer.Name + "/" + inputPlayer.Tag)

	if err != nil {
		log.Print(err)
		return Account{}, err
	}

	var acc Account
	err = json.Unmarshal(body, &acc)
	if err != nil {
		log.Print(err)
		return Account{}, err
	}
	return acc, nil
}

func getRankedEntriesByAcc(acc Account) ([]RankedEntry, error) {
	body, err := callRiot(leagueBaseUrl + entriesPath + acc.PUUID)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var rankedEntries []RankedEntry
	err = json.Unmarshal(body, &rankedEntries)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return rankedEntries, nil
}

func getActiveGamesByPuuid(puuid string) (currentGame CurrentGameInfo, found bool) {
	body, err := callRiot(leagueBaseUrl + specatatorPath + puuid)
	if err != nil {
		log.Print(err)
		return CurrentGameInfo{}, false
	}

	var currentGameInfo CurrentGameInfo
	err = json.Unmarshal(body, &currentGameInfo)
	if err != nil {
		log.Print(err)
		return CurrentGameInfo{}, false
	}

	if currentGameInfo.GameID == 0 {
		return currentGameInfo, false
	}

	return currentGameInfo, true
}

func getMatchHistoryByPuuid(puuid string) (matchIds []string, found bool) {
	amountOfGamesToLoad := 1
	body, err := callRiotWithQueryParams(riotBaseUrl+fmt.Sprintf(matchHistoryPath, puuid), "&queue=420")

	if err != nil {
		log.Print(err)
		return nil, false
	}

	err = json.Unmarshal(body, &matchIds)
	if err != nil {
		log.Print(err)
		return nil, false
	}
	if len(matchIds) < amountOfGamesToLoad {
		return matchIds, true
	}
	return matchIds[:amountOfGamesToLoad], true
}

func getMatchByMatchId(matchId string) (match Match, found bool) {
	body, err := callRiot(riotBaseUrl + fmt.Sprintf(matchPath, matchId))
	if err != nil {
		log.Print(err)
		return Match{}, false
	}

	err = json.Unmarshal(body, &match)
	if err != nil {
		log.Print(err)
		return Match{}, false
	}

	return match, true
}
