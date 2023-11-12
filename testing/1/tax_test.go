package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Tax should be: %.3f, got: %.2f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Tax should be: %.3f, got: %.2f", item.expect, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	amount := 500.0
	expected := 5.0
	for i := 0; i < b.N; i++ {
		result := CalculateTax(amount)
		if result != expected {
			b.Errorf("Tax should be: %.3f, got: %.2f", expected, result)
		}
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	amount := 500.0
	expected := 5.0
	for i := 0; i < b.N; i++ {
		result := CalculateTax2(amount)
		if result != expected {
			b.Errorf("Tax should be: %.3f, got: %.2f", expected, result)
		}
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 0, 0.5, 1, 1.5, 2, 2.5, 3, 500.0, 1000.0, 1501.0}

	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, a float64) {
		result := CalculateTax(a)

		if a < 0 && result != 0.0 {
			t.Errorf("Tax should be: %.3f, got: %.2f", 0.0, result)
		}
		if a > 20000 && result != 20.0 {
			t.Errorf("Tax should be: %.3f, got: %.2f", 20.0, result)
		}
	})

}
