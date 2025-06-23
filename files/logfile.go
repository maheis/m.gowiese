// Package files enthält Beispiele für fileoperations
package files

import (
	"log"
	"os"
)

func LogFile() {
	f, err := os.OpenFile("file.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fLog := log.New(f, "File: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	fLog.Println("Hello File")

	eLog := log.New(os.Stderr, "Stderr: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	eLog.Println("Hello Stderr")

	sLog := log.New(os.Stdout, "Stdout: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	sLog.Println("Hello Stdout")

	// fLog.Fatal("Fatal")
	// fLog.Panic("Panic")
	// Ruft man .Fatal oder .Panic auf, wird das Programm beendet

	sLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	sLog.Println("Stdout another Flags")
	sLog.SetOutput(eLog.Writer())
	sLog.Println("Change from Stdout to Stderr")
}
