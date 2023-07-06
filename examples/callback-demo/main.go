//go:build examples

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/xen0n/go-workwx/v2"
)

type dummyRxMessageHandler struct{}

var _ workwx.RxMessageHandler = dummyRxMessageHandler{}

// OnIncomingMessage 一条消息到来时的回调。
func (dummyRxMessageHandler) OnIncomingMessage(msg *workwx.RxMessage) error {
	// You can do much more!
	fmt.Printf("incoming message: %s\n", msg)
	return nil
}

func main() {
	pAddr := flag.String("addr", "[::]:8000", "address and port to listen on")
	pToken := flag.String("token", "", "configured Token of your work weixin app")
	pEncodingAESKey := flag.String("key", "", "configured EncodingAESKey of your work weixin app")

	flag.Parse()

	hh, err := workwx.NewHTTPHandler(*pToken, *pEncodingAESKey, dummyRxMessageHandler{})
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", hh)

	err = http.ListenAndServe(*pAddr, mux)
	if err != nil {
		panic(err)
	}
}
