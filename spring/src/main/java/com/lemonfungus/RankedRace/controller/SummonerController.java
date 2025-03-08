package com.lemonfungus.RankedRace.controller;

import com.lemonfungus.RankedRace.model.dto.AccountDto;
import com.lemonfungus.RankedRace.model.dto.LeagueEntryDto;
import com.lemonfungus.RankedRace.model.dto.SummonerDto;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import com.lemonfungus.RankedRace.service.RiotApiService;

import java.util.Set;

@RestController
@Slf4j
@AllArgsConstructor
public class SummonerController {
    RiotApiService riotApiService;

    @GetMapping("/")
    public String getSummoner() {
        return riotApiService.getSummoner();
    }
    @GetMapping("/thread")
    public String getThreadName() {
        return Thread.currentThread().toString();
    }

    @GetMapping("/account/{name}/{tag}")
    public AccountDto getAccountByNameAndTag(@PathVariable String name, @PathVariable String tag) {
        return riotApiService.getAccountByNameAndTag(name, tag);
    }

    @GetMapping("/summoner/{name}/{tag}")
    public SummonerDto getSummoner(@PathVariable String name, @PathVariable String tag) {
        var puuid = riotApiService.getAccountByNameAndTag(name, tag).puuid();
        log.info(puuid);
        return riotApiService.getSummonerByPuuid(puuid);
    }

    @GetMapping("/entries/{name}/{tag}")
    public Set<LeagueEntryDto> getEntries(@PathVariable String name, @PathVariable String tag) {
        var puuid = riotApiService.getAccountByNameAndTag(name, tag).puuid();
        log.info(puuid);
        return riotApiService.getEntriesByPuuid(puuid);
    }
}