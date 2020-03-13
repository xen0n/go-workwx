package signature

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
		devMsgSignature := MakeDevMsgSignature(values...)
		msgSignature := "1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011"
		c.So(devMsgSignature, c.ShouldEqual, msgSignature)
	})
}

func TestVerifyHTTPRequestSignature(t *testing.T) {
	token := "kjr2TKI8umCBfVF3wAHk8JiPwma5VBme"

	c.Convey("带有完好签名的 URL 应该能通过校验", t, func() {
		s := "http://test.example.com/?echostr=6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC%2FnIX3ZNt9w%3D%3D&msg_signature=1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb011&timestamp=1583940690&nonce=VHh7ymSeb0jc4lSb"
		u, err := url.Parse(s)
		c.So(err, c.ShouldBeNil)

		ok := VerifyHTTPRequestSignature(token, u, "")
		c.So(ok, c.ShouldBeTrue)
	})

	c.Convey("带请求体的正常请求也应该能通过校验", t, func() {
		token2 := "kz7Yx62CH8SaLN"
		encodingAESKey := "cD0d7jx4tYvVtzqrmh3Dm3QFCXe6f8SlHoMtMh3qQEP"
		_ = encodingAESKey
		s := "http://test.example.com/?msg_signature=a44d1b7dc3dc4b02edd99dc32afa6a26be7c92f6&timestamp=1583995625&nonce=1584392382"
		body := "fyVN9BUnH6xNl0/VCHTCg6XGxxvoFdyZ7VQbAbBcQ79dVccYvdIyJYZSMMhkPy8LPaoe+V27qjoQ553fXavmrPOgxzHrKKf3YIu63pB4/nN0LY4S+GzUFi/LLFj1TI1sOh0q2jT5u1nYz5G3HFNLd3DFUvJVLtZl8mVMwUOzBpzKZhhmCvIJLVlvym9tt+VYwG06MSbmnf9AzYFKhm1BIy+95Q824ilY6Wy8l+vkFqEdASl2k83jMXvlkVRSB3hmNptWTbq8ygDdLtkCyA4UEiZvfTskqqVkbYwLptRl/1F+DVL9SH/TWt2j3dCm4rzY5jQG+fA+2xh2Kk+OrVgafqtQCfhjq8C38A2IwCcg2BONKk8s5ijFxQoszmXvJQ3zLYbjzRWQ40xhryDyPNNse7/slKVW2QMBEUPNOxUcX9k="
		u, err := url.Parse(s)
		c.So(err, c.ShouldBeNil)

		ok := VerifyHTTPRequestSignature(token2, u, body)
		c.So(ok, c.ShouldBeTrue)
	})

	c.Convey("签名受损的 URL 应该通不过校验", t, func() {
		s := "http://test.example.com/?echostr=6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC%2FnIX3ZNt9w%3D%3D&msg_signature=1ba3cb09c0d2c2b3ed6900d37f91a6efae6cb010&timestamp=1583940690&nonce=VHh7ymSeb0jc4lSb"
		u, err := url.Parse(s)
		c.So(err, c.ShouldBeNil)

		ok := VerifyHTTPRequestSignature(token, u, "")
		c.So(ok, c.ShouldBeFalse)
	})
}
