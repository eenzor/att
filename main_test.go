package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testVersion     = "10.1.9"
	testSHA         = "ea11f6e"
	testDescription = "description"
	testMetadata    = `[
  {
    "version": "10.1.9",
    "lastcommitsha": "ea11f6e",
    "description": "description"
  }
]`
)

const (
	// terminal color escape sequences
	red   = "\033[31m"
	green = "\033[32m"
	reset = "\033[0m"
)

func TestFormatVersion(t *testing.T) {

	received, err := formatVersion(testVersion, testSHA, "description")
	if err != nil {
		t.Error(err.Error())
	}
	if received != testMetadata {
		got, want, _ := diff(received, testMetadata)
		t.Errorf("formatVersion returned unexpected response:\ngot\n%v\nwant\n%v", got, want)
	}
}

func TestVersionHandler(t *testing.T) {

	version = testVersion
	commit = testSHA
	description = testDescription
	logFormat = "none"

	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(versionHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	expected := fmt.Sprintf("\"myapplication\": %s\n", testMetadata)

	if rr.Body.String() != expected {
		got, want, _ := diff(rr.Body.String(), expected)
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

// diff takes two strings and compares them, returning them with the differences
// colour coded, in addition to a bool indicating if they differ
func diff(a string, b string) (string, string, bool) {

	if a == b {
		return a, b, false
	}

	a, b = pad(a, b)
	var aBuf bytes.Buffer
	var bBuf bytes.Buffer
	match := true

	for i := range a {
		if a[i] == b[i] {
			if match != true {
				aBuf.WriteString(reset)
				bBuf.WriteString(reset)
			}
			aBuf.WriteByte(a[i])
			bBuf.WriteByte(b[i])
			match = true
		} else {
			aBuf.WriteString(red)
			aBuf.WriteByte(a[i])
			bBuf.WriteString(green)
			bBuf.WriteByte(b[i])
			match = false
		}

	}

	aBuf.WriteString(reset)
	bBuf.WriteString(reset)

	return aBuf.String(), bBuf.String(), true
}

// pad takes two strings and right pads the shorter
// to be the same length as the longer
func pad(a string, b string) (string, string) {

	if len(a) == len(b) {
		return a, b
	}

	padChar := "~"
	flipped := false

	// make sure a is always the longest string
	if len(b) > len(a) {
		a, b = b, a
		flipped = true
	}

	// keep adding the padding character until the strings are equal in length
	for {
		b += padChar
		if len(b) >= len(a) {
			break
		}
	}

	// return the strings in the correct order
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
		t.Errorf("\na=|" + a + "|\nb=|" + b + "|")
	}
}

func TestDiff(t *testing.T) {
	a := "123ABC"
	b := "123XB"

	_, _, e := diff(a, b)
	if !e {
		t.Error("diff says unequal strings are equal")
	}
}
