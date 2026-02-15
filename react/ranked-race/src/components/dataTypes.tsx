import Image from 'next/image';

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
};

export type ColumnKey = keyof MainRankingsData | 'opgg';
// @ts-ignore
const loaderProp = ({ src }) => {
  return src;
};
export const visibleColumns: {
  key: ColumnKey;
  label: string;
  render?: (row: MainRankingsData) => React.ReactNode;
}[] = [
  { key: 'name', label: 'Name' },
  { key: 'role', label: 'Role'},
  {
    key: 'displayRank',
    label: 'Rank',
    render: (row) => {
      if (row.tier) {
        const source = `/static/${row.tier.toLowerCase()}.png`;
        return (
          <div>
            <Image
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
    label: 'Total Games'
  }
];
