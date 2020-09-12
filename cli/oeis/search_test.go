package oeis

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildUrl(t *testing.T) {
	tests := map[string]struct {
		input string
		query string
		want  string
	}{
		"simple": {
			input: "https://example.com",
			query: "1",
			want:  "https://example.com/search?fmt=json&q=1",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := buildURL(tc.input, tc.query)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestValidateJson(t *testing.T) {

	tests := map[string]struct {
		input []byte
		want  OeisQuery
	}{
		"simple": {
			input: []byte(`{"greeting": "Foo","query": "1", "count": 1, "start": 0, "results": null}`),
			want: OeisQuery{
				Greeting: "Foo",
				Query:    "1",
				Count:    1,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := validateJSON(tc.input)
			diff := cmp.Diff(tc.want, *got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
