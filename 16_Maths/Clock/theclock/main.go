package main

import (
	theclock "GoSayHello/16_Maths/Clock"
	"os"
	"time"
)

func main() {
	tm := time.Now()
	theclock.BuildClock(os.Stdout, tm)
}
