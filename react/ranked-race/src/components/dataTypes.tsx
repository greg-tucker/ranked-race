import Image from 'next/image';
import { GameTimer } from './GameTimer';

export type LiveGamePerks = {
  perkIds: number[];
  perkStyle: number;
  perkSubStyle: number;
};

export type LiveGameParticipant = {
  championId: number;
  perks: LiveGamePerks;
  profileIconId: number;
  bot: boolean;
  teamId: number;
  summonerName: string;
  summonerId: string;
  puuid: string;
  spell1Id: number;
  spell2Id: number;
  gameCustomizationObjects: any[];
};

export type LiveGameObservers = {
  encryptionKey: string;
};

export type LiveGame = {
  gameId: number;
  gameType: string;
  gameStartTime: number;
  mapId: number;
  gameLength: number;
  platformId: string;
  gameMode: string;
  bannedChampions: any[];
  gameQueueConfigId: number;
  observers: LiveGameObservers;
  participants: LiveGameParticipant[];
};

export type MainRankingsData = {
  name: string;
  peak: number;
  gained: number;
  current: number;
  wins: number;
  losses: number;
  played: number;
  tier: string;
  rank: string;
  lp: number;
  date: string;
  displayRank: string;
  winRate: number;
  tag: string;
  opgg: string;
  inGame: boolean;
  role: string;
  startTime?: number;
  puuid:string;
};

export type ColumnKey = keyof MainRankingsData | 'opgg';
// @ts-ignore
export const loaderProp = ({ src }) => {
  return src;
};
export const visibleColumns: {
  key: ColumnKey;
  label: string;
  render?: (row: MainRankingsData) => React.ReactNode;
}[] = [
  { key: 'name', label: 'Name' },
  { key:
     'role',
      label: 'Role',
    render: (row) => {  
      return <Image src={`/static/${row.role.toLowerCase()}_icon.webp`} loader={loaderProp} alt={row.role} width={32} height={32} style={{ borderRadius: 6 }} />
  }},
  {
    key: 'displayRank',
    label: 'Rank',
    render: (row) => {
      if (row.tier) {
        const source = `/static/${row.tier.toLowerCase()}.png`;
        return (
          <div style={{display: 'flex', alignItems: 'center'}}>
            <Image
              style={{marginRight:'1rem'}}
              src={source}
              alt=""
              loader={loaderProp}
              width={50}
              height={50}
            />
            {row.displayRank}
          </div>
        );
      }
      return <div> This loser is unranked </div>;
    },
  },
  {
    key: 'opgg',
    label: 'OP.GG',
    render: (row) => (
      <a
        href={`https://www.op.gg/summoners/euw/${encodeURIComponent(row.name)}%20-${row.tag}`}
        target="_blank"
        rel="noopener noreferrer"
      >
        <Image src="/static/opgg.svg" alt="OP.GG" width="50" height="20" />
      </a>
    ),
  },
  {
    key: 'winRate',
    label: 'Win Rate',
    render: (row) => (
        <div className={row.winRate > 50 ? 'green' : 'red'}>{row.winRate}%</div>
    ),
  },
  {
    key: 'wins',
    label: 'Wins',
  },
  {
    key: 'losses',
    label: 'Losses',
  },
  {
    key: 'played',
    label: 'Total Games',
    render: (row) => {
      return (
        <div style={{display:'flex', flexDirection:'row', justifyContent:'space-between'}}><span>{row.played}</span> {row.inGame && row.startTime && <GameTimer startTime={row.startTime} />}</div>
      );
    }
  }
];
