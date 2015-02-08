package myfitnesspal

import (
	"strconv"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDateFormat(t *testing.T) {
	Convey("Given a time.Time instance", t, func() {
		now := time.Date(2014, 4, 5, 6, 0, 0, 0, time.Local)

		Convey("When I call #ToDate", func() {
			dateString := now.Format("2006-01-02")

			Convey("Then I expect a well formatted date string", func() {
				So(dateString, ShouldEqual, "2014-04-05")
			})
		})
	})
}

func TestItoa(t *testing.T) {
	Convey("Given a number => 1,234", t, func() {
		i := 1234

		Convey("When I call #Itoa", func() {
			str := strconv.Itoa(i)

			Convey("When I expect the number to have been parsed successfully", func() {
				So(str, ShouldEqual, "1234")
			})
		})
	})
}
