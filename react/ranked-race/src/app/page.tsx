'use client';

import AudioPlayer from '@/components/AudioPlayer';
import { HeaderMenu } from '@/components/HeaderMenu';
import { MainRankings } from '@/components/MainRankings';
import { Modal } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { useRef } from 'react';

export default function HomePage() {
  const [opened, { close }] = useDisclosure(true);
  const audioRef = useRef<HTMLAudioElement | null>(null);

  return (
    <>
      <audio loop ref={audioRef} id="gods">
        <source src="/audio/gods.mp3" type="audio/mpeg" />
        Gods
      </audio>

      <Modal
        opened={opened}
        onClose={close}
        withCloseButton={false}
        fullScreen
        centered
        transitionProps={{ transition: 'fade', duration: 200 }}
      >
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '90vh' }}>
          <AudioPlayer closeModal={close} audioRef={audioRef} text='ARE YOU READY???' />
        </div>
      </Modal>

      <div>
        <HeaderMenu />
        <main>
          <h3>Most Epic Ranking!</h3>
          <MainRankings />
          <AudioPlayer closeModal={close} audioRef={audioRef} text='Toggle music' />
        </main>
        <footer>Brought to you by Greg, Zi, Nikki</footer>
      </div>
    </>
  );
}
