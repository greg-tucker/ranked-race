package com.lemonfungus.RankedRace.model.entities;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;

import java.util.Date;

@Entity
@AllArgsConstructor
@Data
@Builder
@Table(name = "rankEntry")
public class RankEntryEntity {

    private RankEntryEntity() {

    }

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Integer id;

    String name;
    int peak;
    int gained;
    int current;
    int wins;
    int losses;
    int played;
    String tier;
    String rank;
    int lp;
    Date date;
}
