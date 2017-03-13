package ema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEMA(t *testing.T) {
	e := New(0, 0.9)
	assert.EqualValues(t, 0, e.Get())
	old, new := e.Update(10)
	assert.EqualValues(t, 0, old)
	assert.EqualValues(t, 9, new)
	assert.EqualValues(t, 9, e.Get())
	_, newD := e.UpdateDuration(20)
	assert.EqualValues(t, 18.9, newD)
	assert.EqualValues(t, 18.9, e.GetDuration().Nanoseconds())
}
