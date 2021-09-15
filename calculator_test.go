package calculator_test

import (
	"testing"

	"github.com/wayofthebadger/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 3)
	if want != got {
		t.Errorf("want %f, got %f, want, got")
	}
}
