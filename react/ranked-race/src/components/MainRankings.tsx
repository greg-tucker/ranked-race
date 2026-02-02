'use client';
import { MainRankingsData, visibleColumns } from '@/app/dataTypes';
import { useEffect, useState } from 'react';
import { getCurrentRanking } from '@/app/dataFetcher';
import { Table } from '@mantine/core';

export function MainRankings() {
  const [rankings, setRankings] = useState<MainRankingsData[]>([]);

  useEffect(() => {
    getCurrentRanking((data) => {
      const sortedData = [...data].sort((a, b) => b.gained - a.gained);
      setRankings(sortedData);
    });
  }, []);

  return (
    <div>
      <Table verticalSpacing="xs">
        <Table.Thead>
          <Table.Tr>
            {visibleColumns.map((column) => (
              <Table.Th key={column}>{column}</Table.Th>
            ))}
          </Table.Tr>
        </Table.Thead>
        <Table.Tbody>
          {rankings.map((row) => (
            <Table.Tr key={row.name + 'row'}>
              {visibleColumns.map((column) => (
                <Table.Td key={row.name + column}>{String(row[column])}</Table.Td>
              ))}
            </Table.Tr>
          ))}
        </Table.Tbody>
      </Table>
    </div>
  );
}
