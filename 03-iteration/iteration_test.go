package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat n times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
