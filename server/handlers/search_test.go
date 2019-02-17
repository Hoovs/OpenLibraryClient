package handlers

import (
	"go.uber.org/zap"
	"net/url"
	"testing"
)

func TestCreateUrl(t *testing.T) {
	baseUrl := "http://test?q="
	cases := []struct {
		name        string
		v           url.Values
		validate    func(error) bool
		expectedStr string
	}{
		{
			name: "nil values fails",
			v:    nil,
			validate: func(e error) bool {
				return e != nil
			},
			expectedStr: "",
		}, {
			name: "single word",
			v:    url.Values{"q": []string{"test"}},
			validate: func(e error) bool {
				return e == nil
			},
			expectedStr: "http://test%3Fq=test",
		}, {
			name: "two words",
			v:    url.Values{"q": []string{"test query"}},
			validate: func(e error) bool {
				return e == nil
			},
			expectedStr: "http://test%3Fq=test%20query",
		},
	}

	l, _ := zap.NewDevelopment()
	h := SearchHandler{
		Logger:        l,
		BaseSearchUrl: baseUrl,
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			u, err := h.createLibraryURL(c.v)
			if !c.validate(err) {
				t.Errorf("Error didn't match expected")
			}
			if u != c.expectedStr {
				t.Errorf("%s didn't match expected: %s", u, c.expectedStr)
			}
		})
	}
}
