package com.lemonfungus.RankedRace.controller;

import com.lemonfungus.RankedRace.model.SummonerRankData;
import com.lemonfungus.RankedRace.model.entities.RankEntryEntity;
import com.lemonfungus.RankedRace.service.RankService;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.Map;
import java.util.Set;

@RestController
@Slf4j
@AllArgsConstructor
public class RaceController {
    private RankService rankService;

    @CrossOrigin
    @GetMapping("/rank")
    public Set<SummonerRankData> getRanks(){
        return rankService.getRanks();
    }

    @CrossOrigin
    @GetMapping("/ranktimeline")
    public Map<String, List<RankEntryEntity>> getRankTimeline(){
        return rankService.getRankTimeline();
    }

    @CrossOrigin(origins = "http://localhost:3000")
    @GetMapping("/ranktimeline/{player}")
    public List<Map<String, Object>> getIndividualRankTimeline(@PathVariable String player){
        log.info("Getting data for " + player);
        return rankService.getIndividualRankTimeline(player);
    }
}
