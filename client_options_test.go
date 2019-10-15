package workwx

import (
	"net/http"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestDefaultOptions(t *testing.T) {
	c.Convey("给定一个默认值 options", t, func() {
		opts := defaultOptions()

		c.Convey("opts.HTTP 应该是空值 http.Client", func() {
			c.So(opts.HTTP, c.ShouldNotBeNil)
			c.So(*opts.HTTP, c.ShouldResemble, http.Client{})
		})

		c.Convey("opts.QYAPIHost 应该是企业微信官方 API host", func() {
			c.So(opts.QYAPIHost, c.ShouldEqual, "https://qyapi.weixin.qq.com")
		})
	})
}

func TestWithHTTPClient(t *testing.T) {
	c.Convey("给定一个 options", t, func() {
		opts := options{}

		c.Convey("用 WithHTTPClient 修饰它", func() {
			newClient := http.Client{}
			o := WithHTTPClient(&newClient)
			o.applyTo(&opts)

			c.Convey("options.HTTP 应该变了", func() {
				c.So(opts.HTTP, c.ShouldEqual, &newClient)
			})
		})
	})
}

func TestWithQYAPIHost(t *testing.T) {
	c.Convey("给定一个 options", t, func() {
		opts := options{}

		c.Convey("用 WithQYAPIHost 修饰它", func() {
			o := WithQYAPIHost("http://localhost:8000")
			o.applyTo(&opts)

			c.Convey("options.QYAPIHost 应该变了", func() {
				c.So(opts.QYAPIHost, c.ShouldEqual, "http://localhost:8000")
			})
		})
	})
}
