package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type German struct {
}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(greet string) string {
	return fmt.Sprintf("Hallo %s!", greet)
}

type Italian struct {
}

func (g Italian) LanguageName() string {
	return "Italian"
}

func (g Italian) Greet(greet string) string {
	return fmt.Sprintf("Ciao %s!", greet)
}

type Portuguese struct {
}

func (g Portuguese) LanguageName() string {
	return "Portuguese"
}

func (g Portuguese) Greet(greet string) string {
	return fmt.Sprintf("Olá %s!", greet)
}

func SayHello(greet string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(greet))
}
