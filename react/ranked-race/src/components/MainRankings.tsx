'use client';
import { MainRankingsData, visibleColumns } from '@/app/dataTypes';
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
    <div>
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
          {'render' in column
            ? column.render(row)
            : String(row[column.key])}
        </Table.Td>
      ))}
    </Table.Tr>
  ))}
</Table.Tbody>

      </Table>
    </div>
  );
}
