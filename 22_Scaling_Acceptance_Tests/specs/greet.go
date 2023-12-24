package specs

import "testing"

type Greeter interface {
	Greet() (string, error)
}

func GreeterSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet()
	if err != nil {
		t.Errorf("Did not expect error, %v", err)
	}
	if got != "Hello World" {
		t.Errorf("got %s != %s", got, "Hello World")
	}
}
