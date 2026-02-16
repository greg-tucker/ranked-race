'use client';
import React, { useEffect, useState } from 'react';
import { MainRankingsData, visibleColumns, LiveGame, loaderProp } from '@/components/dataTypes';
import { getCurrentRanking, getGameStats } from '@/app/dataFetcher';
import { Table, Stack, Group, Badge, Avatar } from '@mantine/core';
import { GameTimer } from './GameTimer';
import Image from 'next/image';

function useIsMobile() {
  const [isMobile, setIsMobile] = useState(false);
  useEffect(() => {
    const check = () => setIsMobile(window.innerWidth < 700);
    check();
    window.addEventListener('resize', check);
    return () => window.removeEventListener('resize', check);
  }, []);
  return isMobile;
}

export function MainRankings() {
  const [rankings, setRankings] = useState<MainRankingsData[]>([]);
  const isMobile = useIsMobile();
  const [activeGameData, setActiveGameData] = useState<LiveGame | null>(null);
  const [expandedPuuid, setExpandedPuuid] = useState<string | null>(null);
  const [isLoadingGame, setIsLoadingGame] = useState(false);

  function expand(puuid: string) {
    if (expandedPuuid === puuid) {
      setExpandedPuuid(null);
      setActiveGameData(null);
    } else {
      setExpandedPuuid(puuid);
      setIsLoadingGame(true);
      getGameStats(puuid, (data: LiveGame) => {
        setActiveGameData(data);
        setIsLoadingGame(false);
      });
    }
  }

  useEffect(() => {
    getCurrentRanking((data) => {
      const sortedData = [...data].sort((a, b) => b.current - a.current);
      setRankings(sortedData);
    });
  }, []);

  if (isMobile) {
    return (
      <div className="container">
        <h2 style={{ textAlign: 'center', marginBottom: 24, fontSize: 28, letterSpacing: 1, fontWeight: 700 }}>
          Live League Standings
        </h2>
        <div style={{ display: 'flex', flexDirection: 'column', gap: 16 }}>
          {rankings.map((row) => (
            <div key={row.name} className="glass" style={{ padding: 16, borderRadius: 12, boxShadow: '0 2px 12px #0ea5e933' }}>
              <div style={{ display: 'flex', alignItems: 'center', gap: 12, marginBottom: 8 }}>
                <span style={{ fontWeight: 700, fontSize: 18 }}>{row.name}</span>
                <span style={{ color: '#a8b5d8', fontSize: 14 }}>({row.tag})</span>
              </div>
              <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 8 }}>
                <img src={`/static/${row.tier.toLowerCase()}.png`} alt={row.tier} width={32} height={32} style={{ borderRadius: 6 }} />
                <span style={{ fontWeight: 600 }}>{row.displayRank}</span>
              </div>
              <div style={{ display: 'flex', flexWrap: 'wrap', gap: 12, fontSize: 15 }}>
                <span><b>LP:</b> {row.current}</span>
                <span><b>Peak:</b> {row.peak}</span>
                <span><b>Gained:</b> {row.gained}</span>
                <span><b>Wins:</b> <span className={row.winRate > 50 ? 'green' : 'red'}>{row.wins}</span></span>
                <span><b>Losses:</b> <span className={row.winRate > 50 ? 'green' : 'red'}>{row.losses}</span></span>
                <span><b>Winrate:</b> <span className={row.winRate > 50 ? 'green' : 'red'}>{row.winRate}%</span></span>
                <span><b>Games:</b> {row.played}</span>
                <a href={`https://www.op.gg/summoners/euw/${encodeURIComponent(row.name)}%20-${row.tag}`} target="_blank" rel="noopener noreferrer" className="opgg-link">
                  <img src="/static/opgg.svg" alt="OP.GG" width={50} height={20} style={{ verticalAlign: 'middle' }} />
                </a>
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  }

  return (
    <div className="container">
      <div className="glass" style={{ padding: '2rem', margin: 'auto', maxWidth: 1100, }}>
        <h2 style={{ textAlign: 'center', marginBottom: 32, fontSize: 32, letterSpacing: 1, fontWeight: 700 }}>
          Live League Standings
        </h2>
        <Table verticalSpacing="xs">
          <Table.Thead>
            <Table.Tr>
              {visibleColumns.map((column) => (
                <Table.Th key={column.key}>{column.label}</Table.Th>
              ))}
            </Table.Tr>
          </Table.Thead>

          <Table.Tbody>
            {rankings.map((row) => (
              <React.Fragment key={row.puuid ?? row.name}>
                <Table.Tr onClick={() =>{ if (row.inGame) expand(row.puuid)}} className={row.inGame ? 'inGame' : ''}>
                  {visibleColumns.map((column) => (
                    <Table.Td key={column.key}>
                      {typeof column.render === 'function' ? column.render(row) : String(row[column.key])}
                    </Table.Td>
                  ))}
                </Table.Tr>

                {expandedPuuid === row.puuid  && row.inGame  && (
                  <Table.Tr>
                    <Table.Td colSpan={visibleColumns.length}>
                      {isLoadingGame ? (
                        <div style={{ padding: 16, textAlign: 'center' }}>Loading game data...</div>
                      ) : activeGameData ? (
                        <Stack gap={16} style={{ padding: 16 }}>
                          <Group justify="space-between">
                            <div>
                              <Badge variant="light">{activeGameData.gameMode}</Badge>
                              <span style={{ marginLeft: 12 }}>
                                <GameTimer startTime={activeGameData.gameStartTime} />
                              </span>
                            </div>
                          </Group>

                          <Group grow align="flex-start">
                            {[100, 200].map((teamId) => (
                              <Stack key={teamId} gap={8}>
                                <Badge color={teamId === 100 ? 'blue' : 'red'} style={{ width: 'fit-content' }}>
                                  Team {teamId === 100 ? 'Blue' : 'Red'}
                                </Badge>
                                {activeGameData.participants
                                  .filter((p) => p.teamId === teamId)
                                  .map((participant) => (
                                    <Group
                                      key={participant.puuid}
                                      style={{
                                        padding: 8,
                                        borderRadius: 6,
                                        backgroundColor: 'rgba(255, 255, 255, 0.05)',
                                      }}
                                    >
                                      <Image
                                        src={`https://lolcdn.darkintaqt.com/cdn/champion/${participant.championId}/tile`}
                                        alt="champion"
                                        loader={loaderProp}
                                        width={32}
                                        height={32}
                                        style={{ borderRadius: 4 }}
                                      />
                                      <div style={{ flex: 1 }}>
                                        <div style={{ fontWeight: 500 }}>{participant.summonerName}</div>
                                        <div style={{ fontSize: 12, color: '#a8b5d8' }}>
                                          {participant.bot ? 'Bot' : `Level ${participant.profileIconId}`}
                                        </div>
                                      </div>
                                      <Group gap={4}>
                                        <img
                                          src={`https://ddragon.leagueoflegends.com/cdn/14.1.1/img/spell/Summoner${['D', 'F'][participant.spell1Id < participant.spell2Id ? 0 : 1]}.png`}
                                          alt="spell1"
                                          width={20}
                                          height={20}
                                          style={{ borderRadius: 2 }}
                                        />
                                        <img
                                          src={`https://ddragon.leagueoflegends.com/cdn/14.1.1/img/spell/Summoner${['D', 'F'][participant.spell1Id < participant.spell2Id ? 1 : 0]}.png`}
                                          alt="spell2"
                                          width={20}
                                          height={20}
                                          style={{ borderRadius: 2 }}
                                        />
                                      </Group>
                                    </Group>
                                  ))}
                              </Stack>
                            ))}
                          </Group>
                        </Stack>
                      ) : (
                        <div style={{ padding: 16, textAlign: 'center', color: '#a8b5d8' }}>
                          No active game data available
                        </div>
                      )}
                    </Table.Td>
                  </Table.Tr>
                )}
              </React.Fragment>
            ))}
          </Table.Tbody>
        </Table>
      </div>
    </div>
  );
}
