package com.lemonfungus.RankedRace.config;

import com.lemonfungus.RankedRace.repositories.RankEntryRepository;
import com.lemonfungus.RankedRace.service.RankService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.lemonfungus.RankedRace.service.RiotApiService;
import org.springframework.web.client.RestClient;

@Slf4j
@Configuration
public class RitoConfig {
    @Value( "${riot.key}" )
    private String riotApiKey;
    @Value( "${riot.region}" )
    private String riotRegion;
    @Value( "${riot.url}" )
    private String riotUrl;
    @Value( "${riot.account-url}" )
    private String accountUrl;

    @Bean
    public RiotApiService riotApiService(RestClient riotWebClient, RestClient accountWebClient){
        return new RiotApiService(riotWebClient, accountWebClient, riotApiKey, riotRegion);
    }

    @Bean
    public RankService rankService(RiotApiService riotApiService, RankedRaceProperties rankedRaceProperties
            , RankEntryRepository rankEntryRepository){
        log.info("Propertes {}", rankedRaceProperties.getPlayers());
        return new RankService(riotApiService, rankedRaceProperties, rankEntryRepository);
    }

    @Bean
    public RestClient accountWebClient() {
        return RestClient.builder()
                .baseUrl(accountUrl)
                .defaultHeader("X-Riot-Token", riotApiKey)
                .build();
    }

    @Bean
    public RestClient riotWebClient() {
        return RestClient.builder()
                .baseUrl(riotUrl)
                .defaultHeader("X-Riot-Token", riotApiKey)
                .build();
    }

    @Bean
    public RankedRaceProperties rankedRaceProperties(){
        return new RankedRaceProperties();
    }
}
