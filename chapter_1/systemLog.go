package chapter_1

import (
	"log"
	"os"
)

func SystemLog() {
	// Usually a log file would be placed in the temp directory like this:
	// LOGFILE := path.Join(os.TempDir(), "masteringGo.log")

	// Basic logging options. Create the file if it doesn't exist yet, write only access and append to current content.
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println("This log message will be written to the file.")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Log with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Log with file/line")

	myLogger := log.New(file, "Prefix: ", log.LstdFlags)
	myLogger.Println("from myLogger")

	log.Panicln("Panic log. This also terminates the program with status 2 (log.Fatal also terminates, but with status 1)")
}
