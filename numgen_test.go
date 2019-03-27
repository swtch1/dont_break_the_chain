package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratingNumbers(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name           string
		start          int
		iterations     int
		expectedResult string
	}{
		{"1_iteration", 0, 1, " 1"},
		{"2_iterations", 1, 2, " 2"},
		{"10_iterations", 0, 10, "10"},
		{"100_iterations", 0, 100, "100"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ng := NewNumGen()

			var r string
			for i := 0; i < tt.iterations; i++ {
				r = ng.Next()
			}

			assert.Equal(tt.expectedResult, r)
		})
	}
}
