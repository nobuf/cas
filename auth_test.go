package cas

import "testing"

func TestHasAuthToken(t *testing.T) {
	tt := []struct {
		in       Client
		expected bool
	}{
		{Client{}, false},
		{Client{clientID: "abc"}, false},
		{Client{clientSecret: "abc"}, false},
		{Client{clientID: "abc", clientSecret: "xyz"}, true},
	}
	for _, c := range tt {
		actual := c.in.hasAuthToken()
		if actual != c.expected {
			t.Errorf("expected %v but got %v with %+v", c.expected, actual, c.in)
		}
	}
}
