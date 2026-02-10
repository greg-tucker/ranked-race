import { Button } from '@mantine/core';
import React from 'react';

interface AudioPlayerProps {
  closeModal: () => void;
  audioRef: React.RefObject<HTMLAudioElement>;
  text: string;
}

export default function AudioPlayer({
  closeModal,
  audioRef,
  text,
}: AudioPlayerProps) {
  const playSong = () => {
    if (audioRef.current) {
      audioRef.current.volume = 0.2;
      if (audioRef.current.paused) {
        audioRef.current
          .play()
          .catch((err) => console.error('Play failed:', err)); // Handle play errors
      } else {
        audioRef.current.pause();
      }
    } else {
      console.error('audioRef is null');
    }

    closeModal();
  };

  return (
    <>
      <Button onClick={playSong} mx="auto" mt="xl">
        {text}
      </Button>
    </>
  );
}
