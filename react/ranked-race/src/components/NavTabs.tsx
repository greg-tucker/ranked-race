import { Tabs } from '@mantine/core';
import { MainRankings } from '@/components/MainRankings';
import { Graph } from '@/components/Graph';

export function NavTabs() {
  return (
    <Tabs defaultValue="table">
      <Tabs.List>
        <Tabs.Tab value="table">Table</Tabs.Tab>
        <Tabs.Tab value="timeline">Timeline</Tabs.Tab>
      </Tabs.List>

      <Tabs.Panel value="table">
        <MainRankings />
      </Tabs.Panel>
      <Tabs.Panel value="timeline">
        <Graph />
      </Tabs.Panel>
    </Tabs>
  );
}
