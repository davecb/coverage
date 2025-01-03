package testIterator

import "math"

// FloatSample is returned by float64 value generators.
type FloatSample struct {
	Value float64
	Legal bool
}

// IntSample is returned by int value generators
type IntSample struct {
	Value int
	Legal bool
}

// PositiveFloat is a sample set for positive.
// Notice we inject some decimal points, not just integer-equivalents
var PositiveFloat = []FloatSample{
	{-math.MaxFloat64 + 0, false},
	{-math.MaxFloat64 + 1.1, false},
	{-math.MaxFloat64/2 + .2, false},
	{-(math.MaxFloat64 / 2) + 1.3, false},
	{-3.4, false},
	{-2.5, false},
	{-1.6, false},
	{0, false},
	{1.7, true},
	{2.8, true},
	{3.9, true},
	{4, true},
	{(math.MaxFloat64 / 2) - 1.1, true},
	{math.MaxFloat64/2 + .2, true},
	{math.MaxFloat64 - 1.3, true},
	{math.MaxFloat64, true},
}

// SmallPositiveCounter is a sample-set for ... small positive counters!
var SmallPositiveCounter = []IntSample{
	{-10, false},
	{-5, false},
	{-3, false},
	{-2, false},
	{-1, false},
	{0, false},
	{1, true},
	{2, true},
	{3, true},
	{5, true},
	{10, true},
}

// AllTrue looks for any false values in a vector of booleans
// typically used like testIterator.AllTrue(this.Legal, that.Legal, etc)
func AllTrue(b ...bool) bool {
	for _, v := range b {
		if !v {
			return false
		}
	}
	return true
}
