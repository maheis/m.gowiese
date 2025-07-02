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

// ChoosColor returns a greeting with the chosen color
func (a *App) ChoosColor(color string) string {
	// Name als Globale Variable definieren und hier wieder verwertden. Wie das geht muss ich noch rausfinden...
	// return fmt.Sprintf("Hallo %s, ab geht die Party! 🥳 in der Farbe %s", name, color)
	return fmt.Sprintf("Du hast die Farbe %s gewählt... 🎉", color)
}
