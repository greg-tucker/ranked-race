import Image from 'next/image';
import mapImage from '../../public/static/img.png';

export function HeaderMenu() {
  return (
    <header
      style={{
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center',
      }}
    >
      <Image width={50} height={50} alt="map icon" src={mapImage} />
      <div style={{ marginLeft: 10 }} className="rr-title">
        Welcome to Ranked Race
      </div>
    </header>
  );
}
