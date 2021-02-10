package api

import (
	"strconv"
)

// Block contains information relevant to blocks mined - used by multiple endpoints.
type Block struct {
	Hash                  string  `json:"hash"`
	Number                int64   `json:"number"`
	Type                  string  `json:"type"`
	Miner                 string  `json:"miner"`
	Difficulty            int64   `json:"difficulty"`
	Timestamp             int64   `json:"timestamp"`
	Confirmed             bool    `json:"confirmed"`
	RoundTime             int64   `json:"round_time"`
	Luck                  float64 `json:"luck"`
	ServerName            string  `json:"server_name"`
	BlockReward           int64   `json:"block_reward"`
	BlockFees             int64   `json:"block_fees"`
	UncleInclusionRewards int64   `json:"uncle_inclusion_rewards"`
	TotalRewards          int64   `json:"total_rewards"`
}

// MinerDailyStats contains miner daily stats data from the /miner/{address}/stats and /miner/{address}/daily endpoint.
type MinerDailyStats struct {
	EffectiveHashrate float64 `json:"effective_hashrate"`
	InvalidShares     int64   `json:"invalid_shares"`
	ReportedHashrate  float64 `json:"reported_hashrate"`
	StaleShares       int64   `json:"stale_shares"`
	ValidShares       int64   `json:"valid_shares"`
}

// MinerDailyStats contains miner stats data from the /miner/{address}/stats endpoint.
type MinerStats struct {
	Current WorkerCurrentStats `json:"current"`
	Daily   MinerDailyStats    `json:"daily"`
}

// MinerWorkerCount contains worker count data from the /miner/{address}/workerCount endpoint.
type MinerWorkerCount struct {
	Online  int64 `json:"online"`
	Offline int64 `json:"offline"`
}

// MinerWorker contains worker data entries from the /miner/{address}/workers endpoint.
type MinerWorker struct {
	Name                   string `json:"name"`
	Online                 bool   `json:"online"`
	DuplicateWorkersMerged int64  `json:"duplicate_workers_merged"`
	ReportedHashrate       int64  `json:"reported_hashrate"`
	EffectiveHashrate      int64  `json:"effective_hashrate"`
	ValidShares            int64  `json:"valid_shares"`
	StaleShares            int64  `json:"stale_shares"`
	InvalidShares          int64  `json:"invalid_shares"`
	LastSeen               int64  `json:"last_seen"`
}

// MinerChartData contains chart data entries from the /miner/{address}/chart endpoint.
type MinerChartData struct {
	Timestamp                int64   `json:"timestamp"`
	EffectiveHashrate        int64   `json:"effective_hashrate"`
	AverageEffectiveHashrate float64 `json:"average_effective_hashrate"`
	ReportedHashrate         int64   `json:"reported_hashrate"`
	ValidShares              int64   `json:"valid_shares"`
	StaleShares              int64   `json:"stale_shares"`
	InvalidShares            int64   `json:"invalid_shares"`
}

// MinerPayment contains payment entries from the /miner/{address}/payments endpoint.
type MinerPayment struct {
	Txid      string `json:"txid"`
	Amount    int64  `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	Duration  int64  `json:"duration"`
}

// MinerPaymentData contains paged payment data from the /miner/{address}/payments endpoint.
type MinerPaymentData struct {
	Data         []MinerPayment `json:"data"`
	ItemsPerPage int64          `json:"items_per_page"`
	TotalItems   int64          `json:"total_items"`
	TotalPages   int64          `json:"total_pages"`
}

// MinerPaymentChart contains payment chart data from the /miner/{address}/paymentsChart endpoint.
type MinerPaymentChart struct {
	Amount    int64 `json:"amount"`
	Timestamp int64 `json:"timestamp"`
}

// MinerBlockData contains paged block data from the /miner/{address}/blocks endpoint.
type MinerBlockData struct {
	Data         []Block `json:"data"`
	ItemsPerPage int64   `json:"items_per_page"`
	TotalItems   int64   `json:"total_items"`
	TotalPages   int64   `json:"total_pages"`
}

// MinerBlockCount contains block count data from the /miner/{address}/blockCount endpoint.
type MinerBlockCount struct {
	Confirmed   int64 `json:"confirmed"`
	Unconfirmed int64 `json:"unconfirmed"`
}

// MinerDetails contains overview data from the /miner/{address}/details endpoint.
type MinerDetails struct {
	MinPayoutThreshold int64   `json:"min_payout_threshold"`
	PoolDonation       float64 `json:"pool_donation"`
	MaxFeePrice        int64   `json:"max_free_price"`
	CensoredEmail      string  `json:"censored_email"`
	CensoredIp         string  `json:"censored_ip"`
	FirstJoined        int64   `json:"first_joined"`
}

// MinerGetBalance takes a mining wallet address and gets the balance in gwei. Returns the balance and nil on success,
// or -1 and error on failure.
func MinerGetBalance(address string) (int64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "balance", []string{}); err != nil {
		return -1, err
	}

	return int64(response.Result.(float64)), nil
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

	data.EffectiveHashrate = int64(response.Result.(map[string]interface{})["effective_hashrate"].(float64))
	data.ReportedHashrate = int64(response.Result.(map[string]interface{})["reported_hashrate"].(float64))

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
	data.StaleShares = int64(response.Result.(map[string]interface{})["stale_shares"].(float64))
	data.ValidShares = int64(response.Result.(map[string]interface{})["valid_shares"].(float64))

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

	data.Current.EffectiveHashrate = int64(currentData["effective_hashrate"].(float64))
	data.Current.ReportedHashrate = int64(currentData["reported_hashrate"].(float64))

	data.Daily.EffectiveHashrate = dailyData["effective_hashrate"].(float64)
	data.Daily.InvalidShares = int64(dailyData["invalid_shares"].(float64))
	data.Daily.ReportedHashrate = dailyData["reported_hashrate"].(float64)
	data.Daily.StaleShares = int64(dailyData["stale_shares"].(float64))
	data.Daily.ValidShares = int64(dailyData["valid_shares"].(float64))

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

	data.Offline = int64(responseData["offline"].(float64))
	data.Online = int64(responseData["online"].(float64))

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
				DuplicateWorkersMerged: int64(workerData["duplicate_workers_merged"].(float64)),
				ReportedHashrate:       int64(workerData["reported_hashrate"].(float64)),
				EffectiveHashrate:      int64(workerData["effective_hashrate"].(float64)),
				ValidShares:            int64(workerData["valid_shares"].(float64)),
				StaleShares:            int64(workerData["stale_shares"].(float64)),
				InvalidShares:          int64(workerData["invalid_shares"].(float64)),
				LastSeen:               int64(workerData["last_seen"].(float64)),
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
			Timestamp:                int64(chartData["timestamp"].(float64)),
			EffectiveHashrate:        int64(chartData["effective_hashrate"].(float64)),
			AverageEffectiveHashrate: chartData["average_effective_hashrate"].(float64),
			ReportedHashrate:         int64(chartData["reported_hashrate"].(float64)),
			ValidShares:              int64(chartData["valid_shares"].(float64)),
			StaleShares:              int64(chartData["stale_shares"].(float64)),
			InvalidShares:            int64(chartData["invalid_shares"].(float64)),
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
				Amount:    int64(paymentData["amount"].(float64)),
				Timestamp: int64(paymentData["timestamp"].(float64)),
				Duration:  int64(paymentData["duration"].(float64)),
			})
		}
	}

	data.ItemsPerPage = int64(responseData["items_per_page"].(float64))
	data.TotalItems = int64(responseData["total_items"].(float64))
	data.TotalPages = int64(responseData["total_pages"].(float64))

	return data, nil
}

// MinerGetPaymentCount takes a mining wallet address and gets the number of payments made to that address. Returns the
// number of payments as an int64 and nil on success, or -1 and error on failure.
func MinerGetPaymentCount(address string) (int64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "paymentCount", []string{}); err != nil {
		return -1, err
	}

	return int64(response.Result.(float64)), nil
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
			Amount:    int64(paymentData["amount"].(float64)),
			Timestamp: int64(paymentData["timestamp"].(float64)),
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
				Number:                int64(blockData["number"].(float64)),
				Type:                  blockData["type"].(string),
				Miner:                 blockData["miner"].(string),
				Difficulty:            int64(blockData["difficulty"].(float64)),
				Timestamp:             int64(blockData["timestamp"].(float64)),
				Confirmed:             blockData["confirmed"].(bool),
				RoundTime:             int64(blockData["round_time"].(float64)),
				Luck:                  blockData["difficulty"].(float64),
				ServerName:            blockData["server_name"].(string),
				BlockReward:           int64(blockData["block_reward"].(float64)),
				BlockFees:             int64(blockData["block_fees"].(float64)),
				UncleInclusionRewards: int64(blockData["uncle_inclusion_rewards"].(float64)),
				TotalRewards:          int64(blockData["total_rewards"].(float64)),
			})
		}
	}

	data.ItemsPerPage = int64(responseData["items_per_page"].(float64))
	data.TotalItems = int64(responseData["total_items"].(float64))
	data.TotalPages = int64(responseData["total_pages"].(float64))

	return data, nil
}

// MinerGetBlockCount takes a mining wallet address and gets the number of blocks mined by that address. Returns the number
// of blocks mined as an int64 and nil on success, or -1 and error on failure.
func MinerGetBlockCount(address string) (int64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "blockCount", []string{}); err != nil {
		return -1, err
	}

	return int64(response.Result.(float64)), nil
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

	data.MinPayoutThreshold = int64(responseData["min_payout_threshold"].(float64))
	data.PoolDonation = responseData["pool_donation"].(float64)
	data.MaxFeePrice = int64(responseData["max_fee_price"].(float64))
	data.CensoredEmail = responseData["censored_email"].(string)
	data.CensoredIp = responseData["censored_ip"].(string)
	data.FirstJoined = int64(responseData["first_joined"].(float64))

	return data, nil
}

// MinerGetEstimatedDailyRevenue takes a mining address and gets the estimated daily revenue in gwei. Returns the estimated
// daily revenue as an int64 and nil on success, or -1 and error on failure.
func MinerGetEstimatedDailyRevenue(address string) (int64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "estimatedDailyRevenue", []string{}); err != nil {
		return -1, err
	}

	return int64(response.Result.(float64)), nil
}

// MinerGetRoundShare takes a mining address and gets the current round share in percentage. Returns the round share as
// a float64 and nil on success, or -1 and error on failure.
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
// as an int64 and nil on success, or -1 and error on failure.
func MinerGetTotalPaid(address string) (int64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "totalPaid", []string{}); err != nil {
		return -1, err
	}

	return int64(response.Result.(float64)), nil
}

// MinerGetTotalDonated takes a mining address and gets the total amount of gwei donated from that address to the pool.
// Returns the amount donated as an int64 and nil on success, or -1 and error on failure.
func MinerGetTotalDonated(address string) (int64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Miner, address, "totalDonated", []string{}); err != nil {
		return -1, err
	}

	return int64(response.Result.(float64)), nil
}
