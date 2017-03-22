// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

// New returns a Clock struct with the initialized time.

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	// If 'minute' is over an hour, increment 'hour' respectively.
	if minute >= 60 {
		hour += minute / 60
		minute = minute % 60
	}

	// If 'minute' is under an hour, add the negative hour(s) to 'hour', and add the negative minutes and 60 to 'minute'
	if minute < 0 {
		hour += minute/60 + (-1)
		minute = minute%60 + 60
	}

	// If 'hour' is over a day, update 'hour', but drop whatever would have been added to days.
	if hour >= 24 {
		hour = hour % 24
	}

	// If 'hour' is negative, add the negative number to the positive number, less the amount over -24. This was a tricky one https://github.com/golang/go/issues/448 https://golang.org/ref/spec#Arithmetic_operators
	if hour < 0 {
		hour = hour%24 + 24
	}

	return Clock{
		hour,
		minute}
}

// String implements Stringer interface.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add returns a new Clock struct with updated minutes. Subtraction is done by adding negative integers.
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes

	// If 'minute' is over an hour, increment 'hour' respectively.
	if c.minute >= 60 {
		c.hour += c.minute / 60
		c.minute = c.minute % 60
	}

	// If 'minute' is under an hour, add the negative hour(s) to 'hour', and add the negative minutes and 60 to 'minute'
	if c.minute < 0 {
		c.hour += c.minute/60 + (-1)
		c.minute = c.minute%60 + 60
	}

	// If 'hour' is negative, add the negative number to the positive number, less the amount over -24. This was a tricky one https://github.com/golang/go/issues/448 https://golang.org/ref/spec#Arithmetic_operators
	if c.hour < 0 {
		c.hour = c.hour%24 + 24
	}

	// If 'hour' is over a day, update 'hour', but drop whatever would have been added to days.
	if c.hour >= 24 {
		c.hour = c.hour % 24
	}

	return c
}
