package LogTips

import (
	"fmt"
	"goStandardLib/fmtTips"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func LogsTips() {
	// You can either use the way example below or create Logger instances for each level.
	// 2nd way: var InfoLogger *log.Logger; InfoLogger = log.New(file, "INFO: ", ... )

	fmt.Println("\n-----LOGS-------")
	log.Println("Log message.")

	writeLog(fmtTips.INFO, "Good")
	writeLog(fmtTips.ERROR, "Oh no")
	//writeLog(fmtTips.FATAL, "Dead") // Note: Fatal will kill the program.

	// TRACING
	fmt.Println("\n-----TRACE-------")

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("No file created: %v\n", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Failed to close a file %v.\n", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("Error of trace start. %v.\n", err)
	}

	defer trace.Stop()
	AddRandomNumbers()

}

func writeLog(messageType fmtTips.MessageType, message string) {

	// Open is not that good as OpenFile
	// - permission issues wth writing into the file.

	// OpenFile() will write log into it.
	// - has options like, create|append| ...
	// Sets permissions like 0666

	//file, err := os.Open("log.txt")
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// Sets output of the log.
	log.SetOutput(file)

	switch messageType {
	case fmtTips.INFO:
		// Ldate - local date, Ltime - local time
		// Lshortfile - file name a the line of code that it was executed on
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)

	case fmtTips.WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)

	case fmtTips.ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)

	case fmtTips.FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message)
	}
}

// TRACING
func AddRandomNumbers() {
	firstNumber := rand.Intn(100)
	secondNumber := rand.Intn(100)

	time.Sleep(2 + time.Second)
	var result = firstNumber * secondNumber

	fmt.Printf("Result %d*%d = %d\n", firstNumber, secondNumber, result)
}
