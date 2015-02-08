package myfitnesspal

import "fmt"

const (
	DateFormat = "2006-01-02"
)

var (
	ErrNotLoggedIn = fmt.Errorf("not logged in")
)

type Macros struct {
	Section  string `json:"-"`
	Label    string `json:"label"`
	Calories int    `json:"calories"`
	Carbs    int    `json:"carbs"`
	Fat      int    `json:"fat"`
	Protein  int    `json:"protein"`
	Sodium   int    `json:"sodium"`
	Sugar    int    `json:"sugar"`
}

type DiaryEntry struct {
	Breakfast MacrosArray `json:"breakfast,omitempty"`
	Lunch     MacrosArray `json:"lunch,omitempty"`
	Dinner    MacrosArray `json:"dinner,omitempty"`
	Snacks    MacrosArray `json:"snacks,omitempty"`

	Totals    *Macros `json:"totals"`
	Goal      *Macros `json:"goal"`
	Remaining *Macros `json:"remaining"`
}

type MacrosArray []*Macros

func (m MacrosArray) Totals() *Macros {
	macros := &Macros{}

	if m == nil || len(m) == 0 {
		return macros
	}

	macros.Section = m[0].Section
	macros.Label = m[0].Section + " Totals"
	for _, item := range m {
		macros.Calories = macros.Calories + item.Calories
		macros.Carbs = macros.Carbs + item.Carbs
		macros.Fat = macros.Fat + item.Fat
		macros.Protein = macros.Protein + item.Protein
		macros.Sodium = macros.Sodium + item.Sodium
		macros.Sugar = macros.Sugar + item.Sugar
	}

	return macros
}

func (m MacrosArray) Find(section, label string) *Macros {
	items := m.FindAll(section)

	for _, item := range items {
		if item.Label == label {
			return item
		}
	}

	return nil
}

func (m MacrosArray) FindAll(section string) MacrosArray {
	results := MacrosArray{}

	for _, item := range m {
		if item.Section == section {
			results = append(results, item)
		}
	}

	return results
}
