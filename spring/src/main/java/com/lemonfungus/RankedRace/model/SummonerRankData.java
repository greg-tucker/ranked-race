package com.lemonfungus.RankedRace.model;

import lombok.Builder;

@Builder
public record SummonerRankData(
        String name,
        int peak,
        int gained,
        int current,
        int wins,
        int losses,
        int played,
        String tier,
        String rank,
        int lp
) {
}
