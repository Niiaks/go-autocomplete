package autocomplete

import "testing"

func TestAutoComplete(t *testing.T) {

	trie := InitTrie()
	t.Run("returns true for a match", func(t *testing.T) {
		trie.Insert("master")
		want := true
		got := trie.Search("master")
		assertCorrectMessage(t, got, want)
	})

	t.Run("returns false for no match", func(t *testing.T) {
		trie.Insert("master")
		want := false
		got := trie.Search("masters")
		assertCorrectMessage(t, got, want)
	})

	t.Run("normalizes input", func(t *testing.T) {
		trie.Insert("master")
		want := true
		got := trie.Search("Master")
		assertCorrectMessage(t, got, want)
	})

	t.Run("retrieve all words in trie", func(t *testing.T) {
		trie.Insert("car")
		trie.Insert("cart")
		trie.Insert("cut")

		got := trie.GetAllWords()
		want := 3

		if len(got) != want {
			t.Errorf("want %d got %d", want, len(got))
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("want %t got %t", want, got)
	}
}
