package com.lemonfungus.RankedRace.model.dto;

public record LeagueEntryDto(
        String leagueid,
        String summonerId,
        String puuid,
        String queueType,
        String tier,
        String rank,
        int leaguePoints,
        int wins,
        int losses,
        boolean hotStreak,
        boolean veteran,
        boolean freshBlood,
        boolean inactive
) {
}
