package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Niiaks/go-autocomplete"
)

func loadWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}

func main() {
	trie := autocomplete.InitTrie()

	words, err := loadWords("words.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, w := range words {
		trie.Insert(w)
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Search: ")
	for scanner.Scan() {
		input := scanner.Text()
		results := trie.FuzzySearch(input)
		fmt.Fprint(os.Stdout, "suggestions ", results)
		fmt.Print("Search: ")
	}
}
