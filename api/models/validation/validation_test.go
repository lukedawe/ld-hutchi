package validation

import (
	"testing"
)

// **Gemini generated** table test cases.
func TestNameValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		// --- Success Cases (Must be strictly lowercase and no whitespace) ---
		{
			name:    "Success: Basic Valid Breed Name",
			input:   "beagle",
			wantErr: false,
		},
		{
			name:    "Success: Long Valid Breed Name",
			input:   "labrador",
			wantErr: false,
		},
		{
			name:    "Success: Short Valid Breed Name",
			input:   "pug",
			wantErr: false,
		},

		// --- Failure Cases: Empty/Whitespace ---
		{
			name:    "Fail: Empty String",
			input:   "",
			wantErr: true,
		},
		{
			name:    "Fail: Only Spaces",
			input:   "   \t\n", // Tabs, newlines, and spaces
			wantErr: true,
		},
		{
			name:    "Fail: Leading/Trailing Whitespace",
			input:   " chihuahua ",
			wantErr: true,
		},
		{
			name:    "Fail: Internal Spaces (Multi-word breed name)",
			input:   "great dane",
			wantErr: true,
		},

		// --- Failure Cases: Invalid Characters (Failing Regex for non-a-z) ---
		{
			name:    "Fail: Contains Uppercase",
			input:   "Beagle",
			wantErr: true,
		},
		{
			name:    "Fail: All Uppercase",
			input:   "POODLE",
			wantErr: true,
		},
		{
			name:    "Fail: Mixed Case",
			input:   "Dalmatian",
			wantErr: true,
		},
		{
			name:    "Fail: Contains Numbers",
			input:   "rottweiler1",
			wantErr: true,
		},
		{
			name:    "Fail: Contains Symbols (Hyphenated breed name)",
			input:   "saint-bernard",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.input)

			// 1. Check for expected error state (presence or absence of error)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validateName(%q) error status mismatch. wantErr: %v, got error: %v", tt.input, tt.wantErr, err)
			}
		})
	}
}
