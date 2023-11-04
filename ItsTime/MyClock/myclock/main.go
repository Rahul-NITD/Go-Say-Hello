package main

import (
	myclock "GoSayHello/ItsTime/MyClock"
	"os"
	"time"
)

func main() {
	t := time.Date(1337, time.January, 1, 0, 0, 45, 0, time.UTC)
	myclock.BuildSVG(os.Stdout, t)

}
