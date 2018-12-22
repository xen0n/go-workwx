package workwx

import (
	"net/http"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestWithHTTPClient(t *testing.T) {
	c.Convey("给定一个 options", t, func() {
		opts := options{}

		c.Convey("用 WithHTTPClient 修饰它", func() {
			newClient := http.Client{}
			o := WithHTTPClient(&newClient)
			o.ApplyTo(&opts)

			c.Convey("options.HTTP 应该变了", func() {
				c.So(opts.HTTP, c.ShouldEqual, &newClient)
			})
		})
	})
}
