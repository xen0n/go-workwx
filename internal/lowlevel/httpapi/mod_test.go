package httpapi

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestLowlevelHandler(t *testing.T) {
	c.Convey("E2E HTTP handler tests", t, func() {
		token := "kjr2TKI8umCBfVF3wAHk8JiPwma5VBme"
		encodingAESKey := "4Ma3YBrSBbX2aez8MJpXGBne5LSDwgGqHbhM9WPYIws"
		handler, err := NewLowlevelHandler(token, encodingAESKey, nil)
		c.So(err, c.ShouldBeNil)
		c.So(handler, c.ShouldNotBeNil)

		mux := http.NewServeMux()
		mux.Handle("/test", handler)
		server := httptest.NewServer(mux)
		defer server.Close()

		c.Convey("测试回调模式请求", func() {
			resp, err := http.DefaultClient.Get(server.URL + "/test?echostr=6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC%2FnIX3ZNt9w%3D%3D&msg_signature=1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011&timestamp=1583940690&nonce=VHh7ymSeb0jc4lSb")
			c.So(err, c.ShouldBeNil)
			c.So(resp.StatusCode, c.ShouldEqual, http.StatusOK)
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			c.So(err, c.ShouldBeNil)
			c.So(body, c.ShouldResemble, []byte("94966531020182955848408"))
		})
	})
}
