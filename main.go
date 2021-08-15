package main

import (
	"goStandardLib/fmtTips"
	"goStandardLib/logTips"
	"goStandardLib/timeTips"
	//"strconv"
)

func main() {
	//fmt.Printf("Hello %v \n", runtime.Version())
	//args := os.Args[1:]

	// FLAG - booleans, strings, ...
	// has own data types
	// arch - name of the flag
	// x86 - default value
	// CPU Type - description - see via goStandardLib.exe -h

	//archPtr := flag.String("arch", "x86", "CPU Type")

	//flag.Parse()
	// to access values of the flag - use dereference
	// goStandardLib.exe -arch AMD64 2 3
	//switch *archPtr {
	//case "x86":
	//	fmt.Println("32 bit mode")
	//case "AMD64":
	//	fmt.Println("64 bit mode")
	//}

	//if len(args) == 1 && args[0] == "/help" {
	//	fmt.Println("Enter the price of meal and amount of tip.\n" +
	//		"e.g. goStandardLib.exe 11 2")
	//}
	//if len(args) != 2 {
	//	fmt.Println("Enter 2 arguments or run /help.")
	//	return
	//}
	// Args are strings by default.
	//mealTotal, _ := strconv.ParseFloat(args[0], 32)
	//tipAmount, _ := strconv.ParseFloat(args[1], 32)
	//fmt.Printf("Total: %.2f", calculateTotal(mealTotal, tipAmount))

	// FMT
	fmtTips.FmtTips()

	// LOGS
	logTips.LogsTips()

	// TIME
	timeTips.TimeTips()
}

func calculateTotal(mealTotal, amount float64) float64 {
	total := mealTotal + (mealTotal * (amount / 100))
	return total
}
