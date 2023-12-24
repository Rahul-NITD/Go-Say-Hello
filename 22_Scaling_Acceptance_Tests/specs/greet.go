package specs

import "testing"

type Greeter interface {
	Greet(name string) (string, error)
}

func GreeterSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Mike")
	if err != nil {
		t.Errorf("Did not expect error, %v", err)
	}
	if got != "Hello Mike" {
		t.Errorf("got %s != %s", got, "Hello Mike")
	}
}
