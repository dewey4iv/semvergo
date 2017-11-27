package semvergo

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFromString(t *testing.T) {
	Convey("When calling FromString", t, func() {
		tests := map[string]struct {
			input  string
			output *Version
			err    error
		}{
			"simple passing": {
				input:  "0.0.1",
				output: &Version{0, 0, 1, "", ""},
				err:    nil,
			},
			"alpha branch": {
				input:  "0.0.1-alpha",
				output: &Version{0, 0, 1, "alpha", ""},
				err:    nil,
			},
		}

		for k, v := range tests {
			s := "<nil>"

			if v.output != nil {
				s = v.output.String()
			}

			Convey(fmt.Sprintf("Given case: %s - %s", k, s), func() {
				ver, err := fromString(v.input)

				So(err, ShouldEqual, v.err)

				if v.output != nil {
					So(ver, ShouldNotBeNil)
					So(ver.String(), ShouldEqual, v.output.String())
				} else {
					So(ver, ShouldBeNil)
				}
			})
		}
	})
}
