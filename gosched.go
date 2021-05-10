package gosched

import (
	"time"	
)

type Clock struct {
	Hour int
	Minute int
	Second int
}

func NewClockFromTime( t time.Time) Clock {
	h, m, s := t.Clock();

	return Clock{ 
		Hour: h,
		Minute: m,
		Second: s,
	}
}

type DailyTask struct {
	TriggerAt time.Time
}