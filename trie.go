package autocomplete

import (
	"strings"
)

const AlphabetLength = 26

// Node represents a node in the tree
type Node struct {
	// children holds up to 26 child nodes, one per letter of the alphabet
	children [AlphabetLength]*Node

	// isEnd marks whether this node is the last character of a stored word
	isEnd bool

	// frequency is used to determine how frequent a word has been searched
	// for ranking purposes
	frequency int
}

// Result helps to sort for frequency in a structured way
type Result struct {
	Word      string
	Frequency int
}

// Trie represents the data structure which stores
// all nodes
type Trie struct {
	root *Node
}

// InitTrie initializes a Trie with it's node and
// sets the children with nil indexes
func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert adds a word to the trie
func (t *Trie) Insert(w string) {
	// normalize input for case-insensitive comparison
	w = strings.TrimSpace(strings.ToLower(w))

	wl := len(w)
	currentNode := t.root

	for i := range wl {
		// this makes the word have a character index from
		// 0 - 26 instead of byte char (97...)
		charIdx := w[i] - 'a'

		// traverse to the child node for this character, creating it if it doesn't exist
		if currentNode.children[charIdx] == nil {
			currentNode.children[charIdx] = &Node{}
		}
		currentNode = currentNode.children[charIdx]
	}
	currentNode.isEnd = true
}

// Search browses the trie and returns true if there is a match
func (t *Trie) Search(w string) bool {
	// normalize input for case-insensitive comparison
	w = strings.TrimSpace(strings.ToLower(w))

	wl := len(w)
	currentNode := t.root

	for i := range wl {
		// this makes the word have a character index from
		// 0 - 26 instead of byte char (97...)
		charIdx := w[i] - 'a'

		// traverse to the child node for this character, returning
		// false if it does not exist
		if currentNode.children[charIdx] == nil {
			return false
		}
		currentNode = currentNode.children[charIdx]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

// IncrementFrequency increments the frequency for a word searched
func (t *Trie) IncrementFrequency(w string) {
	w = strings.TrimSpace(strings.ToLower(w))

	wl := len(w)
	currentNode := t.root

	for i := range wl {
		charIdx := w[i] - 'a'

		if currentNode.children[charIdx] == nil {
			return
		}
		// move to the next node
		currentNode = currentNode.children[charIdx]
	}
	if currentNode.isEnd {
		currentNode.frequency++
	}
}

// GetAllWords returns all words in the trie
func (t *Trie) GetAllWords() []Result {
	words := []Result{}
	collectWords(t.root, "", &words)
	return words
}

// Collect traverses through each child node and saves that word
func collectWords(n *Node, currentWord string, words *[]Result) {
	// when reach a node which isEnd is true add it to the list of words
	if n.isEnd {
		res := Result{
			Word:      currentWord,
			Frequency: n.frequency,
		}
		*words = append(*words, res)
	}
	// recursively go through each node adding it's character to the currentWord
	for i := range n.children {
		if n.children[i] != nil {
			collectWords(n.children[i], currentWord+string(rune('a'+i)), words)
		}
	}
}
