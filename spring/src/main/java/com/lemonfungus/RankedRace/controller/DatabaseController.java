package com.lemonfungus.RankedRace.controller;

import com.lemonfungus.RankedRace.model.entities.RankEntryEntity;
import com.lemonfungus.RankedRace.repositories.RankEntryRepository;
import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.util.Optional;

@RestController
@AllArgsConstructor
public class DatabaseController {
    RankEntryRepository rankEntryRepository;

    @GetMapping("/database")
    public Iterable<RankEntryEntity> findAllRankEntries() {
        return this.rankEntryRepository.findAll();
    }

    @GetMapping("/database/{name}")
    public Optional<RankEntryEntity> findAllRankEntries(@PathVariable("name") String name) {
        return this.rankEntryRepository.findTopByNameOrderByGainedDesc(name);
    }

    @PostMapping("/database")
    public RankEntryEntity addOneEmployee(@RequestBody RankEntryEntity entry) {
        return this.rankEntryRepository.save(entry);
    }
}
