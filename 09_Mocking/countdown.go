package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countStart = 3

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := 0; i < countStart; i++ {
		fmt.Fprintln(writer, (countStart - i))
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	ds := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, ds)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (ss *SpySleeper) Sleep() {
	ss.Calls++
}

type SpyCountdownOperations struct {
	calls []string
}

func (sco *SpyCountdownOperations) Sleep() {
	sco.calls = append(sco.calls, "sleep")
}

func (sco *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	sco.calls = append(sco.calls, "write")
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
