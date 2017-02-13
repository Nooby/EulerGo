package euler

import "testing"

var verifyTests = []struct {
	challenge int
	solution  string
	expected  bool
	error     bool
}{

	{-1, "0", false, true},      // Range Check
	{0, "0", false, true},       // Range Check
	{9999, "0", false, true},    // Range Check
	{1, "233168", true, false},  // Success Test
	{1, "233160", false, false}, // Failure Test
	{474, "0", false, true},     // No Solution available
	{470, "0", false, true},     // No Solution available
}

func TestVerify(t *testing.T) {
	for _, tt := range verifyTests {
		solved, err := Verify(tt.challenge, tt.solution)
		if err != nil {
			if !tt.error {
				t.Errorf("Verify(%v, %v) threw error: %v",
					tt.challenge, tt.solution, err)
			}
		} else {
			if tt.error {
				t.Errorf("Verify(%v, %v) expected error",
					tt.challenge, tt.solution)
			}
			if solved != tt.expected {
				t.Errorf("Verify(%v, %v) expected %v got %v",
					tt.challenge, tt.solution, tt.expected, solved)
			}
		}

	}

}
