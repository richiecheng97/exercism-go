package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	for {
		if m >= 0 && m <= 59 {
			break
		}
		if m < 0 {
			m += 60
			h--
		}
		if m > 59 {
			m -= 60
			h++
		}
	}
	for {
		if h >= 0 && h <= 23 {
			break
		}
		if h > 23 {
			h -= 24
		}
		if h < 0 {
			h += 24
		}
	}
	return Clock{
		h: h,
		m: m,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
