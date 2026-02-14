'use client';

import { HeaderMenu } from '@/components/HeaderMenu';
import { NavTabs } from '@/components/NavTabs';
import { useDisclosure } from '@mantine/hooks';
import { useRef } from 'react';

export default function HomePage() {
  const [opened, { close }] = useDisclosure(true);
  const audioRef = useRef<HTMLAudioElement>(null);

  return (
    <>
      <audio loop ref={audioRef} id="gods">
        <source src="/audio/gods.mp3" type="audio/mpeg" />
        Gods
      </audio>

      <div>
        <HeaderMenu />
        <main>
          <div className="glass" style={{ padding: '2.5rem 2rem', margin: '2rem auto', maxWidth: 1200, boxShadow: '0 4px 32px #0ea5e933' }}>
            <h3 style={{ textAlign: 'center', fontSize: 28, fontWeight: 600, marginBottom: 24, letterSpacing: 1 }}>Most Epic Ranking!</h3>
            <NavTabs />
          </div>
        </main>
        <footer>Brought to you by Greg, Zi, Nikki</footer>
      </div>
    </>
  );
}
