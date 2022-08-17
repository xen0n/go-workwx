package workwx

import (
	"testing"
	"time"

	c "github.com/smartystreets/goconvey/convey"
)

var cst = time.FixedZone("CST", 8*3600)

func TestRxMessageText(t *testing.T) {
	c.Convey("解析接收的 XML 消息体", t, func() {
		c.Convey("文本消息", func() {
			body := []byte("<xml><ToUserName><![CDATA[ww6a112864f8022910]]></ToUserName><FromUserName><![CDATA[foobar]]></FromUserName><CreateTime>1583995625</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[x123]]></Content><MsgId>2018405441</MsgId><AgentID>1000002</AgentID></xml>")

			msg, err := fromEnvelope(body)
			c.So(err, c.ShouldBeNil)
			c.So(msg, c.ShouldNotBeNil)
			c.So(msg.String(), c.ShouldEqual, `RxMessage { FromUserID: "foobar", SendTime: 1583995625000000000, MsgType: "text", MsgID: 2018405441, AgentID: 1000002, Event: "", ChangeType: "", Content: "x123" }`)
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

func TestRxMessageEventEditExternalContact(t *testing.T) {
	c.Convey("解析接收的 XML 消息体", t, func() {
		c.Convey("编辑企业客户事件", func() {
			body := []byte("<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[sys]]></FromUserName> <CreateTime>1403610513</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[change_external_contact]]></Event><ChangeType><![CDATA[edit_external_contact]]></ChangeType><UserID><![CDATA[zhangsan]]></UserID><ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mAAAA]]></ExternalUserID><State><![CDATA[teststate]]></State></xml>")

			msg, err := fromEnvelope(body)
			c.So(err, c.ShouldBeNil)
			c.So(msg, c.ShouldNotBeNil)
			c.So(msg.String(), c.ShouldEqual, `RxMessage { FromUserID: "sys", SendTime: 1403610513000000000, MsgType: "event", MsgID: 0, AgentID: 0, Event: "change_external_contact", ChangeType: "edit_external_contact", UserID: "zhangsan", ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mAAAA", State: "teststate" }`)
			c.So(msg.FromUserID, c.ShouldEqual, "sys")
			c.So(msg.SendTime, c.ShouldEqual, time.Date(2014, 6, 24, 19, 48, 33, 0, cst))
			c.So(msg.MsgType, c.ShouldEqual, MessageTypeEvent)
			c.So(msg.MsgID, c.ShouldEqual, 0)
			c.So(msg.AgentID, c.ShouldEqual, 0)
			c.So(msg.Event, c.ShouldEqual, EventTypeChangeExternalContact)
			c.So(msg.ChangeType, c.ShouldEqual, ChangeTypeEditExternalContact)

			{
				e, ok := msg.EventEditExternalContact()
				c.So(ok, c.ShouldBeTrue)
				c.So(e, c.ShouldNotBeNil)
				c.So(e.GetUserID(), c.ShouldEqual, "zhangsan")
				c.So(e.GetExternalUserID(), c.ShouldEqual, "woAJ2GCAAAXtWyujaWJHDDGi0mAAAA")
				c.So(e.GetState(), c.ShouldEqual, "teststate")
			}
		})
	})
}

func TestRxMessageChangeContactUpdateUser(t *testing.T) {
	c.Convey("解析接收的 XML 消息体", t, func() {
		c.Convey("更新成员事件", func() {
			body := []byte(" <xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[sys]]></FromUserName> <CreateTime>1403610513</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[change_contact]]></Event><ChangeType>update_user</ChangeType><UserID><![CDATA[zhangsan]]></UserID><NewUserID><![CDATA[zhangsan001]]></NewUserID><Name><![CDATA[张三]]></Name><Department><![CDATA[1,2,3]]></Department><MainDepartment>1</MainDepartment><IsLeaderInDept><![CDATA[1,0,0]]></IsLeaderInDept><Position><![CDATA[产品经理]]></Position><Mobile>13800000000</Mobile><Gender>1</Gender><Email><![CDATA[zhangsan@gzdev.com]]></Email><Status>1</Status><Avatar><![CDATA[http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0]]></Avatar><Alias><![CDATA[zhangsan]]></Alias><Telephone><![CDATA[020-123456]]></Telephone><Address><![CDATA[广州市]]></Address><ExtAttr><Item><Name><![CDATA[爱好]]></Name><Type>0</Type><Text><Value><![CDATA[旅游]]></Value></Text></Item><Item><Name><![CDATA[卡号]]></Name><Type>1</Type><Web><Title><![CDATA[企业微信]]></Title><Url><![CDATA[https://work.weixin.qq.com]]></Url></Web></Item></ExtAttr></xml>")

			msg, err := fromEnvelope(body)
			c.So(err, c.ShouldBeNil)

			{
				y, ok := msg.EventChangeTypeUpdateUser()
				c.So(ok, c.ShouldBeTrue)
				c.So(y, c.ShouldNotBeNil)
			}
		})
	})

}

func TestRxMessageChangeContactCreateUser(t *testing.T) {
	c.Convey("解析接收的 XML 消息体", t, func() {
		c.Convey("更新成员事件", func() {
			body := []byte(" <xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[sys]]></FromUserName> <CreateTime>1403610513</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[change_contact]]></Event><ChangeType>create_user</ChangeType><UserID><![CDATA[zhangsan]]></UserID><Name><![CDATA[张三]]></Name><Department><![CDATA[1,2,3]]></Department><MainDepartment>1</MainDepartment><IsLeaderInDept><![CDATA[1,0,0]]></IsLeaderInDept><Position><![CDATA[产品经理]]></Position><Mobile>13800000000</Mobile><Gender>1</Gender><Email><![CDATA[zhangsan@gzdev.com]]></Email><Status>1</Status><Avatar><![CDATA[http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0]]></Avatar><Alias><![CDATA[zhangsan]]></Alias><Telephone><![CDATA[020-123456]]></Telephone><Address><![CDATA[广州市]]></Address><ExtAttr><Item><Name><![CDATA[爱好]]></Name><Type>0</Type><Text><Value><![CDATA[旅游]]></Value></Text></Item><Item><Name><![CDATA[卡号]]></Name><Type>1</Type><Web><Title><![CDATA[企业微信]]></Title><Url><![CDATA[https://work.weixin.qq.com]]></Url></Web></Item></ExtAttr></xml>")

			msg, err := fromEnvelope(body)
			c.So(err, c.ShouldBeNil)

			{
				y, ok := msg.EventChangeTypeCreateUser()
				c.So(ok, c.ShouldBeTrue)
				c.So(y, c.ShouldNotBeNil)
			}
		})
	})
}
