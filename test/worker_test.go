package main

import (
	"fmt"
	"testing"

	"../pkg/api"
)

func TestWorkerEndpoints(t *testing.T) {
	// README: edit this address and worker name to whatever address/worker pair you want to test with
	const ADDR = ""
	const WORKER = ""

	t.Run("WorkerGetCurrent", func(t *testing.T) {
		if current, err := api.WorkerGetCurrent(ADDR, WORKER); err == nil {
			fmt.Printf("Get Current: %+v\n", current)
		} else {
			t.Errorf("WorkerGetCurrent failed with: %v", err)
		}
	})

	t.Run("WorkerGetDaily", func(t *testing.T) {
		if daily, err := api.WorkerGetDaily(ADDR, WORKER); err == nil {
			fmt.Printf("Get Current: %+v\n", daily)
		} else {
			t.Errorf("WorkerGetDaily failed with: %v", err)
		}
	})

	t.Run("WorkerGetStats", func(t *testing.T) {
		if stats, err := api.WorkerGetStats(ADDR, WORKER); err == nil {
			fmt.Printf("Get Daily: %+v\n", stats)
		} else {
			t.Errorf("WorkerGetStats failed with: %v", err)
		}
	})

	t.Run("WorkerGetChart", func(t *testing.T) {
		if chart, err := api.WorkerGetChart(ADDR, WORKER); err == nil {
			fmt.Printf("Get Stats: %+v\n", chart)
		} else {
			t.Errorf("WorkerGetChart failed with: %v", err)
		}
	})
}
