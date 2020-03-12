package lowlevel

import (
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
