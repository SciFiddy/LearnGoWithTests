package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const english_hello_and_a_comma = "Hello, "
const spanish_hola_and_a_comma = "Hola, "
const french_and_a_comma = "I Surrender, "

func Hello(name string, language string) string {

	greeting := english_hello_and_a_comma

	if name == "" {
		name = "World"
	}

	switch language {
	case french:
		greeting = french_and_a_comma
	case spanish:
		greeting = spanish_hola_and_a_comma
	}
	return greeting + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
