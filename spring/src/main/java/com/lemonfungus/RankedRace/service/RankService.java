package com.lemonfungus.RankedRace.service;

import com.lemonfungus.RankedRace.config.RankedRaceProperties;
import com.lemonfungus.RankedRace.model.Ranks;
import com.lemonfungus.RankedRace.model.SummonerRankData;
import com.lemonfungus.RankedRace.model.Tiers;
import com.lemonfungus.RankedRace.model.dto.LeagueEntryDto;
import com.lemonfungus.RankedRace.model.entities.RankEntryEntity;
import com.lemonfungus.RankedRace.repositories.RankEntryRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.scheduling.annotation.Scheduled;

import java.util.Date;
import java.util.HashSet;
import java.util.Set;

@Slf4j
public class RankService {
    private final RiotApiService riotApiService;
    private final RankedRaceProperties rankedRaceProperties;
    private final RankEntryRepository rankEntryRepository;

    private final String QUEUE_TYPE = "RANKED_SOLO_5x5";

    public RankService(RiotApiService riotApiService, RankedRaceProperties rankedRaceProperties,
                       RankEntryRepository rankEntryRepository){
        this.riotApiService = riotApiService;
        this.rankedRaceProperties = rankedRaceProperties;
        this.rankEntryRepository = rankEntryRepository;
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

            var peakPlayer = rankEntryRepository.findTopByNameOrderByGainedDesc(player.name());

            outputSet.add(
                    SummonerRankData.builder()
                            .name(player.name())
                            .peak(peakPlayer.map(RankEntryEntity::getGained).orElse(-1))
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

    @Scheduled(fixedRate = 2 * 60 * 1000)
    private void syncDataToDatabase() {
        log.info("Syncing data");
        var allRanks = getRanks();
        for (var rank: allRanks) {
            log.info("Writing data for {}", rank.name());
            rankEntryRepository.save(rank.toPlayerEntryEntity(new Date()));
        }
    }

    private int calcLp(LeagueEntryDto leagueEntryDto) {
        return ((Tiers.valueOf(leagueEntryDto.tier()).ordinal()) * 400) +
                (Ranks.valueOf(leagueEntryDto.rank()).ordinal() * 100)
                + leagueEntryDto.leaguePoints();
    }
}
