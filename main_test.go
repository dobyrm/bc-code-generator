package main

import (
	"testing"
	"fmt"
)

func TestRand(t *testing.T) {
	tests := []struct {
		length    int
		expectErr bool
	}{
		{16, false},
		{32, false},
		{64, false},
		{0, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Rand with length %d", tt.length), func(t *testing.T) {
			randomString, err := Rand(tt.length)
			
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			
			if err == nil && len(randomString) != tt.length {
				t.Errorf("expected string length: %d, got: %d", tt.length, len(randomString))
			}
		})
	}
}
