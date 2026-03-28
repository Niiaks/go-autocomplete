package autocomplete

import (
	"slices"
	"testing"
)

func TestFuzzySearch(t *testing.T) {

	trie := InitTrie()

	trie.Insert("cat")
	trie.Insert("cart")
	trie.Insert("catch")
	trie.Insert("category")
	trie.Insert("search")

	t.Run("should return slice of similar words", func(t *testing.T) {
		got := trie.FuzzySearch("cat")
		want := []Result{
			{Word: "cat", Frequency: 0},
			{Word: "catch", Frequency: 2},
			{Word: "category", Frequency: 1},
		}
		// want := []string{"cat", "catch", "category"}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("should return similar words if there is a misspelling", func(t *testing.T) {
		got := trie.FuzzySearch("serch")
		want := []Result{{Word: "search", Frequency: 1}}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("frequency ranking", func(t *testing.T) {
		trie.IncrementFrequency("cart")
		trie.IncrementFrequency("cart")
		results := trie.FuzzySearch("car")
		if results[0].Word != "cart" {
			t.Errorf("got %s want cart", results[0].Word)
		}
	})
}
