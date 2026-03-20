package autocomplete

import "strings"

// searchWithPrefix returns all words in the trie that start with w
func (t *Trie) searchWithPrefix(w string) []string {
	w = strings.TrimSpace(strings.ToLower(w))

	currentNode := t.root
	words := []string{}

	for i := range w {
		charIdx := w[i] - 'a'

		// traverse through each child node, returning an empty
		// slice if it does not
		if currentNode.children[charIdx] == nil {
			return []string{}
		}
		currentNode = currentNode.children[charIdx]
	}
	collectWords(currentNode, w, &words)
	return words
}
