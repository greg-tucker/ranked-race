'use client';
import { useEffect, useState } from 'react';

interface GameTimerProps {
  startTime: number; // Unix timestamp in milliseconds
}

export function GameTimer({ startTime }: GameTimerProps) {
  const [elapsed, setElapsed] = useState(0);

  useEffect(() => {
    // Calculate initial elapsed time
    const now = Date.now();
    const initialElapsed = Math.max(0, Math.floor((now - startTime) / 1000));
    setElapsed(initialElapsed);

    // Set up interval to update every second
    const interval = setInterval(() => {
      const currentTime = Date.now();
      const newElapsed = Math.max(0, Math.floor((currentTime - startTime) / 1000));
      setElapsed(newElapsed);
    }, 1000);

    return () => clearInterval(interval);
  }, [startTime]);

  const minutes = Math.floor(elapsed / 60);
  const seconds = elapsed % 60;

  const formatTime = (num: number) => num.toString().padStart(2, '0');

  return <span>{formatTime(minutes)}:{formatTime(seconds)}</span>;
}
