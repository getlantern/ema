// Package ema provides an exponential moving average. It can hold floating
// point values up to 3 decimal points in precision and provides a convenience
// interface for keeping EMAs of time.Durations.
package ema

import (
	"sync/atomic"
	"time"
)

const (
	// floating point values are stored to this scale (3 digits behind decimal
	// point).
	scale = 1000
)

// ema holds the Exponential Moving Average of a float64 with a the given
// default α value and a fixed scale of 3 digits. Safe to access concurrently.
// https://en.wikipedia.org/wiki/Moving_average#Exponential_moving_average.
type EMA struct {
	defaultAlpha float64
	v            int64
}

// NewEMA creates an EMA with initial value and alpha
func NewEMA(initial float64, defaultAlpha float64) *EMA {
	return &EMA{defaultAlpha: defaultAlpha, v: scaleToInt(initial)}
}

// Like NewEMA but using time.Duration
func NewEMADuration(initial time.Duration, alpha float64) *EMA {
	return NewEMA(float64(initial), alpha)
}

// UpdateAlpha calculates and stores new EMA based on the duration and α
// value passed in.
func (e *EMA) UpdateAlpha(v float64, α float64) float64 {
	oldEMA := scaleFromInt(atomic.LoadInt64(&e.v))
	newEMA := (1-α)*oldEMA + α*v
	atomic.StoreInt64(&e.v, scaleToInt(newEMA))
	return newEMA
}

// like UpdateAlpha but using the default alpha
func (e *EMA) Update(v float64) float64 {
	return e.UpdateAlpha(v, e.defaultAlpha)
}

// Like Update but using time.Duration
func (e *EMA) UpdateDuration(v time.Duration) time.Duration {
	return time.Duration(e.Update(float64(v)))
}

// Set sets the EMA directly.
func (e *EMA) Set(v float64) {
	atomic.StoreInt64(&e.v, scaleToInt(v))
}

// Like Set but using time.Duration
func (e *EMA) SetDuration(v time.Duration) {
	e.Set(float64(v))
}

// Get gets the EMA
func (e *EMA) Get() float64 {
	return scaleFromInt(atomic.LoadInt64(&e.v))
}

// Like Get but using time.Duration
func (e *EMA) GetDuration() time.Duration {
	return time.Duration(e.Get())
}

func scaleToInt(f float64) int64 {
	return int64(f * scale)
}

func scaleFromInt(i int64) float64 {
	return float64(i) / scale
}
