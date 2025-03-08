package com.lemonfungus.RankedRace.service;

import com.lemonfungus.RankedRace.config.RankedRaceProperties;
import com.lemonfungus.RankedRace.model.Ranks;
import com.lemonfungus.RankedRace.model.SummonerRankData;
import com.lemonfungus.RankedRace.model.Tiers;
import com.lemonfungus.RankedRace.model.dto.LeagueEntryDto;

import java.util.HashSet;
import java.util.Set;

public class RankService {
    private final RiotApiService riotApiService;
    private final RankedRaceProperties rankedRaceProperties;

    private final String QUEUE_TYPE = "RANKED_SOLO_5x5";

    public RankService(RiotApiService riotApiService, RankedRaceProperties rankedRaceProperties){
        this.riotApiService = riotApiService;
        this.rankedRaceProperties = rankedRaceProperties;
    }

    public Set<SummonerRankData> getRanks() {
        var players = rankedRaceProperties.getPlayers();
        var outputSet = new HashSet<SummonerRankData>();
        for (var player: players) {
            var account = riotApiService.getAccountByNameAndTag(player.name(), player.tag());
            var rankedEntries = riotApiService.getEntriesByPuuid(account.puuid());
            var optionalEntry = rankedEntries.stream().filter(entry -> QUEUE_TYPE.equals(entry.queueType())).findFirst();

            if (optionalEntry.isEmpty()) {
                continue;
            }

            var soloqEntry = optionalEntry.get();

            outputSet.add(
                    SummonerRankData.builder()
                            .name(player.name())
                            //TODO: Add db call for peak
                            .peak(999)
                            .gained(calcLp(soloqEntry) - player.lp())
                            .current(calcLp(soloqEntry))
                            .wins(soloqEntry.wins())
                            .losses(soloqEntry.losses())
                            .played(soloqEntry.losses() + soloqEntry.wins())
                            .tier(soloqEntry.tier())
                            .rank(soloqEntry.rank())
                            .lp(soloqEntry.leaguePoints())
                            .build()
            );
        }
        return outputSet;
    }

    private int calcLp(LeagueEntryDto leagueEntryDto) {
        return ((Tiers.valueOf(leagueEntryDto.tier()).ordinal()) * 400) +
                (Ranks.valueOf(leagueEntryDto.rank()).ordinal() * 100)
                + leagueEntryDto.leaguePoints();
    }
}
