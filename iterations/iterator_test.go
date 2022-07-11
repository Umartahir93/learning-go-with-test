package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeatedMessage := func(t testing.TB, repeated string, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		repeatedMessage(t, repeated, expected)
	})

	t.Run("caller specify how many times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		repeatedMessage(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
