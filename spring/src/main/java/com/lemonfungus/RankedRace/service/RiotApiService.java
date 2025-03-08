package com.lemonfungus.RankedRace.service;

import com.lemonfungus.RankedRace.model.dto.AccountDto;
import com.lemonfungus.RankedRace.model.dto.LeagueEntryDto;
import com.lemonfungus.RankedRace.model.dto.SummonerDto;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.web.client.RestClient;

import java.util.Set;

@Slf4j
public class RiotApiService {
    private final String apiKey;
    private final String region;
    private final RestClient restClient;
    private final RestClient accountRestClient;


    public RiotApiService(RestClient restClient, RestClient accountRestClient, String apiKey, String region) {
        log.info("Starting api service with region {} and key {}", region, apiKey);
        this.apiKey = apiKey;
        this.region = region;
        this.restClient = restClient;
        this.accountRestClient = accountRestClient;

    }

    public String getSummoner(){
        return "lmao answer";
    }

    public AccountDto getAccountByNameAndTag(String name, String tag) {
        return accountRestClient
                .get()
                .uri("riot/account/v1/accounts/by-riot-id/{name}/{tag}", name, tag)
                .retrieve()
                .body(AccountDto.class);
    }

    public SummonerDto getSummonerByPuuid(String puuid) {
        return restClient
                .get()
                .uri("lol/summoner/v4/summoners/by-puuid/{puuid}", puuid)
                .retrieve()
                .body(SummonerDto.class);
    }

    public Set<LeagueEntryDto> getEntriesByPuuid(String puuid) {
        return restClient
                .get()
                .uri("lol/league/v4/entries/by-puuid/{puuid}", puuid)
                .retrieve()
                .body(new ParameterizedTypeReference<Set<LeagueEntryDto>>() {
                });
    }
}
