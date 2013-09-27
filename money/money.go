package money

import "fmt"

//
// Money:
//
type Money struct {
	microsPer1x int64 "Price in micros per instance"
}

//
// Conversion Constants
//
const (
	CPM_to_Dollars1x      float64 = 1 / 1e3                                  // 0.001
	Dollars1x_to_Micros1x float64 = 1 * 1e6                                  // 1_000_000
	CPM_to_Micros1x       float64 = CPM_to_Dollars1x * Dollars1x_to_Micros1x // 1_000

	Dollars1x_to_CPM      float64 = 1e3                                      // 1_000
	Micros1x_to_Dollars1x float64 = 1 / 1e6                                  // 0.000001
	Micros1x_to_CPM       float64 = Dollars1x_to_CPM * Micros1x_to_Dollars1x // 0.001
)

//
// Factory Methods
//

// Make a new instance of Money from a CPM
// CPM --> Dollars per 1x --> Micros per 1x
func MakeCostPerMille(value float64) *Money {
	// DEBUGING:
	// fmt.Println("MakeCostPerMille", value, value*CPM_to_Dollars1x)

	return MakeDollarsPer1x(value * CPM_to_Dollars1x)
}

// Make a new instance of Money from a CPM
// CPM --> Dollars per 1x --> Micros per 1x
func MakeCPM(value float64) *Money {
	// DEBUGING:
	// fmt.Println("MakeCostPerMille", value)

	return MakeCostPerMille(value)
}

// Make a new instance of Money from (Total) Dollars
// Dollars per 1x --> Micros per 1x
func MakeDollarsPer1x(value float64) *Money {
	// DEBUGING:
	// fmt.Println("MakeDollarsPer1x", value, int64(value*Dollars1x_to_Micros1x))

	return MakeMicrosPer1x(int64(value * Dollars1x_to_Micros1x))
}

// Make a new instance of Money from (Total) Micros
func MakeMicrosPer1x(value int64) *Money {
	// DEBUGING:
	// fmt.Println("MakeMicrosPer1x", value)

	return &Money{value}
}

//
// Conversion Utilities:
//

func (p *Money) MicrosPer1x() int64 {
	return p.microsPer1x
}

// The “cost per thousand advertising impressions” metric (CPM)
// CPM is useful in comparing the relative efficiency of different advertising
// opportunities or media and in evaluating the costs of overall campaigns.
func (p *Money) CostPerMille() float64 {
	return float64(p.MicrosPer1x()) * Micros1x_to_CPM
}

// Cost per mille, the advertising cost per thousand views
func (p *Money) CPM() float64 {
	return p.CostPerMille()
}

// Dollars per 1x instance
func (p *Money) DollarsPer1x() float64 {
	return float64(p.MicrosPer1x()) * Micros1x_to_Dollars1x
}

//
// String Format Utilities:
//

func (p *Money) String() string {
	return fmt.Sprintf("µ%d(1x)", p.MicrosPer1x())
}

func (p *Money) CostPerMilleStr() string {
	return fmt.Sprintf("$%0.6f(CPM)", p.CostPerMille())
}

func (p *Money) CPMStr() string {
	return p.CostPerMilleStr()
}

func (p *Money) DollarsPer1xStr() string {
	return fmt.Sprintf("$%0.6f(1x)", p.DollarsPer1x())
}
