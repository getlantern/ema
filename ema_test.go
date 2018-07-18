package ema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEMA(t *testing.T) {
	defaultValue := float64(1.1235235)
	e := New(defaultValue, 0.9)
	assert.EqualValues(t, defaultValue, e.Get(), "Should return default value before it's been set")
	assert.EqualValues(t, 10, e.Update(10), "First update should simply set value")
	assert.EqualValues(t, 10, e.Get())
	assert.EqualValues(t, 19, e.Update(20), "Second update should factor into moving average")
	assert.EqualValues(t, 19, e.Get())
	assert.EqualValues(t, 28.9, e.UpdateDuration(30))
	assert.EqualValues(t, 28.9, e.GetDuration().Nanoseconds())
}

func TestBinaryValue(t *testing.T) {
	e1 := New(0, 0.1)
	e2 := New(0, 0.1)
	series1 := []float64{0, 1, 1, 0, 1, 1}
	expected1 := []float64{0, 0.1, 0.19, 0.171, 0.2539, 0.3277}
	series2 := []float64{1, 1, 1, 0, 1, 1}
	expected2 := []float64{1, 1, 1, 0.9, 0.91, 0.919}
	for i := range series1 {
		msg := fmt.Sprintf("Should have expected value in round %d", i)
		assert.InDelta(t, expected1[i], e1.Update(series1[i]), 1e-9, msg)
		assert.InDelta(t, expected1[i], e1.Get(), 1e-3, msg)
		assert.InDelta(t, expected2[i], e2.Update(series2[i]), 1e-9, msg)
		assert.InDelta(t, expected2[i], e2.Get(), 1e-3, msg)
	}
}
