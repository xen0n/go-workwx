package pkcs7

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestPKCS7Pad(t *testing.T) {
	c.Convey("PKCS#7 Padding", t, func() {
		c.Convey("len = 0", func() {
			x := []byte{}
			expectedY := []byte{
				32, 32, 32, 32, 32, 32, 32, 32,
				32, 32, 32, 32, 32, 32, 32, 32,
				32, 32, 32, 32, 32, 32, 32, 32,
				32, 32, 32, 32, 32, 32, 32, 32,
			}
			y := Pad(x)
			c.So(y, c.ShouldResemble, expectedY)
		})

		c.Convey("len = 1", func() {
			x := []byte{65}
			expectedY := []byte{
				65, 31, 31, 31, 31, 31, 31, 31,
				31, 31, 31, 31, 31, 31, 31, 31,
				31, 31, 31, 31, 31, 31, 31, 31,
				31, 31, 31, 31, 31, 31, 31, 31,
			}
			y := Pad(x)
			c.So(y, c.ShouldResemble, expectedY)
		})

		c.Convey("len = 31", func() {
			x := []byte("0123456789abcdef0123456789abcde")
			expectedY := []byte{
				48, 49, 50, 51, 52, 53, 54, 55,
				56, 57, 97, 98, 99, 100, 101, 102,
				48, 49, 50, 51, 52, 53, 54, 55,
				56, 57, 97, 98, 99, 100, 101, 1,
			}
			y := Pad(x)
			c.So(y, c.ShouldResemble, expectedY)
		})

		c.Convey("len = 32", func() {
			x := []byte("0123456789abcdef0123456789abcdef")
			expectedY := []byte{
				48, 49, 50, 51, 52, 53, 54, 55,
				56, 57, 97, 98, 99, 100, 101, 102,
				48, 49, 50, 51, 52, 53, 54, 55,
				56, 57, 97, 98, 99, 100, 101, 102,
				32, 32, 32, 32, 32, 32, 32, 32,
				32, 32, 32, 32, 32, 32, 32, 32,
				32, 32, 32, 32, 32, 32, 32, 32,
				32, 32, 32, 32, 32, 32, 32, 32,
			}
			y := Pad(x)
			c.So(y, c.ShouldResemble, expectedY)
		})

		c.Convey("function should have no side-effect", func() {
			x := []byte("foobar")
			expectedX := []byte("foobar")
			_ = Pad(x)
			c.So(x, c.ShouldResemble, expectedX)
		})
	})
}
