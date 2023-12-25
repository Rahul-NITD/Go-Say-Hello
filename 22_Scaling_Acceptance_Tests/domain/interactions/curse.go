package interactions

import "fmt"

func Curse(name string) string {
	if name == "" {
		return "Go to Hell Nobody"
	}
	return fmt.Sprintf("Go to Hell %s", name)
}
