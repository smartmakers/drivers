package bcd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNonPackedBigEndianBCDDecoder(t *testing.T) {
	Convey("Given a non-packed, big-endian BCD", t, func() {
		var bcd NonPackedBigEndianBCD

		Convey("When unmarshaling 0x00", func() {
			err := bcd.UnmarshalBinary([]byte{0x00})
			So(err, ShouldBeNil)

			Convey("Then the result should be 0", func() {
				So(bcd, ShouldEqual, 0)
			})
		})

		Convey("When unmarshaling 0x01", func() {
			err := bcd.UnmarshalBinary([]byte{0x01})
			So(err, ShouldBeNil)

			Convey("Then the result should be 1", func() {
				So(bcd, ShouldEqual, 1)
			})
		})

		Convey("When unmarshaling 0x09", func() {
			err := bcd.UnmarshalBinary([]byte{0x09})
			So(err, ShouldBeNil)

			Convey("Then the result should be 9", func() {
				So(bcd, ShouldEqual, 9)
			})
		})

		Convey("When unmarshaling 0x0A", func() {
			err := bcd.UnmarshalBinary([]byte{0x0A})

			Convey("Then an error should happen", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When unmarshaling 0x10", func() {
			err := bcd.UnmarshalBinary([]byte{0x10})

			Convey("Then an error should happen", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When unmarshaling 0x0102", func() {
			err := bcd.UnmarshalBinary([]byte{0x01, 0x02})
			So(err, ShouldBeNil)

			Convey("Then the result should be 12", func() {
				So(bcd, ShouldEqual, 12)
			})
		})
	})
}

func TestNonPackedLittleEndianBCDDecoder(t *testing.T) {
	Convey("Given a non-packed, little-endian BCD", t, func() {
		var bcd NonPackedLittleEndianBCD

		Convey("When unmarshaling 0x00", func() {
			err := bcd.UnmarshalBinary([]byte{0x00})
			So(err, ShouldBeNil)

			Convey("Then the result should be 0", func() {
				So(bcd, ShouldEqual, 0)
			})
		})

		Convey("When unmarshaling 0x01", func() {
			err := bcd.UnmarshalBinary([]byte{0x01})
			So(err, ShouldBeNil)

			Convey("Then the result should be 1", func() {
				So(bcd, ShouldEqual, 1)
			})
		})

		Convey("When unmarshaling 0x09", func() {
			err := bcd.UnmarshalBinary([]byte{0x09})
			So(err, ShouldBeNil)

			Convey("Then the result should be 09", func() {
				So(bcd, ShouldEqual, 9)
			})
		})

		Convey("When unmarshaling 0x0A", func() {
			err := bcd.UnmarshalBinary([]byte{0x0A})

			Convey("Then an error should happen", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When unmarshaling 0x10", func() {
			err := bcd.UnmarshalBinary([]byte{0x10})

			Convey("Then an error should happen", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When unmarshaling 0x0102", func() {
			err := bcd.UnmarshalBinary([]byte{0x01, 0x02})
			So(err, ShouldBeNil)

			Convey("Then the result should be 21", func() {
				So(bcd, ShouldEqual, 21)
			})
		})
	})
}
