package sum

import "testing"

//test functions should begin with Test word followed by Uppercase first letter of func name
func TestSumInt(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}

	got := SumInt(x)
	want := 15

	if got!= want {
		t.Errorf("sum of 1 to 5 should be %v; got %v ", want, got)
	}

	x= nil
	got = SumInt(x)
	want = 0


	if got != want {
		t.Errorf("sum of nil slice should be %v; got %v", want, got)
	}

}

// Errorf -> when we want to continue exec
// Fatalf -> when we want to stop test when we have an error