package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/Rahul-NITD/Poker"
)

func TestBlindAlerter(t *testing.T) {
	t.Run("Schedules printing Blind Values", func(t *testing.T) {
		inp := strings.NewReader("Chris wins\n")
		store := NewSTUBStorage()
		alerter := SpyAlerter{}

		cli := poker.NewCLI(&store, inp, &alerter)
		cli.PlayPoker()

		cases := []struct {
			expectedTime time.Duration
			expectedAmt  int
		}{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.expectedAmt, c.expectedTime), func(t *testing.T) {
				if len(alerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
				}
				alert := alerter.alerts[i]

				amountGot := alert.amount
				if amountGot != c.expectedAmt {
					t.Errorf("got amount %d, want %d", amountGot, c.expectedAmt)
				}

				gotScheduledTime := alert.scheduledAt
				if gotScheduledTime != c.expectedTime {
					t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, c.expectedTime)
				}
			})
		}

	})
}