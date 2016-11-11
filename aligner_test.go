package aligner

import "testing"

func TestTester(t *testing.T) {
	actual := Tester("abc")
	expected := "t push --set-upstream mcm masterabc"

	if actual != expected {
		t.Error("Tester function failed expected output")
		t.Logf("expected %s but got %s", expected, actual)
	}
}
