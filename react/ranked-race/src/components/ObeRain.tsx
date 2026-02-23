import React, { useEffect, useMemo } from 'react';

type ObeRainProps = {
  active: boolean;
  count?: number;
  src?: string;
  onFinish?: () => void;
};

export const ObeRain: React.FC<ObeRainProps> = ({ active, count = 40, src = '/static/obe.png', onFinish }) => {
  const items = useMemo(() => {
    return Array.from({ length: count }).map(() => ({
      left: Math.random() * 100,
      delay: Math.random() * 2.5,
      duration: 6 + Math.random() * 6,
      spin: Math.random() * 360,
      scale: 0.6 + Math.random() * 0.9,
      topOffset: -10 - Math.random() * 20,
    }));
  }, [count]);

  useEffect(() => {
    if (!active) return;
    // keep rain + sparkles visible slightly longer than the longest animation
    const timeout = setTimeout(() => {
      onFinish && onFinish();
    }, 14000);
    return () => clearTimeout(timeout);
  }, [active, onFinish]);

  if (!active) return null;

  const sparkleVariants = ['twinkle', 'flare', 'confetti', 'star'];
  const sparkles = Array.from({ length: 60 }).map(() => {
    const variant = sparkleVariants[Math.floor(Math.random() * sparkleVariants.length)];
    return {
      left: Math.random() * 100,
      top: Math.random() * 80,
      delay: Math.random() * 3,
      duration: 2 + Math.random() * 4,
      size: 4 + Math.random() * 16,
      opacity: 0.25 + Math.random() * 0.8,
      rotate: Math.random() * 360,
      variant,
    };
  });

  return (
    <>
      <div className="obe-sparkle-overlay" aria-hidden>
        {sparkles.map((s, i) => (
          <span
            key={i}
            className={`obe-sparkle obe-sparkle--${s.variant}`}
            style={{
              left: `${s.left}%`,
              top: `${s.top}%`,
              width: `${s.size}px`,
              height: `${s.size}px`,
              transform: `rotate(${s.rotate}deg)`,
              animationDelay: `${s.delay}s`,
              animationDuration: `${s.duration}s`,
              opacity: s.opacity,
            }}
          />
        ))}
      </div>

      <div className="obe-rain-overlay" aria-hidden>
        {items.map((it, i) => (
          <img
            key={i}
            src={src}
            className="obe-rain-item"
            style={{
              left: `${it.left}%`,
              animationDelay: `${it.delay}s`,
              animationDuration: `${it.duration}s`,
              transform: `translateY(${it.topOffset}vh) rotate(${it.spin}deg) scale(${it.scale})`,
            }}
            alt="obe"
          />
        ))}
      </div>
    </>
  );
};

export default ObeRain;
