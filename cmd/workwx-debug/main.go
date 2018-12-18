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
	userID := os.Getenv("TEST_WORKWX_USERID")
	chatID := os.Getenv("TEST_WORKWX_CHATID")

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
	if userID == "" {
		fmt.Println("fatal: please set TEST_WORKWX_USERID")
		os.Exit(1)
	}
	if chatID == "" {
		fmt.Println("fatal: please set TEST_WORKWX_CHATID")
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
	time.Sleep(time.Second)

	to1 := workwx.Recipient{
		UserIDs: []string{userID},
	}
	_ = app.SendTextMessage(&to1, "testtest", false)

	to2 := workwx.Recipient{
		ChatID: chatID,
	}
	_ = app.SendTextMessage(&to2, "testtest", false)
}
