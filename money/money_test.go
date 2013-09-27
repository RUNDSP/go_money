package money

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestMoneySpecs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in benchmark mode.")
		return
	}
	r := gospec.NewRunner()
	r.AddSpec(MoneySpecs)
	gospec.MainGoTest(r, t)
}

func MoneySpecs(c gospec.Context) {

	//
	// Factory Methods
	//

	c.Specify("[Money][MakeMicrosPer1x] Returns Micros", func() {
		value := MakeMoneyMicrosPer1x(123456789)
		c.Expect(value.MicrosPer1x, gospec.Equals, int64(123456789))
	})

	c.Specify("[Money][MakeCostPerMille] Returns CPM", func() {
		value := MakeMoneyCostPerMille(float64(0.50))
		c.Expect(value.MicrosPer1x, gospec.Equals, int64(500))
	})

	c.Specify("[Money][MakeCPM] Formats CPM", func() {
		value := MakeMoneyCPM(float64(123456.789))
		c.Expect(value.MicrosPer1x, gospec.Equals, int64(123456789))
	})

	c.Specify("[Money][MakeDollarsPer1x] Formats Dollars", func() {
		value := MakeMoneyDollarsPer1x(float64(123.456789))
		c.Expect(value.MicrosPer1x, gospec.Equals, int64(123456789))
	})

	//
	// Conversion Utilities:
	//

	c.Specify("[Money][MicrosPer1x] Returns Micros", func() {
		value := &Money{123456789}
		c.Expect(value.MicrosPer1x, gospec.Equals, int64(123456789))
	})

	c.Specify("[Money][CostPerMille] Returns CPM", func() {
		value := &Money{123456789}
		c.Expect(value.CostPerMille(), gospec.Equals, float64(123456.789))
	})

	c.Specify("[Money][CPM] Formats CPM", func() {
		value := &Money{123456789}
		c.Expect(value.CPM(), gospec.Equals, float64(123456.789))
	})

	c.Specify("[Money][DollarsPer1x] Formats Dollars", func() {
		value := &Money{123456789}
		c.Expect(value.DollarsPer1x(), gospec.Equals, float64(123.456789))
	})

	//
	// String Format Utilities:
	//

	c.Specify("[Money][String] Formats string", func() {
		value := &Money{123456789}
		c.Expect(value.String(), gospec.Equals, "µ123456789(1x)")
	})

	c.Specify("[Money][CostPerMilleStr] Formats string", func() {
		value := &Money{123456789}
		c.Expect(value.CostPerMilleStr(), gospec.Equals, "$123456.789000(CPM)")
	})

	c.Specify("[Money][CPMStr] Formats string", func() {
		value := &Money{123456789}
		c.Expect(value.CPMStr(), gospec.Equals, "$123456.789000(CPM)")
	})

	c.Specify("[Money][DollarsPer1xStr] Formats string", func() {
		value := &Money{123456789}
		c.Expect(value.DollarsPer1xStr(), gospec.Equals, "$123.456789(1x)")
	})

}

func Benchmark_MakeMicrosPer1x(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeMoneyMicrosPer1x(123456789)
	}
}

func Benchmark_MakeCostPerMille(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeMoneyCostPerMille(float64(123.456789))
	}
}

func Benchmark_MakeCPM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeMoneyCPM(float64(123456.789))
	}
}

func Benchmark_MakeDollarsPer1x(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeMoneyDollarsPer1x(float64(123.456789))
	}
}
