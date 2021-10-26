package sum

import "testing"


/*

	[numbers    | want			]
	1,2,3,4,5	| 15
	nil         | 0
	1 , -1      | 0

*/

//test functions should begin with Test word followed by Uppercase first letter of func name
func TestSumInt(t *testing.T) {
	tt := []struct {
		name    string // struct fields // table tests
		numbers []int
		want    int
	}{
		// struct values
		{
			name:    "one to five",
			numbers: []int{1, 2, 3, 4, 5}, // test cases
			want:    15,
		},
		{
			name:    "nil slice",
			numbers: nil,
			want:    0,
		},
		{
			name:    "one minus one",
			numbers: []int{1, -1},
			want:    0,
		},
		{
			name:    "one to three",
			numbers: []int{1, 2, 3},
			want:    6,
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) { //t.Run is used to create sub tests // which means each test runs as separate unit

			got := SumInt(tc.numbers)
			if got != tc.want {
				t.Errorf("sum of %v want %v; got %v", tc.numbers, tc.want, got)
			}

		})

	}

}

// Errorf -> when we want to continue exec
// Fatalf -> when we want to stop test when we have an error