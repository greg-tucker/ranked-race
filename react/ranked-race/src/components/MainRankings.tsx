'use client';
import { MainRankingsData, visibleColumns } from '@/components/dataTypes';
import { useEffect, useState } from 'react';
import { getCurrentRanking } from '@/app/dataFetcher';
import { Table } from '@mantine/core';

export function MainRankings() {
  const [rankings, setRankings] = useState<MainRankingsData[]>([]);

  useEffect(() => {
    getCurrentRanking((data) => {
      const sortedData = [...data].sort((a, b) => b.current - a.current);
      setRankings(sortedData);
    });
  }, []);

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
              <Table.Tr key={row.name}>
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
