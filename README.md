# go-autocomplete

> Simple trie-based autocomplete with fuzzy matching (Levenshtein distance)

This repository provides a small Go library and a tiny CLI demonstrating:

- A Trie implementation for prefix search and storage
- Fuzzy search using Levenshtein distance (returns nearby words)
- Frequency-based ranking of suggestions

## Features

- Fast prefix-based lookups via a Trie (`InitTrie`, `Insert`, `Search`).
- Fuzzy suggestions via `FuzzySearch` (uses `LevenshteinDistance`, currently considers distance ≤ 2).
- Tracks search frequency for basic ranking (`IncrementFrequency`, results are sorted by `Frequency`).
- Small CLI in `cmd/` that loads a word list and provides interactive suggestions.

## Requirements

- Go 1.25+ (module set to `go 1.25.7` in `go.mod`).

## Install / Run

Run the interactive CLI (loads `cmd/words.txt`):

```bash
go run ./cmd
```

Or build the CLI binary:

```bash
go build -o go-autocomplete ./cmd
./go-autocomplete
```

## Use as a library

Import the module and use the `autocomplete` package:

```go
import (
    "fmt"
    "github.com/Niiaks/go-autocomplete"
)

func example() {
    t := autocomplete.InitTrie()
    t.Insert("cat")
    t.Insert("catch")

    // fuzzy search (handles minor misspellings)
    results := t.FuzzySearch("serch")
    for _, r := range results {
        fmt.Println(r.Word, r.Frequency)
    }
}
```

## Public API (overview)

- `InitTrie() *Trie` — create a new Trie instance.
- `(*Trie) Insert(w string)` — insert a word into the trie.
- `(*Trie) Search(w string) bool` — exact word lookup.
- `(*Trie) FuzzySearch(w string) []Result` — fuzzy suggestions (returns `[]Result`).
- `(*Trie) IncrementFrequency(w string)` — increments a word's stored frequency.
- `(*Trie) GetAllWords() []Result` — returns all stored words and frequencies.
- `LevenshteinDistance(w1, w2 string) int` — utility to compute edit distance.
- `Result` — struct `{ Word string; Frequency int }` used for suggestion results.

All library code is in package `autocomplete`.

## Tests

Run the test suite:

```bash
go test ./...
```

The repo includes unit tests for the Trie, fuzzy search, and distance functions.

## Data

- `cmd/words.txt` contains a sample word list used by the CLI. Replace or extend it for larger datasets.

## Limitations & Notes

- The implementation normalizes input using `strings.ToLower()` and assumes English lowercase letters (a–z) when indexing child nodes.
- The data structure is in-memory and not safe for concurrent writes; consider adding synchronization for concurrent use.

## Contributing

Contributions, bug reports and PRs are welcome. Please open issues for design discussions.

## License

This project is licensed under the MIT License — see [LICENSE](LICENSE).
