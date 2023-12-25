package specs

import "testing"

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t testing.TB, meanGreeter MeanGreeter) {
	got, err := meanGreeter.Curse("Anu")
	if err != nil {
		t.Errorf("Did not expect error, %v", err)
	}
	if got != "Go to Hell Anu" {
		t.Errorf("got %s != %s", got, "Go to Hell Anu")
	}
}
