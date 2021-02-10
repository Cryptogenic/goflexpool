package api

import (
	"strconv"
)

// PoolHashrate contains pool hashrate stats data from the /pool/hashrate endpoint.
type PoolHashrate struct {
	As    int64 `json:"as"`
	Au    int64 `json:"au"`
	Eu    int64 `json:"eu"`
	Sa    int64 `json:"sa"`
	Total int64 `json:"total"`
	Us    int64 `json:"us"`
}

// PoolHashrateChartData contains pool data entries from the /pool/hashrateChart endpoint.
type PoolHashrateChartData struct {
	As        int64 `json:"as"`
	Au        int64 `json:"au"`
	Eu        int64 `json:"eu"`
	Sa        int64 `json:"sa"`
	Timestamp int64 `json:"timestamp"`
	Total     int64 `json:"total"`
	Us        int64 `json:"us"`
}

// PoolBlockCount contains pool block data from the /pool/blockCount endpoint.
type PoolBlockCount struct {
	Confirmed   int64 `json:"confirmed"`
	Unconfirmed int64 `json:"unconfirmed"`
}

// PoolBlockData contains paged block data from the /pool/blocks endpoint.
type PoolBlockData struct {
	Data         []Block `json:"data"`
	ItemsPerPage int64   `json:"items_per_page"`
	TotalItems   int64   `json:"total_items"`
	TotalPages   int64   `json:"total_pages"`
}

// PoolMinerInfo contains miner data for the top miners from the /pool/topMiners endpoint.
type PoolMinerInfo struct {
	Address      string  `json:"address"`
	Hashrate     int64   `json:"hashrate"`
	TotalWorkers int64   `json:"total_workers"`
	Balance      int64   `json:"balance"`
	PoolDonation float64 `json:"pool_donation"`
	FirstJoined  int64   `json:"first_joined"`
}

// PoolDonatorInfo contains donation data for the top donators from the /pool/topDonators endpoint.
type PoolDonatorInfo struct {
	Address      string  `json:"address"`
	PoolDonation float64 `json:"pool_donation"`
	TotalDonated int64   `json:"total_donated"`
	FirstJoined  int64   `json:"first_joined"`
}

// PoolAvgLuckRoundTime contains luck data from the /pool/avgLuckRoundtime endpoint.
type PoolAvgLuckRoundTime struct {
	Luck      float64 `json:"luck"`
	RoundTime float64 `json:"round_time"`
}

// PoolGetHashrate gets the hashrate of the pool for each region in hashes per second. Returns a PoolHashrate instance and
// nil on success, or an empty PoolHashrate and error on failure.
func PoolGetHashrate() (PoolHashrate, error) {
	var (
		response Response
		data     PoolHashrate
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "hashrate", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})

	data.As = int64(responseData["as"].(float64))
	data.Au = int64(responseData["au"].(float64))
	data.Eu = int64(responseData["eu"].(float64))
	data.Sa = int64(responseData["sa"].(float64))
	data.Total = int64(responseData["total"].(float64))
	data.Us = int64(responseData["us"].(float64))

	return data, nil
}

// PoolGetHashrateChart gets a list of hashrate chart data for the pool. Returns a slice of PoolHashrateChartData instances
// and nil on success, or an empty slice and error on failure.
func PoolGetHashrateChart() ([]PoolHashrateChartData, error) {
	var (
		response Response
		data     []PoolHashrateChartData
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "hashrateChart", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.([]interface{})

	for _, statDataPoint := range responseData {
		statData := statDataPoint.(map[string]interface{})

		data = append(data, PoolHashrateChartData{
			As:        int64(statData["as"].(float64)),
			Au:        int64(statData["au"].(float64)),
			Eu:        int64(statData["eu"].(float64)),
			Sa:        int64(statData["sa"].(float64)),
			Timestamp: int64(statData["timestamp"].(float64)),
			Total:     int64(statData["total"].(float64)),
			Us:        int64(statData["us"].(float64)),
		})
	}

	return data, nil
}

// PoolGetMinersOnline gets how many miners are currently active on the pool. Returns the active miner count and nil on
// success, or -1 and error on failure.
func PoolGetMinersOnline() (int64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "minersOnline", []string{}); err != nil {
		return -1, err
	}

	return int64(response.Result.(float64)), nil
}

// PoolGetWorkersOnline gets how many workers are currently active on the pool. Returns the active worker count and nil on
// success, or -1 and error on failure.
func PoolGetWorkersOnline() (int64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "workersOnline", []string{}); err != nil {
		return -1, err
	}

	return int64(response.Result.(float64)), nil
}

// PoolGetBlocks takes a page number and gets a list of blocks the pool has mined from that page. Returns a PoolBlockData
// instance and nil on success, or an empty PoolBlockData and error on failure.
func PoolGetBlocks(page int) (PoolBlockData, error) {
	var (
		response Response
		data     PoolBlockData
		err      error
	)

	pageStr := strconv.Itoa(page)

	if response, err = sendAPIRequest(Pool, "", "blocks", []string{"page=" + pageStr}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})
	blocksData := responseData["data"].([]interface{})

	for _, blockDataPoint := range blocksData {
		blockData := blockDataPoint.(map[string]interface{})

		data.Data = append(data.Data, Block{
			Hash:                  blockData["hash"].(string),
			Number:                int64(blockData["number"].(float64)),
			Type:                  blockData["type"].(string),
			Miner:                 blockData["miner"].(string),
			Difficulty:            int64(blockData["difficulty"].(float64)),
			Timestamp:             int64(blockData["timestamp"].(float64)),
			Confirmed:             blockData["confirmed"].(bool),
			RoundTime:             int64(blockData["round_time"].(float64)),
			Luck:                  blockData["number"].(float64),
			ServerName:            blockData["server_name"].(string),
			BlockReward:           int64(blockData["block_reward"].(float64)),
			BlockFees:             int64(blockData["block_fees"].(float64)),
			UncleInclusionRewards: int64(blockData["uncle_inclusion_rewards"].(float64)),
			TotalRewards:          int64(blockData["total_rewards"].(float64)),
		})
	}

	data.ItemsPerPage = int64(responseData["items_per_page"].(float64))
	data.TotalItems = int64(responseData["total_items"].(float64))
	data.TotalPages = int64(responseData["total_pages"].(float64))

	return data, nil
}

// PoolGetBlockCount gets how many blocks have been mined by the pool. Returns a PoolBlockCount instance and nil on success,
// or an empty PoolBlockCount and error on failure.
func PoolGetBlockCount() (PoolBlockCount, error) {
	var (
		response Response
		data     PoolBlockCount
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "blockCount", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})

	data.Confirmed = int64(responseData["confirmed"].(float64))
	data.Unconfirmed = int64(responseData["unconfirmed"].(float64))

	return data, nil
}

// PoolGetTopMiners gets a list of the top miners in the pool. Returns a slice of PoolMinerInfo instances and nil on
// success, or an empty slice and error on failure.
func PoolGetTopMiners() ([]PoolMinerInfo, error) {
	var (
		response Response
		data     []PoolMinerInfo
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "topMiners", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.([]interface{})

	for _, minerDataPoint := range responseData {
		minerData := minerDataPoint.(map[string]interface{})

		data = append(data, PoolMinerInfo{
			Address:      minerData["address"].(string),
			Hashrate:     int64(minerData["hashrate"].(float64)),
			TotalWorkers: int64(minerData["total_workers"].(float64)),
			Balance:      int64(minerData["balance"].(float64)),
			PoolDonation: minerData["pool_donation"].(float64),
			FirstJoined:  int64(minerData["first_joined"].(float64)),
		})
	}

	return data, nil
}

// PoolGetTopDonators gets a list of the top donators in the pool. Returns a slice of PoolDonatorInfo instances and nil
// on success, or an empty slice and error on failure.
func PoolGetTopDonators() ([]PoolDonatorInfo, error) {
	var (
		response Response
		data     []PoolDonatorInfo
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "topDonators", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.([]interface{})

	for _, donatorDataPoint := range responseData {
		donatorData := donatorDataPoint.(map[string]interface{})

		data = append(data, PoolDonatorInfo{
			Address:      donatorData["address"].(string),
			PoolDonation: donatorData["pool_donation"].(float64),
			TotalDonated: int64(donatorData["total_donated"].(float64)),
			FirstJoined:  int64(donatorData["first_joined"].(float64)),
		})
	}

	return data, nil
}

// PoolGetAverageLuckRoundTime gets the pool's current average luck as a percent and roundtime in seconds. Returns a
// PoolAvgLuckRoundTime instance and nil on success, or an empty PoolAvgLuckRoundTime and error on failure.
func PoolGetAverageLuckRoundTime() (PoolAvgLuckRoundTime, error) {
	var (
		response Response
		data     PoolAvgLuckRoundTime
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "avgLuckRoundtime", []string{}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})

	data.Luck = responseData["luck"].(float64)
	data.RoundTime = responseData["round_time"].(float64)

	return data, nil
}

// PoolGetCurrentLuck gets the pool's current luck as a percent. Returns the current luck as a float64 and nil on success,
// or 0.0 and error on failure.
func PoolGetCurrentLuck() (float64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "currentLuck", []string{}); err != nil {
		return 0.0, err
	}

	return response.Result.(float64), nil
}

// PoolGetAverageBlockReward gets the pool's average block reward in gwei. Returns the average block reward as an int64
// and nil on success, or -1 and error on failure.
func PoolGetAverageBlockReward() (int64, error) {
	var (
		response Response
		err      error
	)

	if response, err = sendAPIRequest(Pool, "", "averageBlockReward", []string{}); err != nil {
		return -1, err
	}

	return int64(response.Result.(float64)), nil
}
