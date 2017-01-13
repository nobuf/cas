package cas

import "testing"

var equalTests = []struct {
	in1 MovieID
	in2 MovieID
	out bool
}{
	{"1234", "1234", true},
	{"1234", "1233", false},
}

func TestMovieID_Equal(t *testing.T) {
	for _, tt := range equalTests {
		if tt.in1.Equal(tt.in2) != tt.out {
			t.Errorf("Unexpected results at: %v", tt)
		}
	}
}

var greaterThanTests = []struct {
	in1 MovieID
	in2 MovieID
	out bool
}{
	{"1234", "1234", false},
	{"1234", "1233", true},
	{"1234", "1235", false},
}

func TestMovieID_GreaterThan(t *testing.T) {
	for _, tt := range greaterThanTests {
		if tt.in1.GreaterThan(tt.in2) != tt.out {
			t.Errorf("Unexpected results at: %v", tt)
		}
	}
}
