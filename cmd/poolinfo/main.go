package main

import (
	"fmt"
	"os"
	"time"

	//"time"

	"github.com/cryptogenic/goflexpool/pkg/api"
	"github.com/cryptogenic/goflexpool/pkg/utils"
)

const (
	FlexpoolCurrentN               = 2_000_000     // 2 million share pool
	FlexpoolCurrentShareDifficulty = 4_000_000_000 // 4GH/s difficulty
)

func secondsToHhMmSs(secondsIn uint) string {
	hours := secondsIn / 60 / 60
	minutes := (secondsIn / 60) % 60
	seconds := secondsIn % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func main() {
	var (
		err             error
		poolHashrate    api.PoolHashrate
		poolMinerCount  int
		poolWorkerCount int
	)

	// Get basic pool info (hashrates, miners and workers online)
	if poolHashrate, err = api.PoolGetHashrate(); err != nil {
		fmt.Printf("Failed to get pool hashrate: %v\n", err.Error())
		os.Exit(1)
	}

	if poolMinerCount, err = api.PoolGetMinersOnline(); err != nil {
		fmt.Printf("Failed to get online miner count: %v\n", err.Error())
		os.Exit(1)
	}

	if poolWorkerCount, err = api.PoolGetWorkersOnline(); err != nil {
		fmt.Printf("Failed to get online worker count: %v\n", err.Error())
		os.Exit(1)
	}

	// Get last 10 pages of blocks
	var pageBlockData api.PoolBlockData
	var blocks []api.Block

	for i := 0; i < 10; i++ {
		if pageBlockData, err = api.PoolGetBlocks(i); err != nil {
			fmt.Printf("Failed to get pool blockdata: %v\n", err.Error())
			os.Exit(1)
		}

		for _, block := range pageBlockData.Data {
			blocks = append(blocks, block)
		}

		// Lighten the load on flexpool API
		time.Sleep(200 * time.Millisecond)
	}

	// Convert hashrates to GH/s for display
	poolHashrateTotalGHs := utils.ConvertHashrate(poolHashrate.Total, utils.HashesPerSecond, utils.GigaHashesPerSecond)
	poolHashrateAsGHs := utils.ConvertHashrate(poolHashrate.As, utils.HashesPerSecond, utils.GigaHashesPerSecond)
	poolHashrateAuGHs := utils.ConvertHashrate(poolHashrate.Au, utils.HashesPerSecond, utils.GigaHashesPerSecond)
	poolHashrateEuGHs := utils.ConvertHashrate(poolHashrate.Eu, utils.HashesPerSecond, utils.GigaHashesPerSecond)
	poolHashrateSaGHs := utils.ConvertHashrate(poolHashrate.Sa, utils.HashesPerSecond, utils.GigaHashesPerSecond)
	poolHashrateUsGHs := utils.ConvertHashrate(poolHashrate.Us, utils.HashesPerSecond, utils.GigaHashesPerSecond)

	// Get PPLNS share window, uncle rate, average block reward, and average blocks per day
	pplnsShareWindowSeconds := utils.CalculatePPLNSShareWindow(FlexpoolCurrentN, FlexpoolCurrentShareDifficulty, poolHashrate.Total)
	uncleRate := utils.CalculateUncleRate(blocks)
	averageBlockRewardGwei := utils.CalculateAverageBlockReward(blocks)
	averageBlocksPerDay := utils.CalculateAverageBlocksPerDay(blocks)

	averageBlockRewardEth := utils.ConvertGweiToEth(averageBlockRewardGwei)

	// Do pretty printing
	fmt.Printf("Flexpool Stats\n-\n\n")
	fmt.Printf("Miners: %d (Workers: %d)\n\n", poolMinerCount, poolWorkerCount)
	fmt.Printf("Hashrate: %vGH/s\n", poolHashrateTotalGHs)
	fmt.Printf("\tAs: %vGH/s\n", poolHashrateAsGHs)
	fmt.Printf("\tAu: %vGH/s\n", poolHashrateAuGHs)
	fmt.Printf("\tEu: %vGH/s\n", poolHashrateEuGHs)
	fmt.Printf("\tSa: %vGH/s\n", poolHashrateSaGHs)
	fmt.Printf("\tUs: %vGH/s\n\n", poolHashrateUsGHs)

	fmt.Printf("PPLNS share window: %s (%d)\n", secondsToHhMmSs(pplnsShareWindowSeconds), pplnsShareWindowSeconds)
	fmt.Printf("Uncle rate: %.2f%%\n", uncleRate*100)
	fmt.Printf("Average blocks per day: %d (average reward: %.8f)\n", averageBlocksPerDay, averageBlockRewardEth)
	fmt.Printf("\t* Averages and uncle rate are over a 100 block period")
}
