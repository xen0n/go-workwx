package httpapi

import (
	"net/url"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestURLValuesForEchoTestAPI(t *testing.T) {
	c.Convey("合法的测试请求参数可以解出 EchoTestAPIArgs", t, func() {
		s := "http://test.example.com/?echostr=6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC%2FnIX3ZNt9w%3D%3D&msg_signature=1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011&timestamp=1583940690&nonce=VHh7ymSeb0jc4lSb"
		u, err := url.Parse(s)
		c.So(err, c.ShouldBeNil)

		req := URLValuesForEchoTestAPI(u.Query())
		args, err := req.ToEchoTestAPIArgs()
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
