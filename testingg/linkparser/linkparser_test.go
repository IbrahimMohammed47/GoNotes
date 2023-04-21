package linkparser_test

import (
	"os"
	"testing"
	"testingg/linkparser"
)

// one test
func TestParseHtml(t *testing.T) {
	r, err := os.Open("../examples/ex2.html")
	if err != nil {
		t.Errorf("Failed to open file: %s", err)
	}

	got, err := linkparser.ParseHtml(r)
	if err != nil {
		t.Errorf("failed to parse HTML: %s", err)
	}
	expected := []linkparser.Link{
		{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
		{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "I'm a nested a",
		},
		{
			Href: "https://github.com/gophercises",
			Text: "Gophercises is on Github !",
		},
		{
			Href: "/dog",
			Text: "Something in a span Text not in a span Bold text!",
		},
	}
	for i, l := range got {
		if l.Href != expected[i].Href {
			t.Errorf("Expected %s, got %s", expected[i].Href, l.Href)
		}
		if l.Text != expected[i].Text {
			t.Errorf("Expected %s, got %s", expected[i].Text, l.Text)
		}
	}
}

type subTest struct {
	name     string
	input    string
	expected string
}

// one test with subtests
func TestParseHtml2(t *testing.T) {
	table := []subTest{
		{name: "subtest_ex1", input: "../examples/ex1.html", expected: "/other-page"},
		{name: "subtest_ex2", input: "../examples/ex2.html", expected: "https://www.twitter.com/joncalhoun"},
	}

	for _, st := range table {
		t.Run(st.name, func(t *testing.T) {
			r, err := os.Open(st.input)
			if err != nil {
				t.Errorf("Failed to open file: %s", err)
			}
			got, err := linkparser.ParseHtml(r)
			if err != nil {
				t.Errorf("failed to parse HTML: %s", err)
			}
			if got[0].Href != st.expected {
				t.Errorf("Expected %s, got %s", st.expected, got[0].Href)
			}

		})
	}
}
