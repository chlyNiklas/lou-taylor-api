package middleware

import "testing"

func TestMatchPath(t *testing.T) {
	tests := []struct {
		a     string
		b     string
		match bool
	}{
		{
			a:     "/asdf/{joggölid}/asdf",
			b:     "/asdf/ruedi",
			match: false,
		},
		{
			a:     "/asdf/{joggölid}",
			b:     "/asdf/ruedi",
			match: true,
		},
		{
			a:     "/asdf/{joggölid}",
			b:     "/asdf/ruedi/asdf",
			match: false,
		},
		{
			a:     "/asdf/{joggölid}/asdf/{someId}",
			b:     "/asdf/ruedi/asdf/3",
			match: true,
		},
	}

	for _, test := range tests {
		if matchPath(test.a, test.b) != test.match {
			t.Errorf("Expected %s & %s to %v", test.a, test.b, test.match)
		}
	}

}
