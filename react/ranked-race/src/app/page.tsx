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
          <h3>Most Epic Ranking!</h3>
          <NavTabs />
          {/* <AudioPlayer closeModal={close} audioRef={audioRef.current} text='Toggle music' /> */}
        </main>

        <footer>Brought to you by Greg, Zi, Nikki</footer>
      </div>
    </>
  );
}
