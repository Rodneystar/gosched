package gosched

import (
	"fmt"
	"time"
)

type Clock struct {
	Seconds int
}

func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", c.Seconds/60/60, (c.Seconds/60)%60, c.Seconds%60)
}

func NewClockFromTime(t time.Time) Clock {
	h, m, s := t.Clock()

	return Clock{
		Seconds: s + m*60 + h*60*60,
	}
}

type DailyTask struct {
	TriggerAt Clock
}
