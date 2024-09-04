package basics

import "testing"

func TestDivideTwo(t *testing.T) {
	want := 2.5
	output, err := DivideTwo(5, 2)

	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	if want != output {
		t.Errorf("Output %v is not equal to want %v", output, want)
	}

}
