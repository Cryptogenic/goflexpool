package main

import (
	"flag"
	"fmt"
	"github.com/cryptogenic/goflexpool/pkg/api"
	"github.com/cryptogenic/goflexpool/pkg/utils"
	"os"
	"time"
)

func main() {
	var (
		err                error
		balanceGwei        uint
		metaDetails        api.MinerDetails
		roundShare         float64
		dailyEstimatedGwei uint
		totalPaidGwei      uint
		totalDonatedGwei   uint
		workers            []api.MinerWorker
		paymentData        api.MinerPaymentData
		blockData          api.MinerBlockData

		minerAddress string
	)

	// Take an address to check from argument
	flag.StringVar(&minerAddress, "address", "", "Mining wallet address")
	flag.Parse()

	if minerAddress == "" {
		fmt.Printf("No address given, exiting.\n")
		os.Exit(1)
	}

	// Get balance
	if balanceGwei, err = api.MinerGetBalance(minerAddress); err != nil {
		fmt.Printf("Unable to get wallet balance: %v\n", err.Error())
		os.Exit(1)
	}

	balanceEth := utils.ConvertGweiToEth(balanceGwei)

	// Get meta details
	if metaDetails, err = api.MinerGetDetails(minerAddress); err != nil {
		fmt.Printf("Unable to get wallet details: %v\n", err.Error())
		os.Exit(1)
	}

	// Get round share
	if roundShare, err = api.MinerGetRoundShare(minerAddress); err != nil {
		fmt.Printf("Unable to get round share: %v\n", err.Error())
		os.Exit(1)
	}

	// Get estimated daily eth
	if dailyEstimatedGwei, err = api.MinerGetEstimatedDailyRevenue(minerAddress); err != nil {
		fmt.Printf("Unable to get estimated daily revenue: %v\n", err.Error())
		os.Exit(1)
	}

	dailyEstimatedEth := dailyEstimatedGwei

	// Get total paid
	if totalPaidGwei, err = api.MinerGetTotalPaid(minerAddress); err != nil {
		fmt.Printf("Unable to get total paid: %v\n", err.Error())
		os.Exit(1)
	}

	totalPaidEth := totalPaidGwei

	// Get total donate
	if totalDonatedGwei, err = api.MinerGetTotalDonated(minerAddress); err != nil {
		fmt.Printf("Unable to get total donated: %v\n", err.Error())
		os.Exit(1)
	}

	totalDonatedEth := totalDonatedGwei

	// Get workers
	if workers, err = api.MinerGetWorkers(minerAddress); err != nil {
		fmt.Printf("Unable to get worker listing: %v\n", err.Error())
		os.Exit(1)
	}

	// Get payments
	if paymentData, err = api.MinerGetPayments(minerAddress, 0); err != nil {
		fmt.Printf("Unable to get last 10 payments: %v\n", err.Error())
		os.Exit(1)
	}

	// Get mined blocks
	if blockData, err = api.MinerGetBlocks(minerAddress, 0); err != nil {
		fmt.Printf("Unable to get last 10 blocks mined: %v\n", err.Error())
		os.Exit(1)
	}

	// Do pretty printing
	fmt.Printf("Flexpool Miner '%s' Stats\n-\n\n", minerAddress)
	fmt.Printf("Unpaid Balance: %.8f\n", balanceEth)

	fmt.Printf("Min Payout Threshold: %.4f eth \t\t Donation Percent: %.4f% \t Round Share: %.8f%\n",
		utils.ConvertGweiToEth(metaDetails.MinPayoutThreshold),
		metaDetails.PoolDonation,
		roundShare)

	fmt.Printf("Estimated Daily Eth: %.8f eth \t Total Paid: %.8f eth \t Total Donated: %.8f eth\n\n",
		utils.ConvertGweiToEth(dailyEstimatedEth),
		utils.ConvertGweiToEth(totalPaidEth),
		utils.ConvertGweiToEth(totalDonatedEth))

	fmt.Printf("Workers:\n")

	if len(workers) > 0 {
		for _, worker := range workers {
			fmt.Printf("\t %s (effective hashrate: %dMH/s) \t (valid: %d, stale: %d, invalid: %d)\n",
				worker.Name,
				utils.ConvertHashrate(worker.EffectiveHashrate, utils.HashesPerSecond, utils.MegaHashesPerSecond),
				worker.ValidShares,
				worker.InvalidShares,
				worker.InvalidShares)
		}
	} else {
		fmt.Printf("\t None currently active.\n")
	}

	fmt.Printf("\nLast 10 payments:\n")

	if paymentData.Data != nil {
		for _, payment := range paymentData.Data {
			fmt.Printf("\t Txn: %s (amount: %.8f eth) \t %s\n",
				payment.Txid,
				utils.ConvertGweiToEth(payment.Amount),
				time.Unix(int64(payment.Timestamp), 0))
		}
	} else {
		fmt.Printf("\t No payments made.\n")
	}

	fmt.Printf("\nLast 10 blocks mined: \n")

	if blockData.Data != nil {
		for _, block := range blockData.Data {
			fmt.Printf("\t %d (type: %s) (reward: %.8f) \t %s\n",
				block.Number,
				block.Type,
				utils.ConvertGweiToEth(block.TotalRewards),
				time.Unix(int64(block.Timestamp), 0))
		}
	} else {
		fmt.Printf("\t No blocks mined yet.\n")
	}
}
