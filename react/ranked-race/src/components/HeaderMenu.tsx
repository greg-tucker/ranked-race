import Image from 'next/image';
import mapImage from '../../public/static/img.png';

export function HeaderMenu() {
  return (
    <header style={{
      display: 'flex',
      flexDirection: 'row',
      alignItems: 'center',
      justifyContent: 'center',
      gap: 24,
      minHeight: 80,
    }}>
      <Image width={60} height={60} alt="Ranked Race Logo" src={mapImage} style={{ borderRadius: 12, boxShadow: '0 2px 12px #0ea5e9' }} />
      <div>
        <div className="rr-title">Ranked Race</div>
        <div style={{ color: '#a8b5d8', fontSize: 18, fontWeight: 400, marginTop: 2, letterSpacing: 1 }}>
          你好 你要回我的家吗
        </div>
      </div>
    </header>
  );
}
