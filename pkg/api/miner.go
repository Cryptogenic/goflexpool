package api

import (
	"strconv"
)

// Block contains information relevant to blocks mined - used by multiple endpoints.
type Block struct {
	Hash                  string  `json:"hash"`
	Number                uint    `json:"number"`
	Type                  string  `json:"type"`
	Miner                 string  `json:"miner"`
	Difficulty            uint    `json:"difficulty"`
	Timestamp             uint    `json:"timestamp"`
	Confirmed             bool    `json:"confirmed"`
	RoundTime             uint    `json:"round_time"`
	Luck                  float64 `json:"luck"`
	ServerName            string  `json:"server_name"`
	BlockReward           uint    `json:"block_reward"`
	BlockFees             uint    `json:"block_fees"`
	UncleInclusionRewards uint    `json:"uncle_inclusion_rewards"`
	TotalRewards          uint    `json:"total_rewards"`
}

// MinerDailyStats contains miner daily stats data from the /miner/{address}/stats and /miner/{address}/daily endpoint.
type MinerDailyStats struct {
	EffectiveHashrate float64 `json:"effective_hashrate"`
	InvalidShares     int     `json:"invalid_shares"`
	ReportedHashrate  float64 `json:"reported_hashrate"`
	StaleShares       int     `json:"stale_shares"`
	ValidShares       int     `json:"valid_shares"`
}

// MinerDailyStats contains miner stats data from the /miner/{address}/stats endpoint.
type MinerStats struct {
	Current WorkerCurrentStats `json:"current"`
	Daily   MinerDailyStats    `json:"daily"`
}

// MinerWorkerCount contains worker count data from the /miner/{address}/workerCount endpoint.
type MinerWorkerCount struct {
	Online  int `json:"online"`
	Offline int `json:"offline"`
}

// MinerWorker contains worker data entries from the /miner/{address}/workers endpoint.
type MinerWorker struct {
	Name                   string `json:"name"`
	Online                 bool   `json:"online"`
	DuplicateWorkersMerged int    `json:"duplicate_workers_merged"`
	ReportedHashrate       uint   `json:"reported_hashrate"`
	EffectiveHashrate      uint   `json:"effective_hashrate"`
	ValidShares            int    `json:"valid_shares"`
	StaleShares            int    `json:"stale_shares"`
	InvalidShares          int    `json:"invalid_shares"`
	LastSeen               int    `json:"last_seen"`
}

// MinerChartData contains chart data entries from the /miner/{address}/chart endpoint.
type MinerChartData struct {
	Timestamp                int     `json:"timestamp"`
	EffectiveHashrate        uint    `json:"effective_hashrate"`
	AverageEffectiveHashrate float64 `json:"average_effective_hashrate"`
	ReportedHashrate         uint    `json:"reported_hashrate"`
	ValidShares              int     `json:"valid_shares"`
	StaleShares              int     `json:"stale_shares"`
	InvalidShares            int     `json:"invalid_shares"`
}

// MinerPayment contains payment entries from the /miner/{address}/payments endpoint.
type MinerPayment struct {
	Txid      string `json:"txid"`
	Amount    uint   `json:"amount"`
	Timestamp uint   `json:"timestamp"`
	Duration  uint   `json:"duration"`
}

// MinerPaymentData contains paged payment data from the /miner/{address}/payments endpoint.
type MinerPaymentData struct {
	Data         []MinerPayment `json:"data"`
	ItemsPerPage int            `json:"items_per_page"`
	TotalItems   int            `json:"total_items"`
	TotalPages   int            `json:"total_pages"`
}

// MinerPaymentChart contains payment chart data from the /miner/{address}/paymentsChart endpoint.
type MinerPaymentChart struct {
	Amount    uint `json:"amount"`
	Timestamp uint `json:"timestamp"`
}

// MinerBlockData contains paged block data from the /miner/{address}/blocks endpoint.
type MinerBlockData struct {
	Data         []Block `json:"data"`
	ItemsPerPage int     `json:"items_per_page"`
	TotalItems   int     `json:"total_items"`
	TotalPages   int     `json:"total_pages"`
}

// MinerBlockCount contains block count data from the /miner/{address}/blockCount endpoint.
type MinerBlockCount struct {
	Confirmed   int `json:"confirmed"`
	Unconfirmed int `json:"unconfirmed"`
}

// MinerDetails contains overview data from the /miner/{address}/details endpoint.
type MinerDetails struct {
	MinPayoutThreshold uint    `json:"min_payout_threshold"`
	PoolDonation       float64 `json:"pool_donation"`
	MaxFeePrice        uint    `json:"max_free_price"`
	CensoredEmail      string  `json:"censored_email"`
	CensoredIp         string  `json:"censored_ip"`
	FirstJoined        uint    `json:"first_joined"`
}

// MinerGetBalance takes a mining wallet address and gets the balance in gwei. Returns the balance and nil on success,
// or 0 and error on failure.
func MinerGetBalance(address string) (uint, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "balance", []string{}); err != nil {
		return 0, err
	}

	return uint(response.Result.(float64)), nil
}

// MinerGetCurrent takes a mining wallet address and gets the current effective and reported hashrate of that address.
// Returns a WorkerCurrentStats instance and nil on success, or an empty WorkerCurrentStats and error on failure.
func MinerGetCurrent(address string) (WorkerCurrentStats, error) {
	var (
		response Response
		data     WorkerCurrentStats
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "current", []string{}); err != nil {
		return data, err
	}

	data.EffectiveHashrate = uint(response.Result.(map[string]interface{})["effective_hashrate"].(float64))
	data.ReportedHashrate = uint(response.Result.(map[string]interface{})["reported_hashrate"].(float64))

	return data, nil
}

// MinerGetDaily takes a mining wallet address and gets the effective and reported hashrate of that address, as well
// as it's amount of stale and valid shares over the last 24 hours. Returns a MinerDailyStats instance and nil on success,
// or an empty MinerDailyStats and error on failure.
func MinerGetDaily(address string) (MinerDailyStats, error) {
	var (
		response Response
		data     MinerDailyStats
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "daily", []string{}); err != nil {
		return data, err
	}

	data.ReportedHashrate = response.Result.(map[string]interface{})["reported_hashrate"].(float64)
	data.EffectiveHashrate = response.Result.(map[string]interface{})["effective_hashrate"].(float64)
	data.StaleShares = int(response.Result.(map[string]interface{})["stale_shares"].(float64))
	data.ValidShares = int(response.Result.(map[string]interface{})["valid_shares"].(float64))

	return data, nil
}

// MinerGetStats takes a mining wallet address and gets the current and daily stats of that address. Returns a
// MinerStats instance and nil on success, or an empty MinerStats instance and error on failure.
func MinerGetStats(address string) (MinerStats, error) {
	var (
		response Response
		data     MinerStats
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "stats", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})
	currentData := responseData["current"].(map[string]interface{})
	dailyData := responseData["daily"].(map[string]interface{})

	data.Current.EffectiveHashrate = uint(currentData["effective_hashrate"].(float64))
	data.Current.ReportedHashrate = uint(currentData["reported_hashrate"].(float64))

	data.Daily.EffectiveHashrate = dailyData["effective_hashrate"].(float64)
	data.Daily.InvalidShares = int(dailyData["invalid_shares"].(float64))
	data.Daily.ReportedHashrate = dailyData["reported_hashrate"].(float64)
	data.Daily.StaleShares = int(dailyData["stale_shares"].(float64))
	data.Daily.ValidShares = int(dailyData["valid_shares"].(float64))

	return data, nil
}

// MinerGetWorkerCount takes a mining wallet address and gets the offline and online worker counts for that address. Returns
// a MinerWorkerCount instance and nil on success, or an empty MinerWorkerCount instance and error on failure.
func MinerGetWorkerCount(address string) (MinerWorkerCount, error) {
	var (
		response Response
		data     MinerWorkerCount
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "workerCount", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})

	data.Offline = int(responseData["offline"].(float64))
	data.Online = int(responseData["online"].(float64))

	return data, nil
}

// MinerGetWorkers takes a mining wallet address and gets a list of the active workers for that address. Returns a slice
// of MinerWorker instances and nil on success, or an empty slice and error on failure.
func MinerGetWorkers(address string) ([]MinerWorker, error) {
	var (
		response Response
		data     []MinerWorker
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "workers", []string{}); err != nil {
		return data, err
	}

	if response.Result != nil {
		responseData := response.Result.([]interface{})

		for _, worker := range responseData {
			workerData := worker.(map[string]interface{})

			data = append(data, MinerWorker{
				Name:                   workerData["name"].(string),
				Online:                 workerData["online"].(bool),
				DuplicateWorkersMerged: int(workerData["duplicate_workers_merged"].(float64)),
				ReportedHashrate:       uint(workerData["reported_hashrate"].(float64)),
				EffectiveHashrate:      uint(workerData["effective_hashrate"].(float64)),
				ValidShares:            int(workerData["valid_shares"].(float64)),
				StaleShares:            int(workerData["stale_shares"].(float64)),
				InvalidShares:          int(workerData["invalid_shares"].(float64)),
				LastSeen:               int(workerData["last_seen"].(float64)),
			})
		}
	}

	return data, nil
}

// MinerGetChart takes a mining wallet address and gets a list of the chart data for that address. Returns a slice of
// MinerChartData instances and nil on success, or an empty slice and error on failure.
func MinerGetChart(address string) ([]MinerChartData, error) {
	var (
		response Response
		data     []MinerChartData
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "chart", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.([]interface{})

	for _, chart := range responseData {
		chartData := chart.(map[string]interface{})

		data = append(data, MinerChartData{
			Timestamp:                int(chartData["timestamp"].(float64)),
			EffectiveHashrate:        uint(chartData["effective_hashrate"].(float64)),
			AverageEffectiveHashrate: chartData["average_effective_hashrate"].(float64),
			ReportedHashrate:         uint(chartData["reported_hashrate"].(float64)),
			ValidShares:              int(chartData["valid_shares"].(float64)),
			StaleShares:              int(chartData["stale_shares"].(float64)),
			InvalidShares:            int(chartData["invalid_shares"].(float64)),
		})
	}

	return data, nil
}

// MinerGetPayments takes a mining wallet address and a page number, and gets a list of payment data for that address + page.
// Returns a MinerPaymentData instance and nil on success, or an empty MinerPaymentData instance and error on failure.
func MinerGetPayments(address string, page int) (MinerPaymentData, error) {
	var (
		response Response
		data     MinerPaymentData
		err      error
	)

	pageStr := strconv.Itoa(page)

	if response, err = sendAPIRequest(Miner, address, "payments", []string{"page=" + pageStr}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})

	if responseData["data"] != nil {
		for _, payment := range responseData["data"].([]interface{}) {
			paymentData := payment.(map[string]interface{})

			data.Data = append(data.Data, MinerPayment{
				Txid:      paymentData["txid"].(string),
				Amount:    uint(paymentData["amount"].(float64)),
				Timestamp: uint(paymentData["timestamp"].(float64)),
				Duration:  uint(paymentData["duration"].(float64)),
			})
		}
	}

	data.ItemsPerPage = int(responseData["items_per_page"].(float64))
	data.TotalItems = int(responseData["total_items"].(float64))
	data.TotalPages = int(responseData["total_pages"].(float64))

	return data, nil
}

// MinerGetPaymentCount takes a mining wallet address and gets the number of payments made to that address. Returns the
// number of payments as an int and nil on success, or 0 and error on failure.
func MinerGetPaymentCount(address string) (int, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "paymentCount", []string{}); err != nil {
		return 0, err
	}

	return int(response.Result.(float64)), nil
}

// MinerGetPaymentChart takes a mining wallet address and gets a list of payments made to that address. Returns a slice of
// MinerPaymentChart instances and nil on success, an empty slice and error on failure.
func MinerGetPaymentChart(address string) ([]MinerPaymentChart, error) {
	var (
		response Response
		data     []MinerPaymentChart
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "paymentsChart", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.([]interface{})

	for _, payment := range responseData {
		paymentData := payment.(map[string]interface{})

		data = append(data, MinerPaymentChart{
			Amount:    uint(paymentData["amount"].(float64)),
			Timestamp: uint(paymentData["timestamp"].(float64)),
		})
	}

	return data, nil
}

// MinerGetBlocks takes an address and page number, and gets a list of blocks mined from that address. Returns a
// MinerBlockData instance and nil on success, an empty MinerBlockData and error on failure.
func MinerGetBlocks(address string, page int) (MinerBlockData, error) {
	var (
		response Response
		data     MinerBlockData
		err      error
	)

	pageStr := strconv.Itoa(page)

	if response, err = sendAPIRequest(Miner, address, "blocks", []string{"page=" + pageStr}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})

	if responseData["data"] != nil {
		for _, block := range responseData["data"].([]interface{}) {
			blockData := block.(map[string]interface{})

			data.Data = append(data.Data, Block{
				Hash:                  blockData["hash"].(string),
				Number:                uint(blockData["number"].(float64)),
				Type:                  blockData["type"].(string),
				Miner:                 blockData["miner"].(string),
				Difficulty:            uint(blockData["difficulty"].(float64)),
				Timestamp:             uint(blockData["timestamp"].(float64)),
				Confirmed:             blockData["confirmed"].(bool),
				RoundTime:             uint(blockData["round_time"].(float64)),
				Luck:                  blockData["difficulty"].(float64),
				ServerName:            blockData["server_name"].(string),
				BlockReward:           uint(blockData["block_reward"].(float64)),
				BlockFees:             uint(blockData["block_fees"].(float64)),
				UncleInclusionRewards: uint(blockData["uncle_inclusion_rewards"].(float64)),
				TotalRewards:          uint(blockData["total_rewards"].(float64)),
			})
		}
	}

	data.ItemsPerPage = int(responseData["items_per_page"].(float64))
	data.TotalItems = int(responseData["total_items"].(float64))
	data.TotalPages = int(responseData["total_pages"].(float64))

	return data, nil
}

// MinerGetBlockCount takes a mining wallet address and gets the number of blocks mined by that address. Returns the number
// of blocks mined as an int and nil on success, or 0 and error on failure.
func MinerGetBlockCount(address string) (int, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "blockCount", []string{}); err != nil {
		return 0, err
	}

	return int(response.Result.(float64)), nil
}

// MinerGetDetails takes a mining wallet address and gets the overall meta details of that wallet. Returns a MinerDetails
// instance and nil on success, or an empty MinerDetails instance and error on failure.
func MinerGetDetails(address string) (MinerDetails, error) {
	var (
		response Response
		data     MinerDetails
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "details", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})

	data.MinPayoutThreshold = uint(responseData["min_payout_threshold"].(float64))
	data.PoolDonation = responseData["pool_donation"].(float64)
	data.MaxFeePrice = uint(responseData["max_fee_price"].(float64))
	data.CensoredEmail = responseData["censored_email"].(string)
	data.CensoredIp = responseData["censored_ip"].(string)
	data.FirstJoined = uint(responseData["first_joined"].(float64))

	return data, nil
}

// MinerGetEstimatedDailyRevenue takes a mining address and gets the estimated daily revenue in gwei. Returns the estimated
// daily revenue as an int and nil on success, or 0 and error on failure.
func MinerGetEstimatedDailyRevenue(address string) (int, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "estimatedDailyRevenue", []string{}); err != nil {
		return 0, err
	}

	return int(response.Result.(float64)), nil
}

// MinerGetRoundShare takes a mining address and gets the current round share in percentage. Returns the round share as
// a float64 and nil on success, or 0.0 and error on failure.
func MinerGetRoundShare(address string) (float64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "roundShare", []string{}); err != nil {
		return 0.0, err
	}

	return response.Result.(float64), nil
}

// MinerGetTotalPaid takes a mining address and gets the total amount of gwei paid to that address. Returns the amount paid
// as an int and nil on success, or 0 and error on failure.
func MinerGetTotalPaid(address string) (int, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "totalPaid", []string{}); err != nil {
		return 0, err
	}

	return int(response.Result.(float64)), nil
}

// MinerGetTotalDonated takes a mining address and gets the total amount of gwei donated from that address to the pool.
// Returns the amount donated as an int and nil on success, or -0 and error on failure.
func MinerGetTotalDonated(address string) (int, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "totalDonated", []string{}); err != nil {
		return -0, err
	}

	return int(response.Result.(float64)), nil
}
