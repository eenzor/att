package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	reset  = "\033[0m\n"
)

func TestFormatVersion(t *testing.T) {

	v := "10.1.9"
	c := "ea11f6e"

	expected :=
		`[
  {
    "version": "10.1.9",
    "lastcommitsha": "ea11f6e",
    "description": "description"
  }
]`

	received, err := formatVersion(v, c, "description")
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}
	if expected != received {
		fmt.Print(diff(expected, received))
		t.Fail()
	}
}

// a basic diff function to pretty print the difference between two strings
// matching sections are printed in green, expected in yellow, and mismatching in red
// not very robust against missing characters (prints a rainbow)
// but makes it easy to spot where the strings diverge
func diff(a string, b string) string {

	a, b = pad(a, b)
	var buf bytes.Buffer
	match := true
	buf.WriteString(green)

	for i := range a {
		if a[i] == b[i] {
			if match != true {
				buf.WriteString(green)
			}
			buf.WriteByte(a[i])
			match = true
		} else {
			buf.WriteString(yellow)
			buf.WriteByte(a[i])
			buf.WriteString(red)
			buf.WriteByte(b[i])
			match = false
		}

	}
	buf.WriteString(reset)
	return buf.String()
}

func pad(a string, b string) (string, string) {

	padChar := " "
	flipped := false

	// make sure a is always the longest string
	if len(b) > len(a) {
		a, b = b, a
		flipped = true
	}

	// enter a loop
	for {
		// add the padding character to string b
		b += padChar

		// if the length of b is now grater than a
		// return a, and b up to the length of a
		if len(b) > len(a) {
			break
		}
	}

	if flipped {
		return b[0:len(a)], a
	}

	return a, b[0:len(a)]

}

func TestPad(t *testing.T) {
	a := "1234"
	b := "123456"

	a, b = pad(a, b)
	if len(a) != len(b) {
		fmt.Println("a=|" + a + "|")
		fmt.Println("b=|" + b + "|")
		t.Fail()
	}
}

func TestDiff(t *testing.T) {
	a := "123ABC"
	b := "123XB"

	d := diff(a, b)
	// we know the strings differ
	// so we should expect to see the escape sequence for red in there
	if !strings.Contains(d, red) {
		t.Fail()
	}
}
