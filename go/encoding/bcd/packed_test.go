package bcd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPackedBigEndianBCDDecoder(t *testing.T) {
	Convey("Given a packed big-endian BCD decoder", t, func() {
		var bcd PackedBigEndianBCD

		Convey("When decoding 0x00", func() {
			err := bcd.UnmarshalBinary([]byte{0x00})
			So(err, ShouldBeNil)

			Convey("Then the result should be 0", func() {
				So(bcd, ShouldEqual, 0)
			})
		})

		Convey("When decoding 0x01", func() {
			err := bcd.UnmarshalBinary([]byte{0x01})
			So(err, ShouldBeNil)

			Convey("Then the result should be 1", func() {
				So(bcd, ShouldEqual, 1)
			})
		})

		Convey("When decoding 0x09", func() {
			err := bcd.UnmarshalBinary([]byte{0x09})
			So(err, ShouldBeNil)

			Convey("Then the result should be 9", func() {
				So(bcd, ShouldEqual, 9)
			})
		})

		Convey("When decoding 0x0A", func() {
			err := bcd.UnmarshalBinary([]byte{0x0A})

			Convey("Then an error should happen", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When decoding 0x10", func() {
			err := bcd.UnmarshalBinary([]byte{0x10})
			So(err, ShouldBeNil)

			Convey("Then the result should be 10", func() {
				So(bcd, ShouldEqual, 10)
			})
		})

		Convey("When decoding 0xA0", func() {
			err := bcd.UnmarshalBinary([]byte{0xA0})

			Convey("Then an error should happen", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When decoding 0x42", func() {
			err := bcd.UnmarshalBinary([]byte{0x42})
			So(err, ShouldBeNil)

			Convey("Then the result should be 42", func() {
				So(bcd, ShouldEqual, 42)
			})
		})

		Convey("When decoding 0x4711", func() {
			err := bcd.UnmarshalBinary([]byte{0x47, 0x11})
			So(err, ShouldBeNil)

			Convey("Then the result should be 4711", func() {
				So(bcd, ShouldEqual, 4711)
			})
		})

		Convey("When decoding 0x12345678", func() {
			err := bcd.UnmarshalBinary([]byte{0x12, 0x34, 0x56, 0x78})
			So(err, ShouldBeNil)

			Convey("Then the result should be 12345678", func() {
				So(bcd, ShouldEqual, 12345678)
			})
		})
	})
}

func TestPackedLittleEndianBCDDecoder(t *testing.T) {
	Convey("Given a packed little-endian BCD decoder", t, func() {
		var bcd PackedLittleEndianBCD

		Convey("When decoding 0x00", func() {
			err := bcd.UnmarshalBinary([]byte{0x00})
			So(err, ShouldBeNil)

			Convey("Then the result should be 0", func() {
				So(bcd, ShouldEqual, 0)
			})
		})

		Convey("When decoding 0x01", func() {
			err := bcd.UnmarshalBinary([]byte{0x01})
			So(err, ShouldBeNil)

			Convey("Then the result should be 10", func() {
				So(bcd, ShouldEqual, 10)
			})
		})

		Convey("When decoding 0x09", func() {
			err := bcd.UnmarshalBinary([]byte{0x09})
			So(err, ShouldBeNil)

			Convey("Then the result should be 90", func() {
				So(bcd, ShouldEqual, 90)
			})
		})

		Convey("When decoding 0x0A", func() {
			err := bcd.UnmarshalBinary([]byte{0x0A})

			Convey("Then an error should happen", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When decoding 0x10", func() {
			err := bcd.UnmarshalBinary([]byte{0x10})
			So(err, ShouldBeNil)

			Convey("Then the result should be 1", func() {
				So(bcd, ShouldEqual, 1)
			})
		})

		Convey("When decoding 0xA0", func() {
			err := bcd.UnmarshalBinary([]byte{0xA0})

			Convey("Then an error should happen", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When decoding 0x42", func() {
			err := bcd.UnmarshalBinary([]byte{0x42})
			So(err, ShouldBeNil)

			Convey("Then the result should be 24", func() {
				So(bcd, ShouldEqual, 24)
			})
		})

		Convey("When decoding 0x4711", func() {
			err := bcd.UnmarshalBinary([]byte{0x47, 0x11})
			So(err, ShouldBeNil)

			Convey("Then the result should be 1174", func() {
				So(bcd, ShouldEqual, 1174)
			})
		})

		Convey("When decoding 0x12345678", func() {
			err := bcd.UnmarshalBinary([]byte{0x12, 0x34, 0x56, 0x78})
			So(err, ShouldBeNil)

			Convey("Then the result should be 87654321", func() {
				So(bcd, ShouldEqual, 87654321)
			})
		})
	})
}
