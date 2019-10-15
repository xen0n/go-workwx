package workwx_test

import (
	"net/http"

	"github.com/xen0n/go-workwx"
)

func ExampleWorkwx() {
	corpID := "your_corpid"
	corpSecret := "your_corpsecret"
	agentID := int64(1234567)

	client := workwx.New(corpID)

	// there're advanced options
	_ = workwx.New(
		corpID,
		workwx.WithQYAPIHost("http://localhost:8888"),
		workwx.WithHTTPClient(&http.Client{}),
	)

	// work with individual apps
	app := client.WithApp(corpSecret, agentID)
	app.SpawnAccessTokenRefresher()

	// see other examples for more details
}

func ExampleWorkwxApp_SendTextMessage() {
	corpID := "your_corpid"
	corpSecret := "your_corpsecret"
	agentID := int64(1234567)

	client := workwx.New(corpID)

	app := client.WithApp(corpSecret, agentID)
	// preferably do this at app initialization
	app.SpawnAccessTokenRefresher()

	// send to user(s)
	to1 := workwx.Recipient{
		UserIDs: []string{"testuser"},
	}
	_ = app.SendTextMessage(&to1, "send to user(s)", false)

	// "safe" message
	to2 := workwx.Recipient{
		UserIDs: []string{"testuser"},
	}
	_ = app.SendTextMessage(&to2, "safe message", true)

	// send to party(parties)
	to3 := workwx.Recipient{
		PartyIDs: []string{"testdept"},
	}
	_ = app.SendTextMessage(&to3, "send to party(parties)", false)

	// send to tag(s)
	to4 := workwx.Recipient{
		TagIDs: []string{"testtag"},
	}
	_ = app.SendTextMessage(&to4, "send to tag(s)", false)

	// send to chatid
	to5 := workwx.Recipient{
		ChatID: "testchat",
	}
	_ = app.SendTextMessage(&to5, "send to chatid", false)
}
