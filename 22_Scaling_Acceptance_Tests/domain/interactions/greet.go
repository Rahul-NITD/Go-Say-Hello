package interactions

import "fmt"

func Greet(name string) string {
	if name == "" {
		return "Hello World"
	}
	return fmt.Sprintf("Hello %s", name)
}
