package com.lemonfungus.RankedRace.config;

import com.lemonfungus.RankedRace.model.PlayerEntry;
import lombok.Data;
import lombok.Getter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

import java.util.List;

@Data
@ConfigurationProperties(prefix = "rr")
public class RankedRaceProperties {
    private List<PlayerEntry> players;
}
