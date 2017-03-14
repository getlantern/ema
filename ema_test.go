package ema

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
