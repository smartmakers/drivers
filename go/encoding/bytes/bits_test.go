package bytes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBits0x01(t *testing.T) {
	Convey("Given the byte 0x01", t, func() {
		in := byte(0x01)

		Convey("When extracting the first 1", func() {
			out, err := Bits(in, 1, 1)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x01", func() {
				So(out, ShouldEqual, 0x01)
			})
		})

		Convey("When extractring the second bit", func() {
			out, err := Bits(in, 2, 2)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x00", func() {
				So(out, ShouldEqual, 0x00)
			})
		})
	})
}

func TestBits0x02(t *testing.T) {
	Convey("Given the byte 0x02", t, func() {
		in := byte(0x02)

		Convey("When extracting the first bit", func() {
			out, err := Bits(in, 1, 1)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x00", func() {
				So(out, ShouldEqual, 0x00)
			})
		})

		Convey("When extracting the second bit", func() {
			out, err := Bits(in, 2, 2)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x01", func() {
				So(out, ShouldEqual, 0x01)
			})
		})
	})
}

func TestBits0x03(t *testing.T) {
	Convey("Given the byte 0x03", t, func() {
		in := byte(0x03)

		Convey("When extracting the first bit", func() {
			out, err := Bits(in, 1, 1)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x01", func() {
				So(out, ShouldEqual, 0x01)
			})
		})

		Convey("When extracting the second bit", func() {
			out, err := Bits(in, 2, 2)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x01", func() {
				So(out, ShouldEqual, 0x01)
			})
		})

		Convey("When extracting the third bit", func() {
			out, err := bits(in, 3, 3)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x00", func() {
				So(out, ShouldEqual, 0x00)
			})
		})

		Convey("When extracting the first and second bit", func() {
			out, err := Bits(in, 1, 2)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x03", func() {
				So(out, ShouldEqual, 0x03)
			})
		})
	})
}

func TestBits0xA0(t *testing.T) {
	Convey("Given the byte 0x80 (0b10...0)", t, func() {
		in := byte(0x80)

		Convey("When extracting the first bit", func() {
			out, err := Bits(in, 1, 1)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x00", func() {
				So(out, ShouldEqual, 0x00)
			})
		})

		Convey("When extracting the second to last bit", func() {
			out, err := Bits(in, 7, 7)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x00", func() {
				So(out, ShouldEqual, 0x00)
			})
		})

		Convey("When extracting the last bit", func() {
			out, err := Bits(in, 8, 8)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x01", func() {
				So(out, ShouldEqual, 0x01)
			})
		})

		Convey("When extracting the two last bits", func() {
			out, err := Bits(in, 7, 8)
			So(err, ShouldBeNil)

			Convey("Then the result should be 0x02", func() {
				So(out, ShouldEqual, 0x02)
			})
		})

	})
}
