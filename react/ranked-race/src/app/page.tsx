'use client';

import { HeaderMenu } from '@/components/HeaderMenu';
import { MainRankings } from '@/components/MainRankings';

export default function HomePage() {
  return (
    <>
      <HeaderMenu />
      <main>
        <h3>Most Epic Ranking!</h3>
        <MainRankings />
      </main>
      <footer>Brought to you by Greg, Zi, Nikki</footer>
    </>
  );
}
