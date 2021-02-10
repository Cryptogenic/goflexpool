package main

import (
	"fmt"
	"testing"

	"../pkg/api"
)

func TestPoolEndpoints(t *testing.T) {
	t.Run("PoolGetHashrate", func(t *testing.T) {
		if hashrate, err := api.PoolGetHashrate(); err == nil {
			fmt.Printf("Get Hashrate: %+v\n", hashrate)
		} else {
			t.Errorf("PoolGetHashrate failed with: %v", err)
		}
	})

	t.Run("PoolGetHashrateChart", func(t *testing.T) {
		if hashrate, err := api.PoolGetHashrateChart(); err == nil {
			fmt.Printf("Get Hashrate Chart: %+v\n", hashrate)
		} else {
			t.Errorf("PoolGetHashrateChart failed with: %v", err)
		}
	})

	t.Run("PoolGetMinersOnline", func(t *testing.T) {
		if miners, err := api.PoolGetMinersOnline(); err == nil {
			fmt.Printf("Get Miners Online: %+v\n", miners)
		} else {
			t.Errorf("PoolGetMinersOnline failed with: %v", err)
		}
	})

	t.Run("PoolGetWorkersOnline", func(t *testing.T) {
		if workers, err := api.PoolGetWorkersOnline(); err == nil {
			fmt.Printf("Get Workers Online: %+v\n", workers)
		} else {
			t.Errorf("PoolGetWorkersOnline failed with: %v", err)
		}
	})

	t.Run("PoolGetBlocks", func(t *testing.T) {
		if blocks, err := api.PoolGetBlocks(0); err == nil {
			fmt.Printf("Get Pool Blocks: %+v\n", blocks)
		} else {
			t.Errorf("PoolGetBlocks failed with: %v", err)
		}
	})

	t.Run("PoolGetBlockCount", func(t *testing.T) {
		if blockCount, err := api.PoolGetBlockCount(); err == nil {
			fmt.Printf("Get Block Count: %+v\n", blockCount)
		} else {
			t.Errorf("PoolGetBlockCount failed with: %v", err)
		}
	})

	t.Run("PoolGetTopMiners", func(t *testing.T) {
		if topMiners, err := api.PoolGetTopMiners(); err == nil {
			fmt.Printf("Get Top Miners: %+v\n", topMiners)
		} else {
			t.Errorf("PoolGetTopMiners failed with: %v", err)
		}
	})

	t.Run("PoolGetTopDonators", func(t *testing.T) {
		if topDonators, err := api.PoolGetTopDonators(); err == nil {
			fmt.Printf("Get Top Donators: %+v\n", topDonators)
		} else {
			t.Errorf("PoolGetTopDonators failed with: %v", err)
		}
	})

	t.Run("PoolGetAverageLuckRoundTime", func(t *testing.T) {
		if avgLuckRoundTime, err := api.PoolGetAverageLuckRoundTime(); err == nil {
			fmt.Printf("Get Average Luck Round Time: %+v\n", avgLuckRoundTime)
		} else {
			t.Errorf("PoolGetAverageLuckRoundTime failed with: %v", err)
		}
	})

	t.Run("PoolGetCurrentLuck", func(t *testing.T) {
		if luck, err := api.PoolGetCurrentLuck(); err == nil {
			fmt.Printf("Get Current Luck: %+v\n", luck)
		} else {
			t.Errorf("PoolGetCurrentLuck failed with: %v", err)
		}
	})

	t.Run("PoolGetAverageBlockReward", func(t *testing.T) {
		if averageBlockReward, err := api.PoolGetAverageBlockReward(); err == nil {
			fmt.Printf("Get Average Block Reward: %+v\n", averageBlockReward)
		} else {
			t.Errorf("PoolGetAverageBlockReward failed with: %v", err)
		}
	})
}
