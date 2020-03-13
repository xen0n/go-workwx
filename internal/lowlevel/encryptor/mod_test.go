package encryptor

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestWorkwxEncryptorCtor(t *testing.T) {
	c.Convey("合法的 EncodingAESKey 可以构造出非空的 encryptor", t, func() {
		encodingAESKey := "4Ma3YBrSBbX2aez8MJpXGBne5LSDwgGqHbhM9WPYIws"
		enc, err := NewWorkwxEncryptor(encodingAESKey)
		c.So(err, c.ShouldBeNil)
		c.So(enc, c.ShouldNotBeNil)
	})

	c.Convey("非法的 EncodingAESKey 不能成功构造 encryptor", t, func() {
		c.Convey("长度不对", func() {
			c.Convey("空值", func() {
				encodingAESKey := ""
				enc, err := NewWorkwxEncryptor(encodingAESKey)
				c.So(err, c.ShouldNotBeNil)
				c.So(enc, c.ShouldBeNil)
			})
			c.Convey("len = 42", func() {
				encodingAESKey := "123456789012345678901234567890123456789012"
				enc, err := NewWorkwxEncryptor(encodingAESKey)
				c.So(err, c.ShouldNotBeNil)
				c.So(enc, c.ShouldBeNil)
			})
			c.Convey("len = 44", func() {
				encodingAESKey := "12345678901234567890123456789012345678901234"
				enc, err := NewWorkwxEncryptor(encodingAESKey)
				c.So(err, c.ShouldNotBeNil)
				c.So(enc, c.ShouldBeNil)
			})
		})

		c.Convey("非法的 Base64", func() {
			c.Convey("多余的 '='", func() {
				encodingAESKey := "4Ma3YBrSBbX2aez8MJpXGBne5LSDwgGqHbhM9WPYIw="
				enc, err := NewWorkwxEncryptor(encodingAESKey)
				c.So(err, c.ShouldNotBeNil)
				c.So(enc, c.ShouldBeNil)
			})
			c.Convey("非法字符", func() {
				encodingAESKey := "4Ma3YBrSBbX2aez8MJpXGBne5LSDwgGqHbhM9WPYIw$"
				enc, err := NewWorkwxEncryptor(encodingAESKey)
				c.So(err, c.ShouldNotBeNil)
				c.So(enc, c.ShouldBeNil)
			})
		})
	})
}

func TestWorkwxEncryptor(t *testing.T) {
	encodingAESKey := "4Ma3YBrSBbX2aez8MJpXGBne5LSDwgGqHbhM9WPYIws"
	enc, err := NewWorkwxEncryptor(encodingAESKey)
	if err != nil {
		panic(err)
	}

	c.Convey("解密", t, func() {
		base64EncryptedMsg := []byte("6KmUQuPVu7UhjyVqRdbo5SfcRqaHvbUlKSHFvBV2ZuR6TIlKsygcfeSd1GDplg1C5KSKr6UPHCaC/nIX3ZNt9w==")
		payload, err := enc.Decrypt(base64EncryptedMsg)
		c.So(err, c.ShouldBeNil)

		expected := WorkwxPayload{
			Msg:       []byte("94966531020182955848408"),
			ReceiveID: []byte("ww6a112864f8022910"),
		}
		c.So(payload, c.ShouldResemble, expected)
	})

	c.Convey("round-trip", t, func() {
		original := WorkwxPayload{
			Msg:       []byte("foobarbaz123456788"),
			ReceiveID: []byte("ww6a112864f8022910"),
		}

		encrypted, err := enc.Encrypt(&original)
		c.So(err, c.ShouldBeNil)

		decrypted, err := enc.Decrypt([]byte(encrypted))
		c.So(err, c.ShouldBeNil)

		c.So(decrypted, c.ShouldResemble, original)
	})
}
