package commands

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

// newTransportWithKeyLog initializes a HTTP Transport with KeyLogWriter
func newTransportWithKeyLog(keyLog io.Writer) *http.Transport {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{KeyLogWriter: keyLog, InsecureSkipVerify: true},

		// Copy of http.DefaultTransport
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	if err := http2.ConfigureTransport(transport); err != nil {
		panic(err)
	}
	return transport
}
