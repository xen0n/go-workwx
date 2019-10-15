package workwx

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestWorkwxCtor(t *testing.T) {
	c.Convey("不带参数构造 Workwx 实例", t, func() {
		a := New("testcorpid")

		c.Convey("corpid 应该正确设置了", func() {
			c.So(a.CorpID, c.ShouldEqual, "testcorpid")
		})
	})
}
