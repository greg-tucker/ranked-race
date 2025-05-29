package com.lemonfungus.RankedRace.controller;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.lemonfungus.RankedRace.model.SummonerRankData;
import com.lemonfungus.RankedRace.model.entities.RankEntryEntity;
import com.lemonfungus.RankedRace.service.RankService;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.io.ClassPathResource;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;
import java.io.InputStream;
import java.util.*;

@RestController
@Slf4j
@AllArgsConstructor
public class RaceController {
    private RankService rankService;

    private final ObjectMapper objectMapper = new ObjectMapper();

    @CrossOrigin
    @GetMapping("/rank")
    public Set<SummonerRankData> getRanks(){
        return rankService.getRanks();
    }

    @CrossOrigin(origins = "http://localhost:3000")
    @GetMapping("/ranktimeline")
    public List<Map<String, Object>> getRankTimeline(){
        return rankService.getRankTimeline();
    }
}
