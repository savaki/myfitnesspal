package myfitnesspal

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFoodDiary(t *testing.T) {
	Convey("Given the food diary page", t, func() {
		html := `
		<table>
	    <tr class="total">
        <td class="first">Totals</td>
          <td>100</td>
          <td>200</td>
          <td>300</td>
          <td>400</td>
          <td>500</td>
          <td>600</td>
        <td class="empty"></td>
      </tr>
      <tr class="total alt">
        <td class="first">Your Daily Goal </td>
          <td>
          1,234</td>
          <td>
          2,345</td>
          <td>
          3,456</td>
          <td>
          4,567</td>
          <td>
          5,678</td>
          <td>
          6,789</td>
        <td class="empty"></td>
      </tr>
      <tr class="total remaining">
      	<td class="first">Remaining</td>
        <td class="positive">11</td>
        <td class="positive">22</td>
        <td class="positive">33</td>
        <td class="positive">44</td>
        <td class="positive">55</td>
        <td class="positive">66</td>
      	<td class="empty"></td>
      </tr>
    </table>`

		Convey("When I call #parseFoodDiary", func() {
			entry, err := parseFoodDiary(strings.NewReader(html))
			So(err, ShouldBeNil)
			So(entry, ShouldNotBeNil)

			Convey("Then I expect #Totals to be set correctly", func() {
				So(entry.Totals, ShouldNotBeNil)
				So(entry.Totals.Calories, ShouldEqual, 100)
				So(entry.Totals.Carbs, ShouldEqual, 200)
				So(entry.Totals.Fat, ShouldEqual, 300)
				So(entry.Totals.Protein, ShouldEqual, 400)
				So(entry.Totals.Sodium, ShouldEqual, 500)
				So(entry.Totals.Sugar, ShouldEqual, 600)
			})

			Convey("Then I expect #Goal to be set correctly", func() {
				So(entry.Goal, ShouldNotBeNil)
				So(entry.Goal.Calories, ShouldEqual, 1234)
				So(entry.Goal.Carbs, ShouldEqual, 2345)
				So(entry.Goal.Fat, ShouldEqual, 3456)
				So(entry.Goal.Protein, ShouldEqual, 4567)
				So(entry.Goal.Sodium, ShouldEqual, 5678)
				So(entry.Goal.Sugar, ShouldEqual, 6789)
			})

			Convey("Then I expect #Remaining to be set correctly", func() {
				So(entry.Remaining, ShouldNotBeNil)
				So(entry.Remaining.Calories, ShouldEqual, 11)
				So(entry.Remaining.Carbs, ShouldEqual, 22)
				So(entry.Remaining.Fat, ShouldEqual, 33)
				So(entry.Remaining.Protein, ShouldEqual, 44)
				So(entry.Remaining.Sodium, ShouldEqual, 55)
				So(entry.Remaining.Sugar, ShouldEqual, 66)
			})
		})
	})
}
