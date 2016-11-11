package aligner

import "testing"

func TestTester(t *testing.T) {
	actual := Tester("abc")
	expected := "abc"

	if actual != expected {
		t.Error("Tester function failed expected output")
		t.Logf("expected %s but got %s", expected, actual)
	}
}
