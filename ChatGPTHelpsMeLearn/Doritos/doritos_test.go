package doritos_test

import (
	"testing"

	doritos "github.com/ChatGPTTeachesMeLearn/Doritos"
)

func TestDoritos(t *testing.T) {
	t.Run("Test doritos for 5 levels", func(t *testing.T) {
		got := doritos.GetMeADorito(5)
		want := `    *
   * *
  * * *
 * * * *
* * * * *`
		if got != want {
			t.Errorf("got \n%s != \n%s", got, want)
		}
	})
}
