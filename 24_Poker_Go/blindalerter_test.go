package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/Rahul-NITD/Poker"
)

func TestBlindAlerter(t *testing.T) {
	t.Run("Schedules printing Blind Values", func(t *testing.T) {
		inp := strings.NewReader("5\n")
		out := &bytes.Buffer{}
		store := NewSTUBStorage()
		alerter := SpyAlerter{}
		game := poker.NewGame(&alerter, &store)
		cli := poker.NewCLI(inp, out, game)
		cli.PlayPoker()

		cases := GenerateCases()
		AssertAlerts(t, cases, alerter)
	})
	t.Run("Test it prompts for number of users and alerts accordingly", func(t *testing.T) {
		inp := strings.NewReader("7\n")
		out := &bytes.Buffer{}
		store := NewSTUBStorage()
		alerter := SpyAlerter{}
		game := poker.NewGame(&alerter, &store)
		cli := poker.NewCLI(inp, out, game)
		cli.PlayPoker()

		got := out.String()
		want := "Enter Number of Players : "

		if got != want {
			t.Error("Did not ask for number of players")
		}

		cases := []TestAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		AssertAlerts(t, cases, alerter)

	})
}

func AssertAlerts(t *testing.T, cases []TestAlert, alerter SpyAlerter) {
	t.Helper()
	for i, c := range cases {
		t.Run(fmt.Sprintf("%d scheduled for %v", c.Amt, c.Time), func(t *testing.T) {
			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
			}
			alert := alerter.alerts[i]

			amountGot := alert.Amt
			if amountGot != c.Amt {
				t.Errorf("got amount %d, want %d", amountGot, c.Amt)
			}

			gotScheduledTime := alert.Time
			if gotScheduledTime != c.Time {
				t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, c.Time)
			}
		})
	}
}

func GenerateCases() []TestAlert {
	return []TestAlert{
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
}
