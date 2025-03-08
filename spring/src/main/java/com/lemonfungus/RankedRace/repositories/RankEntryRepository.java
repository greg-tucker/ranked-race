package com.lemonfungus.RankedRace.repositories;

import com.lemonfungus.RankedRace.model.entities.RankEntryEntity;
import org.springframework.data.repository.CrudRepository;

import java.util.List;
import java.util.Optional;

public interface RankEntryRepository extends CrudRepository<RankEntryEntity, Integer> {
    List<RankEntryEntity> findByName(String name);

    Optional<RankEntryEntity> findTopByNameOrderByGainedDesc(String name);

    List<RankEntryEntity> findByNameOrderByDate(String name);

}
