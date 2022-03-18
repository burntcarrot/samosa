package internal

import "testing"

func TestTrimString(t *testing.T) {
	tests := []struct {
		name  string
		str   string
		limit int
		want  string
	}{
		{"must not trim strings under the limit", "test", 20, "test"},
		{"must trim strings exceeding the limit", "github.com/burntcarrot/samosa", 20, "...m/burntcarrot/samosa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trimString(tt.str, tt.limit)
			if got != tt.want {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}
