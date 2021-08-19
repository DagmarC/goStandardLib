package stringTips

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func StringTips() {
	fmt.Println("\n----------STRINGS--------")

	// COMPARE---------------------------------------------------------
	stooges := []string{"Lorem", "Ipsum", "Not"}

	for _, stooge := range stooges {

		// strings.Compare - faster than ==

		// Returns 0 if they match
		// 1-not match
		// -1-lexikographic order differs
		fmt.Println(strings.Compare("Lorem", stooge))
	}

	// SPLIT----------------------------------------------------------
	fmt.Println("\t--SPLIT")

	stringsCollectionWithout := strings.Split("Hello|I can| run", "|")
	stringsCollectionWith := strings.SplitAfter("Hello|I can| run", "|")

	for i := range stringsCollectionWithout {
		fmt.Println(stringsCollectionWithout[i])
		fmt.Println(stringsCollectionWith[i])
		fmt.Println("---")
	}
	// Split into 2 parts (==after first split into the half, stops splitting)
	stringsCollectionParts := strings.SplitAfterN("Hello|I can| run", "|", 2)
	for i := range stringsCollectionParts {
		fmt.Println(stringsCollectionParts[i])
	}
	// LOG PARSER and REPLACE
	fmt.Println("---LOG PARSER--")

	const searchTerm = "ERROR"

	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		result := strings.HasPrefix(line, searchTerm)
		if result {
			//fmt.Println(line)
			// -1 means Replace All
			replacedString := strings.Replace(line, searchTerm, "REPLACED", -1)
			fmt.Println(line, replacedString)
		}
	}

	// REGEX
	fmt.Println("---REGEX--")

	text := "somg are great songs"

	r, _ := regexp.Compile(`s(\w[a-z]+)g\b`)
	fmt.Println(r.MatchString(text))
	fmt.Println(r.FindAllString(text, -1))
	fmt.Println(r.FindStringIndex(text))

	newText := r.ReplaceAllString(text, "apple")
	fmt.Println(newText)

	// TRIM
	ttt := "   jjj  "
	newTtt := strings.TrimSpace(ttt)
	fmt.Printf("Before: %q and After: %q\n", ttt, newTtt)
	fmt.Println(strings.TrimPrefix("https://asdf.com", "https://"))
}
