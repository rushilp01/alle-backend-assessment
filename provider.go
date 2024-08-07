package main

import (
	"fmt"
	"math/rand"
)

var providers = [][]string{
	{
		"Hello from Provider A - 1",
		"Hello from Provider A - 2",
		"Hello from Provider A - 3",
	},
	{
		"Hello from Provider B - 1",
		"Hello from Provider B - 2",
		"Hello from Provider B - 3",
	},
	{
		"Hello from Provider C - 1",
		"Hello from Provider C - 2",
		"Hello from Provider C - 3",
	},
}

func getRandomText(provider int) string {
	texts := providers[provider]
	return texts[rand.Intn(len(texts))]

}

func getProvider() int {
	mutex.Lock()
	defer mutex.Unlock()
	return currentProvider
}

func switchProvider(newProvider int) {
	mutex.Lock()
	defer mutex.Unlock()
	currentProvider = newProvider
}

func monitorProviders() {
	for {

		bestProvider := evaluateProviders()
		switchProvider(bestProvider)
		if currentProvider != bestProvider {
			fmt.Printf("Switched to provider %d\n", bestProvider)
		}

	}
}

func evaluateProviders() int {
	bestProvider := 0
	bestScore := -1.0
	for i, metric := range metrics {
		successRate := float64(metric.Successes) / float64(metric.Successes+metric.Failures)
		score := successRate
		if score > bestScore {
			bestScore = score
			bestProvider = i
		}
	}
	return bestProvider
}
