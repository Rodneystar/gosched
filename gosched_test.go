package gosched

import (
	"time"
	"testing"
)

func Test_learnTIme( t *testing.T) {
	task  := DailyTask{}

	task.TriggerAt = time.Now()

	t.Logf("timeNowIs: %s", task.TriggerAt)
	t.Logf("wall time : %s", time.Now().)
}