package bracket_test

import (
	"bracket"
	"testing"
)

func TestCheckOnEmptyStringShouldBeTrue(t *testing.T) {

	if bracket.Check("") != true {
		t.Errorf("Error check \"\" should be true")
	}
}

func TestCheckOnOnlyStartBracketShouldBeFalse(t *testing.T) {
	input := "{"
	expected := false

	if bracket.Check(input) != expected {
		t.Errorf("Error \"%v\" should be false", input)
	}
}
func TestCheckOnOnlyEndBracketShouldBeFalse(t *testing.T) {
	input := "}"
	expected := false

	if bracket.Check(input) != expected {
		t.Errorf("Error \"%v\" should be false", input)
	}
}

func TestCheckOnOnlyAlphabetShouldBeFalse(t *testing.T) {
	input := "X"
	expected := true

	if bracket.Check(input) != expected {
		t.Errorf("Error \"%v\" should be false", input)
	}
}
