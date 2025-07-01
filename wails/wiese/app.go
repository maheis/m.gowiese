package main

import (
	"fmt"
)

// Color returns the name of the color
func (a *App) Color(name string) string {
	return name
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hallo %s, ab geht die Party! 🥳", name)
}
