'use client';

import { HeaderMenu } from '@/components/HeaderMenu';
import { NavTabs } from '@/components/NavTabs';

export default function HomePage() {
  return (
    <>
      <link rel="icon" href="static/img.png" sizes="any" />
      <div>
        <HeaderMenu />
        <main>
          <div className="glass" style={{ padding: '2.5rem 2rem', margin: '2rem auto', maxWidth: 1200, boxShadow: '0 4px 32px #0ea5e933' }}>
            <NavTabs />
          </div>
        </main>
        <footer style={{ width:'100%', position:'relative'}}>Brought to you by Greg, Zi, Nikki</footer>
      </div>
    </>
  );
}
