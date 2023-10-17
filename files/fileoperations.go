package files

import (
	"log"
	"os"
)

func FileOperations() {
	FileCreateAppand()
}

func FileCreateAppand() {
	// https://www.golangprograms.com/files-directories-examples.html
	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//CoPilot:
	// defer ist ein Schlüsselwort in Go, das verwendet wird, um eine Funktion oder einen Aufruf zu verzögern, bis die umgebende Funktion zurückkehrt. Mit anderen Worten, defer wird verwendet, um eine Funktion oder einen Aufruf am Ende einer Funktion auszuführen, unabhängig davon, ob die Funktion erfolgreich abgeschlossen wurde oder einen Fehler zurückgegeben hat.
	// In dem gegebenen Code-Beispiel wird defer f.Close() aufgerufen, um die Datei zu schließen, nachdem sie geöffnet wurde. Das bedeutet, dass die Close()-Funktion automatisch aufgerufen wird, wenn die Funktion FileOperations() beendet wird, unabhängig davon, ob sie erfolgreich abgeschlossen wurde oder einen Fehler zurückgegeben hat.
	// Das Verwenden von defer ist eine nützliche Technik, um sicherzustellen, dass Ressourcen wie Dateien, Netzwerkverbindungen oder Datenbankverbindungen ordnungsgemäß freigegeben werden, wenn sie nicht mehr benötigt werden. Es kann auch dazu beitragen, den Code lesbarer zu machen, indem es die Schließung von Ressourcen an das Ende der Funktion verschiebt, anstatt sie in der Mitte des Codes zu platzieren.

	f.WriteString("Hello File\r")
	f.WriteString("I write Something!\n")
}

func FileRead() {

}
