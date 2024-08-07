package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"sync"
	"time"
)

type ProviderMetrics struct {
	Successes int
	Failures  int
}

var metrics = []ProviderMetrics{
	{7, 8},
	{6, 0},
	{10, 13},
}

var currentProvider int
var mutex sync.Mutex

func main() {
	router := gin.Default()
	router.GET("/stream", streamHandler)
	go monitorProviders()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func streamHandler(c *gin.Context) {
	provider := getProvider()

	clientGone := c.Writer.CloseNotify()

	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case <-time.After(1 * time.Second):
			if provider != currentProvider {
				provider = currentProvider
			}
			message := getRandomText(provider)
			if rand.Intn(100)%5 == 0 || provider == 1 {
				updateMetrics(provider, false)
			} else {
				updateMetrics(provider, true)
			}

			c.SSEvent("message", message)
		}
		return true
	})
}

func updateMetrics(provider int, success bool) {
	if success {
		metrics[provider].Successes++
	} else {
		metrics[provider].Failures++
	}
}
