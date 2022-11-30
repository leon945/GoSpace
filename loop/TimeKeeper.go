package loop

import (
	"time"
)

type TimeKeeper struct {
	// deltaTime is in seconds
	deltaTime float64
	startTime time.Time
}

func (tk *TimeKeeper) setDeltaTime(deltaTime float64) {
	tk.deltaTime = deltaTime
}

func (tk *TimeKeeper) DeltaTime() float64 {
	return tk.deltaTime
}

func (tk *TimeKeeper) Start() {
	tk.startTime = time.Now()
}

func (tk *TimeKeeper) End() {
	currentTime := time.Now()
	deltaNano := currentTime.UnixNano() - tk.startTime.UnixNano()
	deltaSeconds := float64(deltaNano) / float64(1000000000)

	tk.setDeltaTime(deltaSeconds * 4000)
}
