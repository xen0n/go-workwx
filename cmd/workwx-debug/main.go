// +build debug

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/xen0n/go-workwx"
)

func main() {
	corpID := os.Getenv("TEST_WORKWX_CORPID")
	corpSecret := os.Getenv("TEST_WORKWX_CORPSECRET")
	agentIDStr := os.Getenv("TEST_WORKWX_AGENTID")

	if corpID == "" {
		fmt.Print("fatal: please set TEST_WORKWX_CORPID")
		os.Exit(1)
	}
	if corpSecret == "" {
		fmt.Print("fatal: please set TEST_WORKWX_CORPSECRET")
		os.Exit(1)
	}
	if agentIDStr == "" {
		fmt.Println("fatal: please set TEST_WORKWX_AGENTID")
		os.Exit(1)
	}

	agentID, err := strconv.Atoi(agentIDStr)
	if err != nil {
		fmt.Println("fatal: AgentID '%s' is not valid integer")
		os.Exit(1)
	}

	c := workwx.Default()
	c.CorpID = corpID

	app := c.WithApp(corpSecret, int64(agentID))
	app.SpawnAccessTokenRefresher()
	time.Sleep(5 * time.Second)
}
