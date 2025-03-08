package com.lemonfungus.RankedRace.model;

import com.lemonfungus.RankedRace.model.entities.RankEntryEntity;
import lombok.Builder;

import java.util.Date;

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
    public RankEntryEntity toPlayerEntryEntity(Date date) {
        return RankEntryEntity.builder()
                .date(date)
                .losses(losses)
                .name(name)
                .gained(gained)
                .current(current)
                .wins(wins)
                .losses(losses)
                .played(played)
                .tier(tier)
                .rank(rank)
                .lp(lp)
                .build();
    }
}
