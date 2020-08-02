package main

import "fmt"

const englishPrefixHello = "Hello, "
const englishSuffixWorld = "World"

const french = "French"
const frenchPrefixHello = "Bonjour, "

const spanish = "Spanish"
const spanishPrefixHello = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = englishSuffixWorld
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefixHello
	case spanish:
		prefix = spanishPrefixHello
	default:
		prefix = englishPrefixHello
	}
	return
}

func main() {
	fmt.Println(Hello(englishSuffixWorld, ""))
}
