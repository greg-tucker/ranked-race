package com.lemonfungus.RankedRace.controller;

import com.lemonfungus.RankedRace.model.SummonerRankData;
import com.lemonfungus.RankedRace.service.RankService;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Set;

@RestController
@Slf4j
@AllArgsConstructor
public class RaceController {
    private RankService rankService;

    @GetMapping("/rank")
    public Set<SummonerRankData> getRanks(){
        return rankService.getRanks();
    }
}
