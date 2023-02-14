package funfacts

import (
	"reflect"
	"testing"
)

// Definer korrekte typer for input og want.

func TestGetFunFacts(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{input: "sun", want: funFacts.Sun},
		{input: "luna", want: funFacts.Luna},
		{input: "terra", want: funFacts.Terra},
		{input: "unknown", want: []string{}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
