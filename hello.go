package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	// println("Hello, World!")
	fmt.Println("Hello, World!")

	// vars
	var a int = 1
	var b int = 2
	var c int = a + b
	var d float32 = 1.1

	println(d)

	var asd string = "asddd"

	fmt.Println(c)

	fmt.Println(asd)
	// fmt.Println(asd + strconv.Itoa(c))
	fmt.Printf("%s %d\n", asd, c)

	fmt.Println(reflect.TypeOf(d))

	fmt.Println((d + float32(c)))

	// short hand declaration of variables
	myString := "Hello, World!"
	fmt.Println(myString)

	const myConst string = "Hello, World!"

	if c == 3 || c > 2 {
		fmt.Println("c is 3")
	} else if c == 2 {
		fmt.Println("c is 2")
	} else {
		fmt.Println("c is something else")
	}

	// Arrays
	var myArray [3]int
	var myArray2 [3]string

	for i := 0; i < len(myArray); i++ {
		myArray[i] = i
	}

	fmt.Println(myArray)
	fmt.Println(myArray2)

	// Maps

	//  map[string]int{"one": 1, "three": 3, "two": 2}
	myMap := make(map[string]int)

	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	fmt.Println(myMap)
	fmt.Println(myMap["one"])

	// Lists
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	// Loops

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for index, value := range myMap {
		fmt.Println(index, value)
	}

	// Functions
	fmt.Println(myFunction())

	// Structs

	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{name: "John", age: 25}
	fmt.Println(myStruct)

}

func myFunction() string {
	return "This is my function"
}
