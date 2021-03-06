package utils

import (
	"github.com/cryptogenic/goflexpool/pkg/api"
)

// CalculateExpectedRoundTime takes a given network hashrate, pool hashrate, and average block time, and calculates what
// the expected roundtime should be in seconds. The units for the network and pool hashrate do not matter as long as they're
// consistent and the same unit. Megahashes/second is recommended for reasonable accuracy.
func CalculateExpectedRoundTime(networkHashrate uint, poolHashrate uint, averageBlocktime int) uint {
	return networkHashrate / poolHashrate * uint(averageBlocktime)
}

// CalculatePPLNSShareWindow takes the N value, the pool's current share difficulty, and the pool's hashrate to calculate
// how many seconds it will take for those shares to be exhausted / expire. The units for the share difficulty and pool
// hashrate do not matter as long as they're consistent and the same unit. Megahashes/second is recommended for reasonable
// accuracy.
func CalculatePPLNSShareWindow(N int, shareDifficulty uint, poolHashrate uint) uint {
	return (uint(N) * shareDifficulty) / poolHashrate
}

// CalculateUncleRate takes a slice of api.Block instances and returns the uncle rate as a percentage as a float64.
func CalculateUncleRate(blocks []api.Block) float64 {
	uncleBlocks := float64(0)

	for _, block := range blocks {
		if block.Type == "uncle" {
			uncleBlocks++
		}
	}

	return uncleBlocks / float64(len(blocks))
}

// CalculateAverageBlockReward takes a slice of api.Block instances and calculates the average reward per block in gwei
// and returns it as an int.
func CalculateAverageBlockReward(blocks []api.Block) uint {
	blockRewardTotal := uint(0)

	for _, block := range blocks {
		blockRewardTotal += block.TotalRewards
	}

	return blockRewardTotal / uint(len(blocks))
}

// CalculateAverageBlocksPerDay takes a slice of api.Block instances and calculates the number of blocks found per day
// based on the roundtime of the blocks.
func CalculateAverageBlocksPerDay(blocks []api.Block) int {
	blockTimeTotal := 0

	for _, block := range blocks {
		blockTimeTotal += int(block.RoundTime)
	}

	// Seconds / 60 = Minutes / 60 = Hours / 24 = 1 Day
	blockTimeTotalDays := blockTimeTotal / 60 / 60 / 24
	return len(blocks) / blockTimeTotalDays
}
