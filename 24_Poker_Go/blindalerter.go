package poker

import "time"

type BlindAlerter interface {
	ScheduleAlertAfter(duration time.Duration, amount int)
}
