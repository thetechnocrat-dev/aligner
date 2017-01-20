package aligner

func Tester(str string) string {
	return str
}

func Align(seq1, seq2 string, matchScore, mismatchPenalty, gapPenalty,
	gapOpeningPenalty float64) []string {

	type node struct {
		Score float64
		Path  string
	}

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
		dpTable[i][0] = node{gapPenalty * float64(i), "up"}
	}
	for j := 1; j < colCount; j++ {
		dpTable[0][j] = node{gapPenalty * float64(j), "right"}
	}

	// fill in Scores
	for i := 1; i < rowCount; i++ {
		for j := 1; j < colCount; j++ {
			var pairScore float64
			if seq1[i-1] == seq2[j-1] {
				pairScore = matchScore
			} else {
				pairScore = mismatchPenalty
			}

			rightScore := dpTable[i][j-1].Score + gapPenalty
			upScore := dpTable[i-1][j].Score + gapPenalty
			diagScore := dpTable[i-1][j-1].Score + pairScore

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
		if currentNode.Path == "up" {
			result = append(result, string(seq1[i])+"_")
			i -= 1
		} else if currentNode.Path == "right" {
			result = append(result, "_"+string(seq2[j]))
			j -= 1
		} else if currentNode.Path == "diag" {
			result = append(result, string(seq1[i])+string(seq2[j]))
			i -= 1
			j -= 1
		} else if currentNode.Path == "start" {
			done = true
		}
	}

	// reverse the result array for correct answer
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
