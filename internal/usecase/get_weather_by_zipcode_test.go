package usecase

import (
	"testing"
)

func TestIsValidZipCode(t *testing.T) {
	tests := []struct {
		zipcode  string
		expected bool
	}{
		{"12345678", true},
		{"1234567", false},
		{"123456789", false},
		{"abcdefgh", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.zipcode, func(t *testing.T) {
			result := isValidZipCode(test.zipcode)
			if result != test.expected {
				t.Errorf("isValidZipCode(%s) = %v, expected %v", test.zipcode, result, test.expected)
			}
		})
	}
}

func TestGetWeatherByZipcodeUseCase_Execute_InvalidZipcode(t *testing.T) {
	uc := NewGetWeatherByZipcodeUseCase("dummy-api-key")
	_, err := uc.Execute("invalid")
	if err == nil {
		t.Error("expected error for invalid zipcode, got nil")
	}
	if err.Error() != "invalid zipcode" {
		t.Errorf("expected error message 'invalid zipcode', got '%s'", err.Error())
	}
}
