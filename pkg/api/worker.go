package api

// WorkerCurrentStats contains hashrate stats - used by multiple endpoints.
type WorkerCurrentStats struct {
	EffectiveHashrate int64 `json:"effective_hashrate"`
	ReportedHashrate  int64 `json:"reported_hashrate"`
}

// WorkerDailyStats contains daily hashrate and share stats from the /worker/{address}/{worker}/daily endpoint.
type WorkerDailyStats struct {
	EffectiveHashrate int64 `json:"effective_hashrate"`
	InvalidShares     int64 `json:"invalid_shares"`
	ReportedHashrate  int64 `json:"reported_hashrate"`
	StaleShares       int64 `json:"stale_shares"`
	ValidShares       int64 `json:"valid_shares"`
}

// WorkerStats contains current and daily stats from the /worker/{address}/{worker}/stats endpoint.
type WorkerStats struct {
	Current WorkerCurrentStats `json:"current"`
	Daily   WorkerDailyStats   `json:"daily"`
}

// WorkerChartData contains chart data entries from the /worker/{address}/{worker}/chart endpoint.
type WorkerChartData struct {
	Timestamp                int64   `json:"timestamp"`
	EffectiveHashrate        int64   `json:"effective_hashrate"`
	AverageEffectiveHashrate float64 `json:"average_effective_hashrate"`
	ReportedHashrate         int64   `json:"reported_hashrate"`
	ValidShares              int64   `json:"valid_shares"`
	StaleShares              int64   `json:"stale_shares"`
	InvalidShares            int64   `json:"invalid_shares"`
}

// WorkerGetCurrent takes a mining wallet address and worker name, and gets the current effective and reported hashrate of
// that address. Returns a WorkerCurrentStats instance and nil on success, or an empty WorkerCurrentStats and error on failure.
func WorkerGetCurrent(address string, worker string) (WorkerCurrentStats, error) {
	var (
		response Response
		data     WorkerCurrentStats
		err      error
	)

	if response, err = sendAPIRequest(Worker, address, worker, []string{"current"}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})

	data.EffectiveHashrate = int64(responseData["effective_hashrate"].(float64))
	data.ReportedHashrate = int64(responseData["reported_hashrate"].(float64))

	return data, nil
}

// WorkerGetDaily takes a mining wallet address and worker name, and gets the daily effective and reported hashrate of that
// address as well as it's amount of stale and valid shares over the last 24 hours. Returns a WorkerDailyStats instance
// and nil on success, an empty WorkerDailyStats and error on failure.
func WorkerGetDaily(address string, worker string) (WorkerDailyStats, error) {
	var (
		response Response
		data     WorkerDailyStats
		err      error
	)

	if response, err = sendAPIRequest(Worker, address, worker, []string{"daily"}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})

	data.EffectiveHashrate = int64(responseData["effective_hashrate"].(float64))
	data.InvalidShares = int64(responseData["invalid_shares"].(float64))
	data.ReportedHashrate = int64(responseData["reported_hashrate"].(float64))
	data.StaleShares = int64(responseData["stale_shares"].(float64))
	data.ValidShares = int64(responseData["valid_shares"].(float64))

	return data, nil
}

// WorkerGetStats takes a mining wallet address and worker name, and gets the current and daily stats of that worker. Returns
// a WorkerStats instance and nil on success, or an empty WorkerStats instance and error on failure.
func WorkerGetStats(address string, worker string) (WorkerStats, error) {
	var (
		response Response
		data     WorkerStats
		err      error
	)

	if response, err = sendAPIRequest(Worker, address, worker, []string{"stats"}); err != nil {
		return data, err
	}

	responseData := response.Result.(map[string]interface{})
	currentData := responseData["current"].(map[string]interface{})
	dailyData := responseData["daily"].(map[string]interface{})

	data.Current.EffectiveHashrate = int64(currentData["effective_hashrate"].(float64))
	data.Current.ReportedHashrate = int64(currentData["reported_hashrate"].(float64))

	data.Daily.EffectiveHashrate = int64(dailyData["effective_hashrate"].(float64))
	data.Daily.InvalidShares = int64(dailyData["invalid_shares"].(float64))
	data.Daily.ReportedHashrate = int64(dailyData["reported_hashrate"].(float64))
	data.Daily.StaleShares = int64(dailyData["stale_shares"].(float64))
	data.Daily.ValidShares = int64(dailyData["valid_shares"].(float64))

	return data, nil
}

// WorkerGetChart takes a mining wallet address and worker name, and gets a list of chart data for that address. Returns
// a slice of MinerChartData instances and nil on success, or an empty slice and error on failure.
func WorkerGetChart(address string, worker string) ([]WorkerChartData, error) {
	var (
		response Response
		data     []WorkerChartData
		err      error
	)

	if response, err = sendAPIRequest(Worker, address, worker, []string{"chart"}); err != nil {
		return data, err
	}

	responseData := response.Result.([]interface{})

	for _, chartDataPoint := range responseData {
		chartData := chartDataPoint.(map[string]interface{})

		data = append(data, WorkerChartData{
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
