package norwegish

import "testing"

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
		got := Translate(tt.in)
		if got != tt.out {
			t.Errorf("expected Translate(%s) to equal `%s`, got %#v", tt.in, tt.out, got)
		}
	}
}
