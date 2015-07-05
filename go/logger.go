package main

import (
	"fmt"
	"time"
)

// Log Format
const format = "[%s] [%s] [%s]\n"

// INFO
type Info struct {
	message  string
	severity string
}

func (l Info) String() string {
	l.severity = "info"
	return output(l.severity, l.message)
}

// Debug
type Debug struct {
	message  string
	severity string
}

func (l Debug) String() string {
	l.severity = "debug"
	return output(l.severity, l.message)
}

// Fatal
type Fatal struct {
	message  string
	severity string
}

func (l Fatal) String() string {
	l.severity = "fatal"
	return output(l.severity, l.message)
}

func output(severtiy string, message string) string {
	return fmt.Sprintf(format, time.Now(), severtiy, message)
}

func main() {
	info := Info{message: "something info about program"}
	fmt.Println(info)

	debug := Debug{message: "something debug about program"}
	fmt.Println(debug)

	fatal := Fatal{message: "something Fatal about program"}
	fmt.Println(fatal)
}
