package reflection

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type person struct {
	personId int
	name     string
}

type customer struct {
	customerId int
	name       string
	company    string
}

func ReflectionPackage() {

	newPerson := person{0, "Fred"}

	// Reflection
	fmt.Printf("The type is: %v \n", reflect.TypeOf(newPerson))
	fmt.Printf("The value is: %v \n", reflect.ValueOf(newPerson))
	fmt.Printf("The kind is: %v \n", reflect.ValueOf(newPerson).Kind())
	fmt.Println("--------")

	newCustomer := customer{
		customerId: 1,
		name:       "Alojz",
		company:    "SWI",
	}
	// Reflection usage: 1 person for both types.
	addPerson(newCustomer)
	addPerson(newCustomer)

	// Slice with reflection
	customers := make([]customer, 3)
	customers = append(customers, customer{
		customerId: 1,
		name:       "Jane",
		company:    "SAP",
	})
	customers = append(customers, customer{
		customerId: 2,
		name:       "Josh",
		company:    "SAP",
	})
	customers = append(customers, customer{
		customerId: 3,
		name:       "Lucas",
		company:    "SQIW",
	})

	customerType := reflect.TypeOf(customers)
	newCustomersList := reflect.MakeSlice(customerType, 0, 0)
	newCustomersList = reflect.Append(newCustomersList, reflect.ValueOf(customer{
		customerId: 4,
		name:       "Reflect",
		company:    "Artin",
	}))
	fmt.Println()
	fmt.Printf("Original list: %v\n\n", customers)
	fmt.Printf("Via reflection list: %v\n\n", newCustomersList)

	fmt.Println("--COOKBOOK--")
	ourTitle := "the go standard library"
	newTitle := properTitle(ourTitle)
	fmt.Println(newTitle)

	fmt.Println("=====TIMED FUNCTION=====")
	// properTitle function as a parameter to time its execution.
	timed := MakeTimedFunction(properTitle).(func(string) string)
	newTitle = timed(ourTitle)
	fmt.Println(newTitle)

}

// With reflection you can create one function for both customer and person.
func addPerson(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		value := reflect.ValueOf(p)

		switch reflect.TypeOf(p).Name() {
		case "person":
			personSqlString := "INSERT INTO person (pesonId, name) VALUES (?,?)"
			fmt.Println(personSqlString)
			fmt.Printf("Added %v", value.Field(1)) // real values would being inserted.
		case "customer":
			customerSqlString := "INSERT INTO person (customerId, name, company) VALUES (?,?,?)"
			fmt.Println(customerSqlString)
			fmt.Printf("Added %v", value.Field(1)) // real values would being inserted.

		}
		return true
	}
	return false
}

// properTitle Cookbook
func properTitle(input string) string {
	words := strings.Fields(input)
	smallWords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallWords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

// MakeTimedFunction Creates our function with reflection that we
// will use to time the function being passed in - Good for timing of execution
func MakeTimedFunction(f interface{}) interface{} {
	reflection := reflect.TypeOf(f)

	if reflection.Kind() != reflect.Func {
		panic("Expects a function.")
	}
	valueOfFunction := reflect.ValueOf(f)

	wrapperFunction := reflect.MakeFunc(reflection, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := valueOfFunction.Call(in) // Passed arg into the function.
		end := time.Now()

		fmt.Printf("Callin %s took %v\n", runtime.FuncForPC(valueOfFunction.Pointer()).Name(), end.Sub(start).Nanoseconds())

		return out
	})
	return wrapperFunction.Interface()
}
