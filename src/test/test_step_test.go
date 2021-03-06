package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcNumRuns(t *testing.T) {
	// Helper for assert
	nr := func(a, b int) []interface{} { return []interface{}{a, b} }

	// Base case when no flags are passed
	assert.Equal(t, nr(1, 1), nr(calcNumRuns(0, 0)))
	// Trivially flaky test; run n times, one success is enough
	assert.Equal(t, nr(3, 1), nr(calcNumRuns(0, 3)))
	// Non-flaky test with multiple runs; run n times, must succeed every time
	assert.Equal(t, nr(3, 3), nr(calcNumRuns(3, 0)))
	// This is where it gets fiddly; when we pass both flags we should run the exact number
	// of times but maintain a proportionate amount of flakiness.
	assert.Equal(t, nr(1, 1), nr(calcNumRuns(1, 1)))
	assert.Equal(t, nr(1, 1), nr(calcNumRuns(1, 3)))
	assert.Equal(t, nr(3, 1), nr(calcNumRuns(3, 3)))
	assert.Equal(t, nr(6, 2), nr(calcNumRuns(6, 3)))
	assert.Equal(t, nr(7, 3), nr(calcNumRuns(7, 3)))
}
