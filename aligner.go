package aligner

import (
	"math"
)

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
	xTable := make([][]node, rowCount)
	yTable := make([][]node, rowCount)
	mTable := make([][]node, rowCount)

	// make rows
	for i := range xTable {
		xTable[i] = make([]node, colCount)
		yTable[i] = make([]node, colCount)
		mTable[i] = make([]node, colCount)
	}

	// add initial values
	xTable[0][0] = node{0, "start"}
	yTable[0][0] = node{0, "start"}
	mTable[0][0] = node{0, "start"}
	for i := 1; i < rowCount; i++ {
		xTable[i][0] = node{math.Inf(-1), "yx"}
		yTable[i][0] = node{gapOpeningPenalty + (gapPenalty * float64(i)), "yy"}
		mTable[i][0] = node{math.Inf(-1), "ym"}
	}
	for j := 1; j < colCount; j++ {
		xTable[0][j] = node{gapOpeningPenalty + (gapPenalty * float64(j)), "xx"}
		yTable[0][j] = node{math.Inf(-1), "xy"}
		mTable[0][j] = node{math.Inf(-1), "xm"}
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

			xx := xTable[i][j-1].Score + gapPenalty
			yx := yTable[i][j-1].Score + gapOpeningPenalty + gapPenalty
			mx := mTable[i][j-1].Score + gapOpeningPenalty + gapPenalty
			if xx >= yx && xx >= mx {
				xTable[i][j] = node{xx, "xx"}
			} else if yx >= mx {
				xTable[i][j] = node{yx, "yx"}
			} else {
				xTable[i][j] = node{mx, "mx"}
			}

			yy := yTable[i-1][j].Score + gapPenalty
			xy := xTable[i-1][j].Score + gapOpeningPenalty + gapPenalty
			my := mTable[i-1][j].Score + gapOpeningPenalty + gapPenalty
			if yy >= xy && yy >= my {
				yTable[i][j] = node{yy, "yy"}
			} else if xy >= my {
				yTable[i][j] = node{xy, "xy"}
			} else {
				yTable[i][j] = node{my, "my"}
			}

			mm := mTable[i-1][j-1].Score + pairScore
			xm := xTable[i-1][j-1].Score + pairScore
			ym := yTable[i-1][j-1].Score + pairScore
			if mm >= xm && mm >= ym {
				mTable[i][j] = node{mm, "mm"}
			} else if xm >= ym {
				mTable[i][j] = node{xm, "xm"}
			} else {
				mTable[i][j] = node{ym, "ym"}
			}
		}
	}

	// use dp table to backtrace sequence pairings
	var result []string
	i := len(seq1)
	j := len(seq2)
	var currentNode node
	if xTable[i][j].Score >= mTable[i][j].Score && xTable[i][j].Score >= yTable[i][j].Score {
		currentNode = xTable[i][j]
	} else if yTable[i][j].Score >= mTable[i][j].Score {
		currentNode = yTable[i][j]
	} else {
		currentNode = mTable[i][j]
	}
	done := false
	for !done {
		if currentNode.Path == "xx" {
			result = append(result, "_"+string(seq2[j-1]))
			j -= 1
			currentNode = xTable[i][j]
		} else if currentNode.Path == "yx" {
			result = append(result, "_"+string(seq2[j-1]))
			j -= 1
			currentNode = yTable[i][j]
		} else if currentNode.Path == "mx" {
			result = append(result, "_"+string(seq2[j-1]))
			j -= 1
			currentNode = mTable[i][j]
		} else if currentNode.Path == "yy" {
			result = append(result, string(seq1[i-1])+"_")
			i -= 1
			currentNode = yTable[i][j]
		} else if currentNode.Path == "xy" {
			result = append(result, string(seq1[i-1])+"_")
			i -= 1
			currentNode = xTable[i][j]
		} else if currentNode.Path == "my" {
			result = append(result, string(seq1[i-1])+"_")
			i -= 1
			currentNode = mTable[i][j]
		} else if currentNode.Path == "mm" {
			result = append(result, string(seq1[i-1])+string(seq2[j-1]))
			i -= 1
			j -= 1
			currentNode = mTable[i][j]
		} else if currentNode.Path == "xm" {
			result = append(result, string(seq1[i-1])+string(seq2[j-1]))
			i -= 1
			j -= 1
			currentNode = xTable[i][j]
		} else if currentNode.Path == "ym" {
			result = append(result, string(seq1[i-1])+string(seq2[j-1]))
			i -= 1
			j -= 1
			currentNode = yTable[i][j]
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
