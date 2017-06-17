package statsd

import (
	"testing"
	"sweetcook-backend/utils/statsd"
	"time"
)

func Test_Timing(t *testing.T)  {
	t1 := time.Now()
	m := 0
	for i:=0;i<1000;i++ {
		m += 1
	}
	statsd.Timing("shiba.timing", t1, time.Now())
	t.Log("finished.")
}