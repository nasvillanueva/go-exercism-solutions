package clock

import (
	"fmt"
	"math"
)

const testVersion = 4

type Clock struct {
	Hour   int
	Minute int
}

func New(hour, minute int) Clock {
	c := Clock{getHour(hour), 0}
	return c.Add(minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	minutes = c.Minute + minutes
	for math.Abs(float64(minutes)) >= 60 || minutes < 0 {
		if minutes >= 60 {
			minutes -= 60
			c.Hour++
		} else {
			minutes += 60
			c.Hour--
		}
	}
	c.Minute = minutes
	c.Hour = getHour(c.Hour)
	return c
}

func getHour(hour int) int {

	for math.Abs(float64(hour)) >= 24 || hour < 0 {
		if hour >= 24 {
			hour -= 24
		} else {
			hour += 24
		}
	}
	if hour < 0 {
		hour += 24
	}
	return hour
}
