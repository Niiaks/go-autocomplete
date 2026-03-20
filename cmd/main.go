package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Niiaks/go-autocomplete"
)

func main() {
	trie := autocomplete.InitTrie()

	// mock words for trie insertion
	words := []string{
		"car", "cart", "card", "care", "careful",
		"cut", "cute", "cup", "curl",
		"search", "searcher", "searched",
		"go", "golang", "gopher",
		"tree", "trie", "traverse",
		"apple", "application", "apply",
		"ball", "balloon", "band",
		"cat", "catch", "category",
	}

	for _, w := range words {
		trie.Insert(w)
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Search: ")
	for scanner.Scan() {
		input := scanner.Text()
		results := trie.FuzzySearch(input)
		fmt.Println("suggestions ", results)
		fmt.Print("Search: ")
	}
}
