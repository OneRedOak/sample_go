package main

import "testing"

func TestMain(t *testing.T) {
	answer := 2.0

	main()

	if getResult() != answer {
		t.Errorf("Main did not return the square root of 4 as %v", answer)
	}
}