package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	hour = (hour + minute/60) % 24
	minute = minute % 60
	if minute < 0 {
		hour -= 1
		minute += 60
	}
	if hour < 0 {
		hour += 24
	}
	c := Clock{hour, minute}
	return c
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour%24, c.minute%60)
}
