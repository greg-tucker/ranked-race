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
};

export const visibleColumns :(keyof MainRankingsData)[] = [
  'name',
  'peak',
  'gained',
  'current',
  'wins',
  'losses',
];
