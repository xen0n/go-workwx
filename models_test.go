package workwx

import (
	"encoding/json"
	"net/url"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestReqAccessToken(t *testing.T) {
	c.Convey("构造一个 reqAccessToken", t, func() {
		a := reqAccessToken{
			CorpID:     "foo",
			CorpSecret: "bar",
		}

		c.Convey("序列化结果应该符合预期", func() {
			v := a.intoURLValues()

			expected := url.Values{
				"corpid":     []string{"foo"},
				"corpsecret": []string{"bar"},
			}
			c.So(v, c.ShouldResemble, expected)
		})
	})
}

func TestRespCommon(t *testing.T) {
	payloadOk := []byte(`{"errcode":0,"errmsg":"ok"}`)
	payloadErr := []byte(`{"errcode":40014,"errmsg":"invalid access_token"}`)

	c.Convey("构造一个成功响应的 respCommon", t, func() {
		var a respCommon
		err := json.Unmarshal(payloadOk, &a)

		c.Convey("应该能成功反序列化", func() {
			c.So(err, c.ShouldBeNil)

			c.Convey("反序列化之后字段应该正确对应", func() {
				c.So(a.ErrCode, c.ShouldEqual, 0)
				c.So(a.ErrMsg, c.ShouldEqual, "ok")
			})

			c.Convey("IsOk 应该为真", func() {
				c.So(a.IsOK(), c.ShouldBeTrue)
			})
		})
	})

	c.Convey("构造一个失败响应的 respCommon", t, func() {
		var a respCommon
		err := json.Unmarshal(payloadErr, &a)

		c.Convey("应该能成功反序列化", func() {
			c.So(err, c.ShouldBeNil)

			c.Convey("反序列化之后字段应该正确对应", func() {
				c.So(a.ErrCode, c.ShouldEqual, 40014)
				c.So(a.ErrMsg, c.ShouldEqual, "invalid access_token")
			})

			c.Convey("IsOk 应该为假", func() {
				c.So(a.IsOK(), c.ShouldBeFalse)
			})
		})
	})
}

func TestReqMessage(t *testing.T) {
	c.Convey("构造 reqMessage", t, func() {
		content := map[string]interface{}{"content": "test"}
		a := reqMessage{
			AgentID: 233,
			MsgType: "text",
			Content: content,
		}

		c.Reset(func() {
			a.ToUser = nil
			a.ToParty = nil
			a.ToTag = nil
			a.ChatID = ""
			a.Content = content
			a.IsSafe = false
		})

		c.Convey("故意放一个不能 marshal 的 Content", func() {
			a.Content = map[string]interface{}{
				"heh": make(chan struct{}),
			}

			c.Convey("执行序列化", func() {
				result, err := a.intoBody()

				c.Convey("序列化应该失败", func() {
					c.So(err, c.ShouldNotBeNil)
					c.So(result, c.ShouldBeNil)
				})
			})
		})

		c.Convey("发给用户列表 & 设置为保密信息", func() {
			a.ToUser = []string{"foo", "bar", "baz"}
			a.IsSafe = true

			c.Convey("执行序列化", func() {
				result, err := a.intoBody()

				c.Convey("序列化应该成功", func() {
					c.So(err, c.ShouldBeNil)

					c.Convey("序列化结果应该符合预期", func() {
						expectedPayload := []byte(`{
								"touser": "foo|bar|baz",
								"toparty": "",
								"totag": "",
								"msgtype": "text",
								"agentid": 233,
								"text": {"content": "test"},
								"safe": 1
								}`)
						var expected map[string]interface{}
						err := json.Unmarshal(expectedPayload, &expected)
						c.So(err, c.ShouldBeNil)

						var actual map[string]interface{}
						err = json.Unmarshal(result, &actual)
						c.So(err, c.ShouldBeNil)

						// we're comparing JSON *outputs*
						// so assertions.ShouldEqualJSON is not suitable
						c.So(actual, c.ShouldResemble, expected)
					})
				})
			})
		})

		c.Convey("发给 chatid", func() {
			a.ChatID = "quux"

			c.Convey("执行序列化", func() {
				result, err := a.intoBody()

				c.Convey("序列化应该成功", func() {
					c.So(err, c.ShouldBeNil)

					c.Convey("序列化结果应该符合预期", func() {
						expectedPayload := []byte(`{
								"chatid": "quux",
								"msgtype": "text",
								"agentid": 233,
								"text": {"content": "test"},
								"safe": 0
								}`)
						var expected map[string]interface{}
						err := json.Unmarshal(expectedPayload, &expected)
						c.So(err, c.ShouldBeNil)

						var actual map[string]interface{}
						err = json.Unmarshal(result, &actual)
						c.So(err, c.ShouldBeNil)

						// we're comparing JSON *outputs*
						// so assertions.ShouldEqualJSON is not suitable
						c.So(actual, c.ShouldResemble, expected)
					})
				})
			})
		})
	})
}
