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
  displayRank: string
  winRate: string
};

export const visibleColumns :(keyof MainRankingsData)[] = [
  'name',
  'displayRank',
  'winRate',
  'wins',
  'losses',
];
