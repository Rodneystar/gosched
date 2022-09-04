package gosched

import (
	"testing"
	"time"
)

func Test_learnTIme(t *testing.T) {
	task := DailyTask{}

	task.TriggerAt = NewClockFromTime(time.Now())

	t.Logf("timeNowIs: %s\n", task.TriggerAt.String())
	t.Logf("wall time : %s\n", time.Now())
}

func Test_repeatTimer(t *testing.T) {
	time.Sleep()
}

func Test_clockString(t *testing.T) {
	clock := NewClockFromTime(time.Date(2020, time.September, 2, 2, 2, 2, 2, time.UTC))

	if clock.String() != "02:02:02" {
		t.Errorf("time was not formatted correctly\n")
	}

	clock = NewClockFromTime(time.Date(2020, time.September, 2, 23, 59, 59, 99999, time.UTC))

	if clock.String() != "23:59:59" {
		t.Errorf("time was not formatted correctly, should be %s\n", "23:59:59")
	}

	clock = NewClockFromTime(time.Date(2020, time.September, 2, 0, 0, 0, 00001, time.UTC))

	if clock.String() != "00:00:00" {
		t.Errorf("time was not formatted correctly, shoould be %s\n", "00:00:00")
	}
}
