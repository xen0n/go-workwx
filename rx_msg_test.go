package workwx

import (
	"testing"
	"time"

	c "github.com/smartystreets/goconvey/convey"
)

var cst = time.FixedZone("CST", 8*3600)

func TestRxMessage(t *testing.T) {
	c.Convey("解析接收的 XML 消息体", t, func() {
		c.Convey("文本消息", func() {
			body := []byte("<xml><ToUserName><![CDATA[ww6a112864f8022910]]></ToUserName><FromUserName><![CDATA[foobar]]></FromUserName><CreateTime>1583995625</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[x123]]></Content><MsgId>2018405441</MsgId><AgentID>1000002</AgentID></xml>")

			msg, err := fromEnvelope(body)
			c.So(err, c.ShouldBeNil)
			c.So(msg, c.ShouldNotBeNil)
			c.So(msg.String(), c.ShouldEqual, `RxMessage { FromUserID: "foobar", SendTime: 1583995625000000000, MsgType: "text", MsgID: 2018405441, AgentID: 1000002, Content: "x123" }`)
			c.So(msg.FromUserID, c.ShouldEqual, "foobar")
			c.So(msg.SendTime, c.ShouldEqual, time.Date(2020, 3, 12, 14, 47, 5, 0, cst))
			c.So(msg.MsgType, c.ShouldEqual, MessageTypeText)
			c.So(msg.MsgID, c.ShouldEqual, 2018405441)
			c.So(msg.AgentID, c.ShouldEqual, 1000002)

			{
				e, ok := msg.Text()
				c.So(ok, c.ShouldBeTrue)
				c.So(e, c.ShouldNotBeNil)
				c.So(e.GetContent(), c.ShouldEqual, "x123")
			}

			{
				e, ok := msg.Image()
				c.So(ok, c.ShouldBeFalse)
				c.So(e, c.ShouldBeNil)
			}

			{
				e, ok := msg.Voice()
				c.So(ok, c.ShouldBeFalse)
				c.So(e, c.ShouldBeNil)
			}

			{
				e, ok := msg.Video()
				c.So(ok, c.ShouldBeFalse)
				c.So(e, c.ShouldBeNil)
			}

			{
				e, ok := msg.Location()
				c.So(ok, c.ShouldBeFalse)
				c.So(e, c.ShouldBeNil)
			}

			{
				e, ok := msg.Link()
				c.So(ok, c.ShouldBeFalse)
				c.So(e, c.ShouldBeNil)
			}
		})
	})
}
