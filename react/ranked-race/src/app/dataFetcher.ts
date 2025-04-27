import { Dispatch, SetStateAction } from 'react';
import { MainRankingsData } from '@/app/dataTypes';

export async function getCurrentRanking(
  setter: Dispatch<SetStateAction<MainRankingsData[]>>
) {
  const url = `${process.env.NEXT_PUBLIC_RR_BACKEND_URL}rank`;
  try {
    let response = await fetch(url, {
      credentials: 'same-origin',
    });
    if (!response.ok) {
      throw new Error('Failed to fetch data');
    }
    let json = await response.json();

    setter(json);
  } catch (error) {
    console.error(error);
  }
}

export async function getRankHistoryForPlayer(
  player: string,
  setter: Dispatch<SetStateAction<MainRankingsData[]>>
) {
  const url = `${process.env.NEXT_PUBLIC_RR_BACKEND_URL}ranktimeline/${player}`;
  try {
    let response = await fetch(url, {
      credentials: 'same-origin',
    });
    if (!response.ok) {
      throw new Error('Failed to fetch data');
    }
    let json = await response.json();

    console.log("retrieved mayalover3 data")
    console.log(json)
    setter(json);
  } catch (error) {
    console.error(error);
  }
}
