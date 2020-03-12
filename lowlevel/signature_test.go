package lowlevel

import (
	"net/url"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestMakeDevMsgSignature(t *testing.T) {
	c.Convey("签名的计算应该跟官方的测试工具一样", t, func() {
		values := []string{
			"kjr2TKI8umCBfVF3wAHk8JiPwma5VBme",
			"6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC/nIX3ZNt9w==",
			"VHh7ymSeb0jc4lSb",
			"1583940690",
		}
		devMsgSignature := makeDevMsgSignature(values...)
		msgSignature := "1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011"
		c.So(devMsgSignature, c.ShouldEqual, msgSignature)
	})
}

func TestVerifyURLSignature(t *testing.T) {
	token := "kjr2TKI8umCBfVF3wAHk8JiPwma5VBme"

	c.Convey("带有完好签名的 URL 应该能通过校验", t, func() {
		s := "http://test.example.com/?echostr=6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC%2FnIX3ZNt9w%3D%3D&msg_signature=1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011&timestamp=1583940690&nonce=VHh7ymSeb0jc4lSb"
		u, err := url.Parse(s)
		c.So(err, c.ShouldBeNil)

		ok := VerifyURLSignature(token, u)
		c.So(ok, c.ShouldBeTrue)
	})

	c.Convey("签名受损的 URL 应该通不过校验", t, func() {
		s := "http://test.example.com/?echostr=6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC%2FnIX3ZNt9w%3D%3D&msg_signature=1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb010&timestamp=1583940690&nonce=VHh7ymSeb0jc4lSb"
		u, err := url.Parse(s)
		c.So(err, c.ShouldBeNil)

		ok := VerifyURLSignature(token, u)
		c.So(ok, c.ShouldBeFalse)
	})
}
