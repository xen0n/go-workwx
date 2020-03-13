package envelope

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestProcessor(t *testing.T) {
	token := "kz7Yx62CH8SaLN"
	encodingAESKey := "cD0d7jx4tYvVtzqrmh3Dm3QFCXe6f8SlHoMtMh3qQEP"
	pr, err := NewProcessor(token, encodingAESKey)
	if err != nil {
		panic(err)
	}

	c.Convey("完整的请求应该能正常解出内容", t, func() {
		s := "http://test.example.com/?msg_signature=f265ae551b1932727204c3d707628d01376a6940&timestamp=1583995625&nonce=1584392382"
		body := "<xml><ToUserName><![CDATA[ww6a112864f8022910]]></ToUserName><Encrypt><![CDATA[EUCt7xMcNiyASzZj0Hjc5yDjFQrCum6AfQ3ntHiUzjGQ51xieKmbvtrZ40/EcB2W/W8yH0n4Lqx48gJl/T9HD/R309I0P/r5pIZucK3lyEn48FYMr4YdE0QdL2jIJ3xkcXUr6uzefzCxG6lMvwpAJaOyVCzN7sRRw47njfxy5EIqU6R9ZBhlTzfdnhhOhK/nTwzrZX3SoGlXFA9OBeZ6ru1NWpXFk76x9DUMe0lcxPPiUqK8ctnQcYXSGUHVqC6DfG7E7mab0OmruNN8cBZY5d3dYOBA4OgaH55Q0AJmUpdT8vNiXpXx+6TxT3TIjySXpDrHVyrsb772aYywgg/Nu4kUmGkALwFZlzhjNegR7wDwb9lr4ERXsSSS8JZ8lbBmaQ3F2Tq584xoPj5rIhXAF734ynm4no1g+SdHiNqR328=]]></Encrypt><AgentID><![CDATA[1000002]]></AgentID></xml>"
		u, err := url.Parse(s)
		c.So(err, c.ShouldBeNil)

		payload, err := pr.HandleIncomingMsg(u, []byte(body))
		c.So(err, c.ShouldBeNil)

		expected := Envelope{
			ToUserName: "ww6a112864f8022910",
			AgentID:    "1000002",
			Msg:        []byte("<xml><ToUserName><![CDATA[ww6a112864f8022910]]></ToUserName><FromUserName><![CDATA[foobar]]></FromUserName><CreateTime>1583995625</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[x123]]></Content><MsgId>2018405441</MsgId><AgentID>1000002</AgentID></xml>"),
			ReceiveID:  []byte("ww6a112864f8022910"),
		}
		c.So(payload, c.ShouldResemble, expected)
	})

	c.Convey("被动响应包应该能正常组装", t, func() {
		originalMsg := []byte("foobarbazquux123321")

		out, err := pr.MakeOutgoingEnvelope(originalMsg)
		c.So(err, c.ShouldBeNil)

		// 重新解出数据
		var outDeserialized struct {
			Encrypt      string `xml:"Encrypt"`
			MsgSignature string `xml:"MsgSignature"`
			Timestamp    int64  `xml:"Timestamp"`
			Nonce        string `xml:"Nonce"`
		}

		err = xml.Unmarshal(out, &outDeserialized)
		c.So(err, c.ShouldBeNil)

		// 做一个假的 URL
		s := fmt.Sprintf(
			"http://a.b/?msg_signature=%s&timestamp=%d&nonce=%s",
			outDeserialized.MsgSignature,
			outDeserialized.Timestamp,
			outDeserialized.Nonce,
		)
		u, err := url.Parse(s)
		c.So(err, c.ShouldBeNil)

		// 亏得 Go encoding 库都对错误的字段名很宽恕
		// 只有 Encrypt 元素有意义，不过正好 ;-)
		rtt, err := pr.HandleIncomingMsg(u, out)
		c.So(err, c.ShouldBeNil)
		c.So(rtt.Msg, c.ShouldResemble, originalMsg)
		c.So(rtt.ReceiveID, c.ShouldBeEmpty)
	})
}
