package autocomplete

import (
	"slices"
	"testing"
)

func TestFuzzySearch(t *testing.T) {

	trie := InitTrie()

	trie.Insert("cat")
	trie.Insert("catch")
	trie.Insert("category")
	trie.Insert("search")

	t.Run("should return slice of similar words", func(t *testing.T) {
		got := trie.FuzzySearch("cat")
		want := []string{"cat", "catch", "category"}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("should return similar words if there is a misspelling", func(t *testing.T) {
		got := trie.FuzzySearch("serch")
		want := []string{"search"}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
