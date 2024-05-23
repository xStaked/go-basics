package main

import (
	"container/list"
	"fmt"
	"reflect"
	"strings"

	"github.com/xstaked/go-basics/message"
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

	fmt.Println(asd)
	// fmt.Println(asd + strconv.Itoa(c))
	fmt.Printf("%s %d\n", asd, c)

	fmt.Println(reflect.TypeOf(d))

	fmt.Println((d + float32(c)))

	// short hand declaration of variables
	myString := "Hello, World!"
	fmt.Println(myString)
	// pi := 3.14159

	// fmt.Printf("d: %s PI: %f", "asdasd", pi)

	// var nombre string

	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&nombre)
	// fmt.Println("Hello", nombre)

	// aritmetica()
	// conditionals()

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

	// Inrange loop
	for index, value := range myMap {
		fmt.Println(index, value)
	}

	condition := 0

	// while loop
	for condition <= 5 {
		fmt.Println(condition)
		condition++
	}

	// Is palindrome
	// var word string
	// fmt.Print("Enter a word: ")
	// fmt.Scanln(&word)

	// for i := 0; i < len(word)/2; i++ {
	// 	if word[i] != word[len(word)-1-i] {
	// 		fmt.Println("Not a palindrome")
	// 		break
	// 	}
	// 	if i == len(word)/2-1 {
	// 		fmt.Println("Is a palindrome")
	// 	}
	// }

	// func isPalindrome(text string) {
	// 	var textReverse string
	// 	for i := len(text) - 1; i >= 0; i-- {
	// 		textReverse += string(text[i])
	// 	}

	// 	if text == textReverse {
	// 		fmt.Println("Es palindromo")
	// 	} else {
	// 		fmt.Println("No es palindromo")
	// 	}

	// }

	// Slices
	array2 := [3]int{5, 3, 32}
	fmt.Println(array2)

	var slice1 []string

	slice1 = append(slice1, "Hello")

	fmt.Println(slice1)

	// Functions
	// fmt.Println(myFunction())

	// Structs

	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{name: "John", age: 25}
	fmt.Println(myStruct)

	// isPalindrome()

	// Multiple return values
	// text, number := data()

	// fmt.Println(text, number)

	text := "Hello golang"
	fmt.Println(strings.Replace(text, "golang", "GO", 1))
	fmt.Println(strings.Split(strings.Join(strings.Split(text, " "), "-"), "-"))

	// Pointers
	var number int = 5
	var pointer *int = &number

	fmt.Println(number)
	fmt.Println(pointer)

	message.PublicFunc()
	message.AnotherFunc()

}

func myFunction() string {
	return "This is my function"
}

func aritmetica() {

	var a, b int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter another number: ")
	fmt.Scanln(&b)

	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b
	modulus := a % b

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)
	fmt.Printf("Modulus: %d\n", modulus)

	fmt.Println("Are they equal?", a == b)
	fmt.Println("Are they not equal?", a != b)
	fmt.Println("Is a greater than b?", a > b)
	fmt.Println("Is a less than b?", a < b)
	fmt.Println("Is a greater than or equal to b?", a >= b)
	fmt.Println("Is a less than or equal to b?", a <= b)
}

func conditionals() {

	var (
		n      int
		output string
	)

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	switch n {
	case 1:
		output = "One"
	case 2:
		output = "Two"
	case 3:
		output = "Three"
	default:
		output = "Not 1, 2, or 3"
	}

	fmt.Println(output)
}

func isPalindrome() {
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scanln(&word)
	var textReverse string
	for i := len(word) - 1; i >= 0; i-- {
		textReverse += string(word[i])
	}

	if word == textReverse {
		fmt.Println("Is a palindrome")
	} else {
		fmt.Println("Is not a palindrome")
	}
}

func data() (string, int) {
	return "Hello", 5
}
