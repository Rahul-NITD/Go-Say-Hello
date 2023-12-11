package main

import (
	myclock "GoSayHello/ItsTime/MyClock"
	"os"
	"time"
)

func main() {
	t := time.Now()
	myclock.BuildSVG(os.Stdout, t)

}
