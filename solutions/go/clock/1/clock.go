package clock

import "fmt"

type Clock struct {
	hora, minutos int
}

func New(h, m int) Clock {
	totalMinutos := (h*60 + m) % 1440
	if totalMinutos < 0 {
		totalMinutos += 1440
	}
	return Clock{
		hora:    totalMinutos / 60,
		minutos: totalMinutos % 60,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hora, c.minutos+m)
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hora, c.minutos)
}