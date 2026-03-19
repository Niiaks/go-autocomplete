package autocomplete

import "math"

// LevenshteinDistance calculates the minimum number of edits (insertions, deletions,
// substitutions) needed to transform w1 into w2
func LevenshteinDistance(w1, w2 string) int {

	rows := len(w1) + 1
	cols := len(w2) + 1

	matrix := make([][]int, rows)

	// allocate columns for each row
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// initialize the first row and column
	for i := range rows {
		matrix[i][0] = i
	}
	for j := range cols {
		matrix[0][j] = j
	}

	// iterate through columns and rows to fill in the distances
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			cost := 0
			if w1[i-1] != w2[j-1] {
				cost = 1 // characters differ so substitution cost is ``
			}

			deletion := matrix[i-1][j] + 1
			insertion := matrix[i][j-1] + 1
			substitution := matrix[i-1][j-1] + cost

			// take the cheapest operation
			matrix[i][j] = int(math.Min(math.Min(float64(deletion), float64(insertion)), float64(substitution)))
		}
	}
	// bottom right cell contains the distance
	return matrix[rows-1][cols-1]
}
