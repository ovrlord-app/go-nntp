package nntp

import (
	"errors"
	"testing"
	"time"
)

func TestDateLayoutsGenerated(t *testing.T) {
	const expectedLayoutCount = 2 * 2 * 2 * 2 * 3
	if len(dateLayouts) != expectedLayoutCount {
		t.Fatalf("unexpected number of date layouts; got %d expected %d", len(dateLayouts), expectedLayoutCount)
	}

	seen := make(map[string]struct{}, len(dateLayouts))
	for _, layout := range dateLayouts {
		seen[layout] = struct{}{}
	}
	if len(seen) != len(dateLayouts) {
		t.Fatalf("date layouts should be unique; got %d unique out of %d", len(seen), len(dateLayouts))
	}
}

func TestParseDateSuccess(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  time.Time
	}{
		{
			name:  "rfc5322 with day of week and numeric zone",
			input: "Mon, 2 Jan 2006 15:04:05 -0700",
			want:  time.Date(2006, time.January, 2, 15, 4, 5, 0, time.FixedZone("", -7*60*60)),
		},
		{
			name:  "no day of week and no seconds",
			input: "2 Jan 2006 15:04 -0700",
			want:  time.Date(2006, time.January, 2, 15, 4, 0, 0, time.FixedZone("", -7*60*60)),
		},
		{
			name:  "two-digit year with padded day",
			input: "02 Jan 06 15:04 -0700",
			want:  time.Date(2006, time.January, 2, 15, 4, 0, 0, time.FixedZone("", -7*60*60)),
		},
		{
			name:  "timezone abbreviation",
			input: "Mon, 02 Jan 06 15:04:05 MST",
			want:  time.Date(2006, time.January, 2, 15, 4, 5, 0, time.FixedZone("MST", 0)),
		},
		{
			name:  "numeric zone with abbreviation suffix",
			input: "Mon, 02 Jan 06 15:04:05 -0700 (MST)",
			want:  time.Date(2006, time.January, 2, 15, 4, 5, 0, time.FixedZone("MST", -7*60*60)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseDate(tc.input)
			if err != nil {
				t.Fatalf("parseDate should succeed for %q: %v", tc.input, err)
			}
			if !got.Equal(tc.want) {
				t.Fatalf("parsed time mismatch for %q; got %v expected %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestParseDateFailure(t *testing.T) {
	got, err := parseDate("this is not a date")
	if !errors.Is(err, errDateCannotBeParsed) {
		t.Fatalf("expected errDateCannotBeParsed; got %v", err)
	}
	if !got.IsZero() {
		t.Fatalf("expected zero time on parse failure; got %v", got)
	}
}
