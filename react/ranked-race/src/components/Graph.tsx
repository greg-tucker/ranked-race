import { MainRankingsData } from '@/components/dataTypes';
import { useState, useEffect } from 'react';
import { LineChart } from '@mantine/charts';
import { Group, Image, Text, Stack } from '@mantine/core';

import '@mantine/charts/styles.css';

export function Graph() {
  const [rankHistory, setRankHistory] = useState<MainRankingsData[]>([]);

  useEffect(() => {
    // const fetchAndSet = () => getRankHistory(setRankHistory);
    // fetchAndSet();
    // const interval = setInterval(fetchAndSet, 120000); // poll every 2min
    // return () => clearInterval(interval); // cleanup
  }, []);

  const series = [
    { name: 'mayalover3', color: '#aa80ff', icon: 'static/mayalover3.jpg' },
    { name: 'jigoa', color: '#ff8000', icon: '/static/jigoa.jpg' },
    {
      name: 'oystericetea',
      color: '#ff99ff',
      icon: '/static/oystericetea.jpg',
    },
    {
      name: 'haudyerwheesht',
      color: '#8000ff',
      icon: '/static/haudyerwheesht.jpg',
    },
    { name: 'broclee', color: '#3385ff', icon: '/static/broclee.jpg' },
  ];

  return (
    <Stack>
      <LineChart
        h={400}
        data={rankHistory}
        dataKey="date"
        series={series}
        xAxisLabel="Date"
        yAxisLabel="LP Gained"
        curveType="linear"
        lineChartProps={{
          style: {
            marginTop: 40,
          },
        }}
      />

      {/* Custom Legend */}
      <Group>
        {series.map((item) => (
          <Group key={item.name} gap="xs">
            <Image src={item.icon} width={80} height={80} alt={item.name} />
            <Text size="m" c={item.color}>
              {item.name}
            </Text>
          </Group>
        ))}
      </Group>
    </Stack>
  );
}
