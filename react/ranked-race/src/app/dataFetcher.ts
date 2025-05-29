import { Dispatch, SetStateAction } from 'react';
import { MainRankingsData } from '@/app/dataTypes';
import { useEffect } from 'react';

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

export async function getRankHistory(
  setter: Dispatch<SetStateAction<MainRankingsData[]>>
) {
  const url = `${process.env.NEXT_PUBLIC_RR_BACKEND_URL}ranktimeline`;
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