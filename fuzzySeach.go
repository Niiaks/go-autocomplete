package autocomplete

// FuzzySearch takes in a word and returns related words
// even if there is a mistake in words
func (t *Trie) FuzzySearch(w string) []string {
	related := []string{}
	trieWords := t.GetAllWords()

	words := t.searchWithPrefix(w)
	if len(words) > 0 {
		related = append(related, words...)
	} else {
		for _, word := range trieWords {
			distance := LevenshteinDistance(w, word)
			if distance <= 2 {
				related = append(related, word)
			}
		}
	}
	return related
}
