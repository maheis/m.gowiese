package jsonfile

import (
	"encoding/json"
	"fmt"
	"os"
)

// Struct für die JSON-Datei
// Cooler Generator: https://mholt.github.io/json-to-go/
type Person struct {
	VTer []struct {
		Name        string `json:"name"`
		Alter       int    `json:"alter"`
		Verheiratet bool   `json:"verheiratet"`
		Beruf       string `json:"beruf"`
		Kinder      []struct {
			Name           string `json:"name"`
			Alter          int    `json:"alter"`
			Schulabschluss string `json:"schulabschluss"`
		} `json:"kinder"`
	} `json:"Väter"`
}

// Implementiert das Stringer interface und Formatiert die Ausgabe von fmt.Print()
// Diese Funktion darf nur auf dem Struct direkt genutzt werden, nicht auf einem Pointer!
func (p Person) String() string {
	ret := ""
	for _, v := range p.VTer {
		ret += fmt.Sprintln("  Name: ", v.Name)
		ret += fmt.Sprintln("  Alter: ", v.Alter)
		ret += fmt.Sprintln("  Verheiratet: ", v.Verheiratet)
		ret += fmt.Sprintln("  Beruf: ", v.Beruf)
		for _, k := range v.Kinder {
			ret += fmt.Sprintln("    Kind: ", k.Name)
			ret += fmt.Sprintln("    Alter: ", k.Alter)
			ret += fmt.Sprintln("    Schulabschluss: ", k.Schulabschluss)
		}
	}
	return ret
}

func Jsonfile() {
	// Datei öffnen
	file, err := os.Open("jsonfile/data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// JSON-Datei einlesen und in Struct schreiben
	var person Person
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}

	// Struct bearbeiten und einen Wert manipulieren
	if person.VTer[0].Name == "Mani" {
		person.VTer[0].Name = "Manfred"
	} else {
		person.VTer[0].Name = "Mani"
	}

	// Struct ausgeben mit Stringer interface und meiner Formatierung von oben!
	fmt.Println(person)

	// Neue leere Datei erstellen
	file, err = os.Create("jsonfile/data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Struct in Datei schreiben
	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
}
