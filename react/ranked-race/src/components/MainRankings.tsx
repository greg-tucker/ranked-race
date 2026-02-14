'use client';
import { MainRankingsData, visibleColumns } from '@/components/dataTypes';
import { useEffect, useState } from 'react';
import { getCurrentRanking } from '@/app/dataFetcher';
import { Table } from '@mantine/core';

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
      <div className="glass" style={{ padding: '2rem', margin: 'auto', maxWidth: 1100, boxShadow: '0 4px 32px #0ea5e933' }}>
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
              <Table.Tr key={row.name} className={row.inGame ?  'inGame' : ''}>
                {visibleColumns.map((column) => (
                  <Table.Td key={column.key}>
                    {typeof column.render === 'function'
                      ? column.render(row)
                      : String(row[column.key])}
                  </Table.Td>
                ))}
              </Table.Tr>
            ))}
          </Table.Tbody>
        </Table>
      </div>
    </div>
  );
}
