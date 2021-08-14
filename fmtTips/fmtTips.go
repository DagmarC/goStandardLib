package fmtTips

import (
	"bufio"
	"fmt"
	"os"
)

// Coloring output--------------------------------------
type MessageType int

// Coloring output
const (
	INFO MessageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

// Coloring output
const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

// FmtTips --------------------------------------------------------------
func FmtTips() {
	fmt.Println("--------FMT-------")
	// Table formatting
	fmt.Printf("|%7.2f|%7.2f|%7.2f\n", 23.222, 4.22, 5555.4)
	fmt.Printf("|%7.2f|%7.2f|%7.2f", 232.22, 4.223, 55.4)

	// Coloring output--------------------------------------------------
	showMessage(INFO, "Opening a file.")
	showMessage(WARNING, "If file error then fail.")

	file, err := os.Open("test.txt")
	if err != nil {
		showMessage(ERROR, err.Error())
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			showMessage(ERROR, err.Error())
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	//------------------------------------------------------------------

}

//--------------------------------------------------------------
// Coloring output
func showMessage(messageType MessageType, message string) {
	switch messageType {
	case INFO:
		printMessage := fmt.Sprintf("\nInfo: \n%s\n", message)
		fmt.Printf(InfoColor, printMessage)
	case WARNING:
		printMessage := fmt.Sprintf("\nWarning: \n%s\n", message)
		fmt.Printf(WarningColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: \n%s\n", message)
		fmt.Printf(ErrorColor, printMessage)
	}
}

//--------------------------------------------------------------
