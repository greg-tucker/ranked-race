'use client';
import React, { useEffect, useRef } from 'react';

type Props = {
  active: boolean;
  onFinish?: () => void;
  carSrc?: string;
  driverSrc?: string;
  containerWidth?: number;
  durationMs?: number;
  extraPx?: number;
};

export default function CarDrive({ active, onFinish, carSrc, driverSrc, containerWidth = 1100, durationMs = 2000, extraPx = 0 }: Props) {
  const wrapperRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    if (!active) return;
    const el = wrapperRef.current;
    const handle = () => onFinish && onFinish();
    if (el) el.addEventListener('animationend', handle);
    return () => { if (el) el.removeEventListener('animationend', handle); };
  }, [active, onFinish]);

  if (!active) return null;

  const carImage = carSrc ?? '/static/car.png';
  const driverImage = driverSrc ?? '/static/joe2.png';

  return (
    <div
      ref={wrapperRef}
      className="crochecha-car-wrapper"
      style={{ pointerEvents: 'none', ['--cw' as any]: `${containerWidth}px`, ['--dur' as any]: `${durationMs}ms`, ['--extra' as any]: `${extraPx}px` } as React.CSSProperties}
    >
      <div className="crochecha-car">
        <img src={carImage} alt="car" className="car-img" />
        <img src={driverImage} alt="driver" className="driver-img" />
      </div>

      <style jsx>{`
        .crochecha-car-wrapper {
          position: fixed;
          left: 50%;
          transform: translateX(-50%);
          top: 30%;
          width: min(var(--cw, 1100px), 100vw);
          height: 160px;
          z-index: 9999;
          overflow: hidden;
        }

        .crochecha-car {
          position: absolute;
          right: calc(-260px - var(--extra));
          display: flex;
          align-items: center;
          /* animate right so the car travels the full width of the container */
          animation: drive-across var(--dur) linear forwards;
        }

        .car-img {
          width: 240px;
          height: 120px;
          display: block;
          object-fit: cover;
          position: relative;
          z-index: 1;
          border-radius: 18px;
          filter: drop-shadow(0 8px 18px rgba(0,0,0,0.35));
        }

        .driver-img {
          position: absolute;
          width: 72px;
          height: 72px;
          border-radius: 50%;
          border: 3px solid rgba(255,255,255,0.9);
          object-fit: cover;
          left: 50px;
          top: 0px;
          z-index: 3;
        }

        @keyframes drive-across {
          0% { right: calc(-260px - var(--extra)); opacity: 1; }
          85% { opacity: 1; }
          /* end with the car fully past the left edge of the container */
          100% { right: calc(100% + 260px + var(--extra)); opacity: 0; }
        }
      `}</style>
    </div>
  );
}
