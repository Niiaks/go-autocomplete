package autocomplete

import (
	"cmp"
	"slices"
)

// FuzzySearch takes in a word and returns related words
// even if there is a mistake in words
func (t *Trie) FuzzySearch(w string) []Result {
	related := []Result{}
	trieWords := t.GetAllWords()

	res := t.searchWithPrefix(w)
	if len(res) > 0 {
		return res
	}

	for _, v := range trieWords {
		distance := LevenshteinDistance(w, v.Word)
		if distance <= 2 {
			related = append(related, v)
		}
	}

	// return results sorted by frequency in descending order
	slices.SortFunc(related, func(a, b Result) int {
		return cmp.Compare(b.Frequency, a.Frequency)
	})
	return related
}
