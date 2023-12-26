package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAfter(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amt int)

func (b BlindAlerterFunc) ScheduleAlertAfter(duration time.Duration, amt int) {
	b(duration, amt)
}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
