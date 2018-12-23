package workwx

import (
	"fmt"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestRecipientValidation(t *testing.T) {

	manyStringsFactory := func(n int) []string {
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = fmt.Sprintf("x%d", i)
		}
		return result
	}

	c.Convey("Recipient 校验逻辑", t, func() {
		a := Recipient{}

		c.Reset(func() {
			a = Recipient{}
		})

		c.Convey("正向用例", func() {
			c.Convey("发给单个用户", func() {
				a.UserIDs = []string{"foo"}

				c.Convey("应该只通过 isValidForMessageSend 校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeTrue)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给 1000 个用户", func() {
				a.UserIDs = manyStringsFactory(1000)

				c.Convey("应该只通过 isValidForMessageSend 校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeTrue)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给单个部门", func() {
				a.PartyIDs = []string{"foo"}

				c.Convey("应该只通过 isValidForMessageSend 校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeTrue)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给 100 个部门", func() {
				a.PartyIDs = manyStringsFactory(100)

				c.Convey("应该只通过 isValidForMessageSend 校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeTrue)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给单个标签", func() {
				a.TagIDs = []string{"foo"}

				c.Convey("应该只通过 isValidForMessageSend 校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeTrue)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给 100 个标签", func() {
				a.TagIDs = manyStringsFactory(100)

				c.Convey("应该只通过 isValidForMessageSend 校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeTrue)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给 chatid", func() {
				a.ChatID = "foo"

				c.Convey("应该只通过 isValidForAppchatSend 校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeFalse)
					c.So(a.isValidForAppchatSend(), c.ShouldBeTrue)
				})
			})

			// TODO: 这个不知道行不行，要线上验证一下，现在是放行的
			c.Convey("发给单个用户和单个部门", func() {
				a.UserIDs = []string{"foo"}
				a.PartyIDs = []string{"bar"}

				c.Convey("目前应该只通过 isValidForMessageSend 校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeTrue)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})
		})

		c.Convey("反向用例", func() {
			c.Convey("空 Recipient", func() {
				c.Convey("不应该通过校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeFalse)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给 chatid 和单个用户", func() {
				a.UserIDs = []string{"foo"}
				a.ChatID = "bar"

				c.Convey("不应该通过校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeFalse)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给 1001 个用户", func() {
				a.UserIDs = manyStringsFactory(1001)

				c.Convey("不应该通过校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeFalse)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给 101 个部门", func() {
				a.PartyIDs = manyStringsFactory(101)

				c.Convey("不应该通过校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeFalse)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})

			c.Convey("发给 101 个标签", func() {
				a.TagIDs = manyStringsFactory(101)

				c.Convey("不应该通过校验", func() {
					c.So(a.isValidForMessageSend(), c.ShouldBeFalse)
					c.So(a.isValidForAppchatSend(), c.ShouldBeFalse)
				})
			})
		})
	})
}
