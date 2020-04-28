package bdd_demo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConvey(t *testing.T) {
	Convey("Given birth",t , func() {
		a := 2
		b := 4

		Convey("When a difference b", func() {
			c := b - a

			Convey("Then the result is bigger than zero", func() {
				So(c, ShouldBeGreaterThanOrEqualTo, 0)
			})
		})
	})
}