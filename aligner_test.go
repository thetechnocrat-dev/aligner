package aligner

import (
	"testing"
)

func TestTester(t *testing.T) {
	actual := Tester("abc")
	expected := "abc"

	if actual != expected {
		t.Error("Tester function failed expected output")
		t.Logf("expected %s but got %s", expected, actual)
	}
}

// Utility functions and structs
type AlignTest struct {
	seq1, seq2, message string
	expected            []string
}

func compareArrays(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, element1 := range arr1 {
		element2 := arr2[i]
		if element1 != element2 {
			return false
		}
	}

	return true
}

// Tests
func TestAlign(t *testing.T) {
	const matchScore = float64(4)
	const mismatchPenalty = float64(-6)
	const gapPenalty = float64(-4)
	const gapOpeningPenalty = float64(-8)

	alignTests := []AlignTest{
		AlignTest{
			"gcat",
			"gcat",
			"Align does not work with identical sequences",
			[]string{"gg", "cc", "aa", "tt"},
		},
		AlignTest{
			"gt",
			"ca",
			"Align does not work when there are no matches",
			[]string{"gc", "ta"},
		},
		AlignTest{
			"actag",
			"ctaga",
			"Align does not work with gaps",
			[]string{"a_", "cc", "tt", "aa", "gg", "_a"},
		},
	}

	for _, test := range alignTests {
		actual := Align(test.seq1, test.seq2, matchScore, mismatchPenalty, gapPenalty, gapOpeningPenalty)
		if !compareArrays(actual, test.expected) {
			t.Error(test.message)
			t.Logf("expected %v, actual %v", test.expected, actual)
		}
	}
}
