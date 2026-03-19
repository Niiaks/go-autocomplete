package autocomplete

import "testing"

func TestDistance(t *testing.T) {
	t.Run("insertion", func(t *testing.T) {
		got := LevenshteinDistance("car", "cart")
		want := 1
		assertMessage(t, got, want)
	})

	t.Run("empty strings", func(t *testing.T) {
		got := LevenshteinDistance("", "cat")
		want := 3
		assertMessage(t, got, want)
	})

	t.Run("identical strings", func(t *testing.T) {
		got := LevenshteinDistance("cat", "cat")
		want := 0
		assertMessage(t, got, want)
	})

	t.Run("completely different strings", func(t *testing.T) {
		got := LevenshteinDistance("cap", "dog")
		want := 3
		assertMessage(t, got, want)
	})

	t.Run("single characters (same)", func(t *testing.T) {
		got := LevenshteinDistance("a", "a")
		want := 0
		assertMessage(t, got, want)
	})

	t.Run("deletion", func(t *testing.T) {
		got := LevenshteinDistance("cart", "cat")
		want := 1
		assertMessage(t, got, want)
	})

	t.Run("single characters (different)", func(t *testing.T) {
		got := LevenshteinDistance("a", "b")
		want := 1
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
