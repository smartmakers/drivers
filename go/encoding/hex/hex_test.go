package hex

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	test "github.com/smartmakers/drivers/go/testing"
)

func TestUnmarshalTextHex01(t *testing.T) {
	Convey("Given the hex string 0x01", t, func() {
		text := "01"

		Convey("When unmarshaling it", func() {
			h := Hex{}
			err := h.UnmarshalText([]byte(text))
			So(err, ShouldBeNil)

			Convey("Then the result should be the byte 01", func() {
				So(h, ShouldResemble, Hex{0x01})
			})
		})
	})
}

func TestMarshalTextHex01(t *testing.T) {
	Convey("Given the hex value 0x01", t, func() {
		h := Hex{0x01}

		Convey("When unmarshaling it", func() {
			t, err := h.MarshalText()
			So(err, ShouldBeNil)

			Convey("Then the result should be the hex string 01", func() {
				So(string(t), ShouldEqual, "01")
			})
		})
	})
}

func TestUnmarshalJSONHex01(t *testing.T) {
	Convey("Given a json object with a hex field with value 0x0102", t, func() {
		j := `{"hex":"0102"}`

		Convey("When unmarshaling it", func() {
			var obj struct {
				Hex Hex `json:"hex"`
			}

			err := json.Unmarshal([]byte(j), &obj)
			So(err, ShouldBeNil)

			Convey("Then the result should be the byte 0102", func() {
				So(obj.Hex, ShouldResemble, Hex{0x01, 0x02})
			})
		})
	})
}

func TestMarshalJSONHex01(t *testing.T) {
	Convey("Given a json object with a hex field with value 0x0102", t, func() {
		var obj struct {
			Hex Hex `json:"hex"`
		}
		obj.Hex = Hex{0x01, 0x02}

		Convey("When unmarshaling it", func() {

			bytes, err := json.Marshal(obj)
			So(err, ShouldBeNil)

			Convey("Then the result should be the byte 0102", func() {
				test.ExpectJSON(t, `{"hex": "0102"}`, string(bytes))
			})
		})
	})
}
