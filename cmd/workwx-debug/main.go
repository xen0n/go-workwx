// +build debug

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/xen0n/go-workwx"
)

func main() {
	corpID := os.Getenv("TEST_WORKWX_CORPID")
	corpSecret := os.Getenv("TEST_WORKWX_CORPSECRET")

	if corpID == "" {
		fmt.Print("fatal: please set TEST_WORKWX_CORPID")
		os.Exit(1)
	}
	if corpSecret == "" {
		fmt.Print("fatal: please set TEST_WORKWX_CORPSECRET")
		os.Exit(1)
	}

	c := workwx.Default()
	c.CorpID = corpID

	app := c.WithApp(corpSecret)
	app.SpawnAccessTokenRefresher()
	time.Sleep(5 * time.Second)
}
