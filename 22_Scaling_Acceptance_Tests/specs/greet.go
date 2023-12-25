package specs

import "testing"

type Greeter interface {
	Greet(name string) (string, error)
}

func GreeterSpecification(t *testing.T, greeter Greeter) {
	t.Run("Test Mike", func(t *testing.T) {
		got, err := greeter.Greet("Mike")
		if err != nil {
			t.Errorf("Did not expect error, %v", err)
		}
		if got != "Hello Mike" {
			t.Errorf("got %s != %s", got, "Hello Mike")
		}
	})
	t.Run("Test No Name", func(t *testing.T) {
		got, err := greeter.Greet("")
		if err != nil {
			t.Errorf("Did not expect error, %v", err)
		}
		if got != "Hello World" {
			t.Errorf("got %s != %s", got, "Hello World")
		}
	})
}
