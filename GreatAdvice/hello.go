package main

import "fmt"

// We want our domain to be separate from side effects
// separate out print statement and the sentence

// refactoring

// META
const space = " "
const exlamationMark = "!"
const emptyString = ""

// Languages Supported
const hindiLang = "Hindi"
const kannadaLang = "Kannada"

// address Everyone
const addressEveryoneEnglish = "everybody"
const addressEveryoneHindi = "sablog"
const addressEveryoneKannada = "yelru"

// polite greetings
const aVeryPoliteGreetingInEnglish = "Go do your f*cking laundry"
const aVeryPoliteGreetingInHindi = "Kapde dhole saale"
const aVeryPoliteGreetingInKannada = "batte toleyo ley"

// Domain code
func Hello(name string, language string) string {
	if name == "" {
		switch language {
		case hindiLang:
			name = addressEveryoneHindi
		case kannadaLang:
			name = addressEveryoneKannada
		default:
			name = addressEveryoneEnglish
		}
	}
	switch language {
	case hindiLang:
		return aVeryPoliteGreetingInHindi + space + name + exlamationMark
	case kannadaLang:
		return aVeryPoliteGreetingInKannada + space + name + exlamationMark
	default:
		return aVeryPoliteGreetingInEnglish + space + name + exlamationMark
	}
}

// Side Effect
func main() {
	fmt.Println(Hello("Rahul", "english"))
}
