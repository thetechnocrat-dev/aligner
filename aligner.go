package aligner

func Tester(str string) string {
	return str
}

// Utility function and structs
func initDpTable(seq1, seq2 string) [][]int {
	dpTable := make([][]int, len(seq1))

	// make rows
	for i := range dpTable {
		dpTable[i] = make([]int, len(seq2))
	}

	// add initial values
	dpTable[0][0] = 0
	for i := 1; i < len(seq1); i++ {
		dpTable[i][0] = -2 * i
	}
	for j := 1; j < len(seq2); j++ {
		dpTable[0][j] = -2 * j
	}

	return dpTable
}

type node struct {
	score int
	path  string
}

func Align(seq1, seq2 string) []string {
	const gap = -2
	const match = 1
	const mismatch = -1

	rowCount := len(seq1) + 1
	colCount := len(seq2) + 1
	dpTable := make([][]node, rowCount)

	// make rows
	for i := range dpTable {
		dpTable[i] = make([]node, colCount)
	}

	// add initial values
	dpTable[0][0] = node{0, "start"}
	for i := 1; i < rowCount; i++ {
		dpTable[i][0] = node{gap * i, "up"}
	}
	for j := 1; j < colCount; j++ {
		dpTable[0][j] = node{gap * j, "right"}
	}

	// fill in scores
	for i := 1; i < rowCount; i++ {
		for j := 1; j < colCount; j++ {
			var pairScore int
			if seq1[i-1] == seq2[j-1] {
				pairScore = match
			} else {
				pairScore = mismatch
			}

			rightScore := dpTable[i][j-1].score + gap
			upScore := dpTable[i-1][j].score + gap
			diagScore := dpTable[i-1][j-1].score + pairScore

			if rightScore >= upScore && rightScore >= diagScore {
				dpTable[i][j] = node{rightScore, "right"}
			} else if upScore >= rightScore && upScore >= diagScore {
				dpTable[i][j] = node{upScore, "up"}
			} else if diagScore >= rightScore && diagScore >= upScore {
				dpTable[i][j] = node{diagScore, "diag"}
			}
		}
	}

	// use dp table to backtrace sequence pairings
	var result []string

	i := len(seq1) - 1
	j := len(seq2) - 1
	done := false
	for !done {
		currentNode := dpTable[i+1][j+1]
		if currentNode.path == "up" {
			result = append(result, string(seq1[i])+"_")
			i -= 1
		} else if currentNode.path == "right" {
			result = append(result, "_"+string(seq2[j]))
			j -= 1
		} else if currentNode.path == "diag" {
			result = append(result, string(seq1[i])+string(seq2[j]))
			i -= 1
			j -= 1
		} else if currentNode.path == "start" {
			done = true
		}
	}

	// reverse the result array for correct answer
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
