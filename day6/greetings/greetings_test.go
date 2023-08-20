package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls Hello() with a name
// checking for a valid output
func TestHelloName(t *testing.T) {
	name := "AIM"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("AIM")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("AIM") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want match for "", error`, msg, err)
	}
}
