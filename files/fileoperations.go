package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func FileOperations() {
	FileCreateAppand()
	FileReadLine()
	FileRead()
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

	f.WriteString("Hello File\r\n")
	f.WriteString("I write Something!\r\n")
	f.WriteString(time.DateTime)
	f.WriteString("\r\n")
}

func FileReadLine() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}
	defer file.Close()

	// Erstelle einen Scanner, um die Datei zeilenweise zu lesen
	scanner := bufio.NewScanner(file)

	// Lese die Datei zeilenweise
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Überprüfe auf Fehler beim Lesen der Datei
	if err := scanner.Err(); err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
	}
}

func FileRead() {
	// Öffne die Datei
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}
	defer file.Close()

	// Bestimme die Größe der Datei
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Fehler beim Bestimmen der Dateigröße:", err)
		return
	}
	size := stat.Size()

	// Lese den Inhalt der Datei
	content := make([]byte, size)
	_, err = file.Read(content)
	if err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
		return
	}

	// Konvertiere den Inhalt in einen String und gib ihn aus
	text := string(content)
	fmt.Println(text)
}
