package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
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
	SummonerName             string                    `json:"summonerName"`
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

type InputPlayer struct {
	Name string
	Tag  string
}

var apiKey = os.Getenv("RIOT_API_KEY")
var riotBaseUrl = "https://europe.api.riotgames.com/"
var leagueBaseUrl = "https://euw1.api.riotgames.com/"
var accountsPath = "riot/account/v1/accounts/by-riot-id/"
var entriesPath = "lol/league/v4/entries/by-puuid/"
var specatatorPath = "/lol/spectator/v5/active-games/by-summoner/"
var apiKeyParam = "?api_key=" + apiKey

func callRiot(url string) ([]byte, error) {
	resp, err := http.Get(url + apiKeyParam)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return body, nil
}

func getAccountFromNameAndTag(inputPlayer InputPlayer) (Account, error) {
	body, err := callRiot(riotBaseUrl + accountsPath + inputPlayer.Name + "/" + inputPlayer.Tag)

	if err != nil {
		log.Fatalln(err)
		return Account{}, err
	}

	var acc Account
	err = json.Unmarshal(body, &acc)
	if err != nil {
		log.Fatalln(err)
		return Account{}, err
	}
	return acc, nil
}

func getRankedEntriesByAcc(acc Account) ([]RankedEntry, error) {
	body, err := callRiot(leagueBaseUrl + entriesPath + acc.PUUID)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var rankedEntries []RankedEntry
	err = json.Unmarshal(body, &rankedEntries)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return rankedEntries, nil
}

func getActiveGamesByPuuid(puuid string) (currentGame CurrentGameInfo, found bool) {
	body, err := callRiot(specatatorPath + puuid)
	if err != nil {
		log.Fatalln(err)
		return CurrentGameInfo{}, false
	}

	var currentGameInfo CurrentGameInfo
	err = json.Unmarshal(body, &currentGameInfo)
	if err != nil {
		log.Fatalln(err)
		return CurrentGameInfo{}, false
	}

	return currentGameInfo, true
}
