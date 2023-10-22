package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("politely greet people [English]", func(t *testing.T) {
		got := Hello("Rahul", "english")
		want := "Go do your f*cking laundry Rahul!"
		assertHelper(t, got, want)
	})
	t.Run("politely tell everybody to do their laundry [English]", func(t *testing.T) {
		got := Hello("", "english")
		want := "Go do your f*cking laundry everybody!"
		assertHelper(t, got, want)
	})
	t.Run("politely tell Rahul to do his laundry [Hindi]", func(t *testing.T) {
		got := Hello("Rahul", "Hindi")
		want := "Kapde dhole saale Rahul!"
		assertHelper(t, got, want)
	})
	t.Run("politely tell everyone to do their laundry [Hindi]", func(t *testing.T) {
		got := Hello("", "Hindi")
		want := "Kapde dhole saale sablog!"
		assertHelper(t, got, want)
	})
	t.Run("politely tell Rahul to do his laundry [Kannada]", func(t *testing.T) {
		got := Hello("Rahul", "Kannada")
		want := "batte toleyo ley Rahul!"
		assertHelper(t, got, want)
	})
	t.Run("politely tell everybody to do their laundry [Kannada]", func(t *testing.T) {
		got := Hello("", "Kannada")
		want := "batte toleyo ley yelru!"
		assertHelper(t, got, want)
	})
}

func assertHelper(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
