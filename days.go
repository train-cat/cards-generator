package main

const (
	Sunday = 1 << iota
	Saturday
	Friday
	Thursday
	Wednesday
	Tuesday
	Monday

	daysPerWeek = 7
)

type (
	// Days is binary representation
	Days uint8

	Day struct {
		Flag   Days
		Letter string
		PosX   int
	}
)

var days []Day

// HasFlag return true if mask has specific flag set
func (d Days) HasFlag(flag Days) bool { return d&flag != 0 }

// AddFlag to the binary representation
func (d *Days) AddFlag(flag Days) { *d |= flag }

// ClearFlag remove flag
func (d *Days) ClearFlag(flag Days) { *d &= ^flag }

// ToggleFlag inverse state of the flag
func (d *Days) ToggleFlag(flag Days) { *d ^= flag }
