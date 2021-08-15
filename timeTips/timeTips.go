package timeTips

import (
	"fmt"
	"log"
	"os"
	"time"
)

var timeLogger *log.Logger

func TimeTips() {
	timeLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Wall vs Monotonic clock
	fmt.Println("\n------TIME------")
	today := time.Now()

	fmt.Println(today.Format(time.ANSIC))
	fmt.Println(today.Format(time.Kitchen))

	//OWN FORMAT, needs to have special signs
	fmt.Println(today.Format("Mon 2 Jan 15:04:05 MST 2006"))
	fmt.Println(today.Format("2 Monday in the year 2006"))
	// Not exact - bad days
	//fmt.Println(today.Format("3 Tuesday in the year 2006"))

	startDate := time.Date(2021, 10, 02, 9, 00, 00, 00, time.UTC)
	fmt.Println(startDate)

	// Mon 2 Jan 15:04:05 MST 2006
	// TIME SPANS
	riskoDate := time.Date(2020, 11, 02, 10, 04, 00, 00, time.UTC)

	elapsedTime := time.Since(riskoDate)
	fmt.Printf("%.0f Hours, %.0f Minutes and %.0f Seconds has passed since Risko`s birth %v.\n",
		elapsedTime.Hours(), elapsedTime.Minutes(), elapsedTime.Seconds(), riskoDate.Format("Monday, 2 January 2006"))

	writeTime(today, "1st elapse") // today is a start

	// Date 6 months from now
	// PAST: -6
	future := today.AddDate(0, 6, 0)
	fmt.Printf("Future: %v\n", future.Format(time.RFC822))

	// PERIOD of X hours, Deadlines, ...:
	p := today.Add(6 * time.Hour)
	fmt.Println("Until period: ", time.Until(p).Hours())

	defer writeTime(today, "End of program")
}

func writeTime(start time.Time, name string) {
	elapsed := time.Since(start)
	timeLogger.Printf("----------%s took %s----------\n", name, elapsed)
}
