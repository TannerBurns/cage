package models

import (
	"math"
	"time"
)

type Timer struct {
	StartTime    time.Time
	StopTime     time.Time
	TotalElapsed float64
}

func NewTimer() *Timer {
	return &Timer{}
}

func (timer *Timer) Start() {
	timer.StartTime = time.Now()
}

func (timer *Timer) Stop() {
	timer.StopTime = time.Now()
	elapsed := timer.StopTime.Sub(timer.StartTime)
	timer.TotalElapsed = math.Round(elapsed.Seconds())
}

func (timer *Timer) Timer(channel chan interface{}) {
	timer.Start()
	for {
		_, ok := <-channel
		if !ok {
			timer.Stop()
			return
		}
	}
}

func (timer *Timer) Elapsed() (elapsed float64) {
	e := time.Now().Sub(timer.StartTime)
	elapsed = math.Round(e.Seconds())
	return
}
