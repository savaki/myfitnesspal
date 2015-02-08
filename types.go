package myfitnesspal

import "fmt"

const (
	DateFormat = "2006-01-02"
)

var (
	ErrNotLoggedIn = fmt.Errorf("not logged in")
)

type Macros struct {
	Calories int `json:"calories"`
	Carbs    int `json:"carbs"`
	Fat      int `json:"fat"`
	Protein  int `json:"protein"`
	Sodium   int `json:"sodium"`
	Sugar    int `json:"sugar"`
}

type DiaryEntry struct {
	Totals    *Macros `json:"totals"`
	Goal      *Macros `json:"goal"`
	Remaining *Macros `json:"remaining"`
}
