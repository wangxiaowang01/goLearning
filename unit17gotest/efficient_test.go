package unit17gotest

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	words := strings.Split("a:b:c", ":")
	assertEqual(len(words), 3)
	// ...
}
func TestSplit02(t *testing.T) {
	// words := strings.Split("a:b:c", ":")
	s := "a:b:c"
	sep := ":"
	words := strings.Split(s, sep)

	if got, want := len(words), 3; got != want {
		t.Errorf("Split(%q, %q) returned %d words, want %d", s, sep, got, want)
	}

}

// type t struct
func TestSplit03(t *testing.T) {
	// words := strings.Split("a:b:c", ":")
	// s := "a:b:c"
	// sep:=":"
	// words := strings.Split(s, sep)
	var testDates = []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:v", ":", 3},
		{"a:c:v", ":", 3},
		{"a:c:v", ":", 3},
	}
	for _, testDate := range testDates {
		words := strings.Split(testDate.s, testDate.sep)

		if got, want := len(words), testDate.want; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", testDate.s, testDate.sep, got, want)
		}
	}

}
func assertEqual(x, y int) {
	if x != y {
		panic(fmt.Sprintf("%d != %d", x, y))
	}
}
