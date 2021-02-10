package main

import (
	"fmt"
	"testing"

	"../pkg/api"
)

func TestMinerEndpoints(t *testing.T) {
	// README: edit this address to whatever address you want to test with
	const ADDR = ""

	t.Run("MinerGetBalance", func(t *testing.T) {
		if balance, err := api.MinerGetBalance(ADDR); err == nil {
			fmt.Printf("Get Balance: %+v\n", balance)
		} else {
			t.Errorf("MinerGetBalance failed with: %v", err)
		}
	})

	t.Run("MinerGetCurrent", func(t *testing.T) {
		if currentStats, err := api.MinerGetCurrent(ADDR); err == nil {
			fmt.Printf("Get Current: %+v\n", currentStats)
		} else {
			t.Errorf("MinerGetCurrent failed with: %v", err)
		}
	})

	t.Run("MinerGetDaily", func(t *testing.T) {
		if dailyStats, err := api.MinerGetDaily(ADDR); err == nil {
			fmt.Printf("Get Daily: %+v\n", dailyStats)
		} else {
			t.Errorf("MinerGetDaily failed with: %v", err)
		}
	})

	t.Run("MinerGetStats", func(t *testing.T) {
		if stats, err := api.MinerGetStats(ADDR); err == nil {
			fmt.Printf("Get Stats: %+v\n", stats)
		} else {
			t.Errorf("MinerGetStats failed with: %v", err)
		}
	})

	t.Run("MinerGetWorkerCount", func(t *testing.T) {
		if workerCount, err := api.MinerGetWorkerCount(ADDR); err == nil {
			fmt.Printf("Get Worker Count: %+v\n", workerCount)
		} else {
			t.Errorf("MinerGetWorkerCount failed with: %v", err)
		}
	})

	t.Run("MinerGetWorkers", func(t *testing.T) {
		if workers, err := api.MinerGetWorkers(ADDR); err == nil {
			fmt.Printf("Get Workers: %+v\n", workers)
		} else {
			t.Errorf("MinerGetWorkers failed with: %v", err)
		}
	})

	t.Run("MinerGetChart", func(t *testing.T) {
		if chartData, err := api.MinerGetChart(ADDR); err == nil {
			fmt.Printf("Get Chart: %+v\n", chartData)
		} else {
			t.Errorf("MinerGetChart failed with: %v", err)
		}
	})

	t.Run("MinerGetPayments", func(t *testing.T) {
		if paymentData, err := api.MinerGetPayments(ADDR, 0); err == nil {
			fmt.Printf("Get Payments: %+v\n", paymentData)
		} else {
			t.Errorf("MinerGetPayments failed with: %v", err)
		}
	})

	t.Run("MinerGetPaymentCount", func(t *testing.T) {
		if paymentCount, err := api.MinerGetPaymentCount(ADDR); err == nil {
			fmt.Printf("Get Payment Count: %+v\n", paymentCount)
		} else {
			t.Errorf("MinerGetPaymentCount failed with: %v", err)
		}
	})

	t.Run("MinerGetPaymentChart", func(t *testing.T) {
		if paymentChartData, err := api.MinerGetPaymentChart(ADDR); err == nil {
			fmt.Printf("Get Payment Chart Data: %+v\n", paymentChartData)
		} else {
			t.Errorf("MinerGetPaymentChart failed with: %v", err)
		}
	})

	t.Run("MinerGetBlocks", func(t *testing.T) {
		if blockData, err := api.MinerGetBlocks(ADDR, 0); err == nil {
			fmt.Printf("Get Blocks: %+v\n", blockData)
		} else {
			t.Errorf("MinerGetBlocks failed with: %+v", err)
		}
	})

	t.Run("MinerGetBlockCount", func(t *testing.T) {
		if blockCount, err := api.MinerGetBlockCount(ADDR); err == nil {
			fmt.Printf("Get Block Count: %+v\n", blockCount)
		} else {
			t.Errorf("MinerGetBlockCount failed with: %+v", err)
		}
	})

	t.Run("MinerGetDetails", func(t *testing.T) {
		if details, err := api.MinerGetDetails(ADDR); err == nil {
			fmt.Printf("Get Details: %v\n", details)
		} else {
			t.Errorf("MinerGetDetails failed with: %+v", err)
		}
	})

	t.Run("MinerGetEstimatedDailyRevenue", func(t *testing.T) {
		if estimatedDailyRevenue, err := api.MinerGetEstimatedDailyRevenue(ADDR); err == nil {
			fmt.Printf("Get Estimated Daily Revenue: %+v\n", estimatedDailyRevenue)
		} else {
			t.Errorf("MinerGetEstimatedDailyRevenue failed with: %v", err)
		}
	})

	t.Run("MinerGetRoundShare", func(t *testing.T) {
		if roundShare, err := api.MinerGetRoundShare(ADDR); err == nil {
			fmt.Printf("Get Round Share: %v\n", roundShare)
		} else {
			t.Errorf("MinerGetRoundShare failed with: %+v", err)
		}
	})

	t.Run("MinerGetTotalPaid", func(t *testing.T) {
		if totalPaid, err := api.MinerGetTotalPaid(ADDR); err == nil {
			fmt.Printf("Get Total Paid: %v\n", totalPaid)
		} else {
			t.Errorf("MinerGetTotalPaid failed with: %+v", err)
		}
	})

	t.Run("MinerGetTotalDonated", func(t *testing.T) {
		if totalDonated, err := api.MinerGetTotalDonated(ADDR); err == nil {
			fmt.Printf("Get Total Donated: %v\n", totalDonated)
		} else {
			t.Errorf("MinerGetTotalDonated failed with: %+v", err)
		}
	})
}
