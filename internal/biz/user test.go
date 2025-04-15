package biz

import "testing"

func TestHashPassword(t *testing.T) {
	// Test cases for the hashPassword function
	tests := []struct {
		input    string
		expected string
	}{
		{"password123", "$2a$10$KIXb7h5jvY1g8m0Z5c3f6OeQ9f9f9f9f9f9f9f9f9f9f9f9f9f9"},
		{"mypassword", "$2a$10$KIXb7h5jvY1g8m0Z5c3f6OeQ9f9f9f9f9f9f9f9f9f9f9f9f9"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := hashPassword(test.input)
			if result != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}
