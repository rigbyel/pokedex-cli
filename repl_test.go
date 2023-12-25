package main

import "testing"
  
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		output []string
	} {
		{
			"hello world!",
			[]string {"hello", "world!"},
		},
		{
			"My FriENd",
			[]string {"my", "friend"},
		},
	}

	for _, value := range cases {
		expected := value.output
		actual := cleanInput(value.input)
		if len(actual) != len(expected) {
			t.Error("lengths are not equal")
			continue
		}
		for j := range actual {
			if actual[j] != expected[j] {
			t.Error("words are not equal")
			continue
			}
		}
	}
}