package norwegish_test

import (
	"testing"

	"github.com/kytrinyx/norwegish"
)

func TestTranslate(t *testing.T) {

	var testCases = []struct {
		in  string
		out string
	}{
		{"hello", "hellø"},
		{"can", "kan"},
		{"thing", "ting"},
		{"wish", "vish"},
	}
	for _, tt := range testCases {
		got := norwegish.Translate(tt.in)
		if got != tt.out {
			t.Errorf("expected Translate(%s) to equal `%s`, got %#v", tt.in, tt.out, got)
		}
	}
}
