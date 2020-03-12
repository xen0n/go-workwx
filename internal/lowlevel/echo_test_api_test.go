package lowlevel

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestURLEchoTestAdapter(t *testing.T) {
	c.Convey("合法的测试请求参数可以解出 EchoTestAPIArgs", t, func() {
		s := "http://test.example.com/?echostr=6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC%2FnIX3ZNt9w%3D%3D&msg_signature=1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011&timestamp=1583940690&nonce=VHh7ymSeb0jc4lSb"
		u, err := url.Parse(s)
		c.So(err, c.ShouldBeNil)

		req := URLValuesEchoTestAdapter(u.Query())
		args, err := req.ParseEchoTestAPIArgs()
		c.So(err, c.ShouldBeNil)

		expected := EchoTestAPIArgs{
			MsgSignature: "1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011",
			Timestamp:    1583940690,
			Nonce:        "VHh7ymSeb0jc4lSb",
			EchoStr:      "6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC/nIX3ZNt9w==",
		}
		c.So(args, c.ShouldResemble, expected)
	})
}

func TestHTTPEchoTestAPIHandler(t *testing.T) {
	c.Convey("HTTP 测试回调模式请求应该能端到端成功", t, func() {
		token := "kjr2TKI8umCBfVF3wAHk8JiPwma5VBme"
		encodingAESKey := "4Ma3YBrSBbX2aez8MJpXGBne5LSDwgGqHbhM9WPYIws"
		handler, err := NewHTTPEchoTestAPIHandler(token, encodingAESKey)
		c.So(err, c.ShouldBeNil)
		c.So(handler, c.ShouldNotBeNil)

		mux := http.NewServeMux()
		mux.Handle("/test", handler)
		server := httptest.NewServer(mux)
		defer server.Close()

		resp, err := http.DefaultClient.Get(server.URL + "/test?echostr=6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC%2FnIX3ZNt9w%3D%3D&msg_signature=1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011&timestamp=1583940690&nonce=VHh7ymSeb0jc4lSb")
		c.So(err, c.ShouldBeNil)
		c.So(resp.StatusCode, c.ShouldEqual, http.StatusOK)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		c.So(err, c.ShouldBeNil)
		c.So(body, c.ShouldResemble, []byte("94966531020182955848408"))
	})
}
