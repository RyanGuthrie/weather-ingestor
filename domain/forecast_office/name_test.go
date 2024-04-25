package forecast_office

import (
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		name       string
		input      string
		wantResult Name
		wantError  bool
	}{
		{"Valid office ID", "TOP", TOP, false},
		{"Unknown office ID", "NonExistingOfficeIdentifier", "", true},
		{"Empty input", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Instances.FromStringLazy(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("parse(%q) error = %v, wantError %v", tt.input, err, tt.wantError)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("parse(%q) = %v, want %v", tt.input, gotResult, tt.wantResult)
			}
		})
	}
}
