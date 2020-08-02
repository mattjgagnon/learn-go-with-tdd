package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertExpectedRepeat := func(t *testing.T, repeated, expected string) {
		t.Helper()

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("test repeat character zero times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertExpectedRepeat(t, repeated, expected)
	})

	t.Run("test repeat character five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertExpectedRepeat(t, repeated, expected)
	})

	t.Run("test repeat character ten times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		assertExpectedRepeat(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeat := Repeat("a", 7)
	fmt.Println(repeat)
	// output: aaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
