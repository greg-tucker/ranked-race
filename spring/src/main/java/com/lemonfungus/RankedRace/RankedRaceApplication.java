package com.lemonfungus.RankedRace;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import lombok.extern.slf4j.Slf4j;
import org.springframework.scheduling.annotation.EnableScheduling;

@SpringBootApplication
@Slf4j
@EnableScheduling
public class RankedRaceApplication {
	public static void main(String[] args) {
		SpringApplication.run(RankedRaceApplication.class, args);
	}

}
