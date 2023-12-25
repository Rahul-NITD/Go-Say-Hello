package specs

import "testing"

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t *testing.T, meanGreeter MeanGreeter) {
	t.Run("Test Anu", func(t *testing.T) {
		got, err := meanGreeter.Curse("Anu")
		if err != nil {
			t.Errorf("Did not expect error, %v", err)
		}
		if got != "Go to Hell Anu" {
			t.Errorf("got %s != %s", got, "Go to Hell Anu")
		}
	})
	t.Run("Test No Name", func(t *testing.T) {
		got, err := meanGreeter.Curse("")
		if err != nil {
			t.Errorf("Did not expect error, %v", err)
		}
		if got != "Go to Hell Nobody" {
			t.Errorf("got %s != %s", got, "Go to Hell Nobody")
		}
	})
}
