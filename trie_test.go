package autocomplete

import "testing"

func TestAutoComplete(t *testing.T) {

	tree := InitTrie()
	t.Run("returns true for a match", func(t *testing.T) {
		tree.Insert("master")
		want := true
		got := tree.Search("master")
		assertCorrectMessage(t, got, want)
	})

	t.Run("returns false for no match", func(t *testing.T) {
		tree.Insert("master")
		want := false
		got := tree.Search("masters")
		assertCorrectMessage(t, got, want)
	})

	t.Run("normalizes input", func(t *testing.T) {
		tree.Insert("master")
		want := true
		got := tree.Search("Master")
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("want %t got %t", want, got)
	}
}
