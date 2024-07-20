package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else {
		return multiply
	}
}

func main() {
	var f = selectFn(2)
	fmt.Println(f(10, 99))

	// #1.1
	PrintGreeting()
	// #1.2
	DisplaySeparator()
	// #1.3
	NotifyStart()
	// #2.1
	GetWelcomeMessage()
	// #2.2
	GetPiValue(3.14)
	// #2.3
	IsServerActive(true)
	// #3.1
	PrintNumber(100)
	// #3.2
	GreetUser("Ibragim")
	// #3.3
	ToggleLight(true)
	// #4.1
	add(10, 89)
	// #4.2
	ConCat("Hello, world!")
	// #4.3
	IsEven(500)
	// #5.1
	var adder func(int, int) int = add
	fmt.Println(adder(10, 89))
	// #5.2
	var concatenator func(string, string) string = addOne
	fmt.Println(concatenator("Ibragim", "hello"))
	// #5.3
	var isPositive func(int) bool = Number
	fmt.Println(isPositive(900))
	// #6.1
	Calculate(10, 20, summ)
	// #6.2
	Execute(false)
	// #6.3
	Apply(10)
	// #7.4

}

// #1.1
func PrintGreeting() {
	fmt.Println("Hello world!")
}

// #1.2
func DisplaySeparator() {
	fmt.Println("--------------------")
}

// #1.3
func NotifyStart() {
	fmt.Println("Program started")
}

// #2.1
func GetWelcomeMessage() (text string) {
	text = "Welcome!"
	return
}

// #2.2
func GetPiValue(Pi) float32 {
	return Pi
}

// #2.3
func IsServerActive(a) bool {
	return a
}

// #3.1
func PrintNumber(a int) {
	fmt.Println(a)
}

// #3.2
func GreetUser(name string) {
	fmt.Println("Привет, ", name)
}

// #3.3
func ToggleLight(a bool) {
	if a == true {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}

// #4.1
func add(a, b int) (result int) {
	result = a + b
	return
}

// #4.2
func ConCat(stringOne, stringTwo string) (result string) {
	result = stringOne + " " + stringTwo
	return
}

// #4.3
func IsEven(a int) bool {
	if a / 2 {
		return true
	} else {
		return false
	}
}

// #5.1
func add(a, b int) int {
	return a + b
}

// #5.2
func addOne(str1, str2 string) string {
	return str1 + str2
}

// #5.3
func Number(num int) bool {
	if num > 0 {
		return true
	} else {
		return false
	}
}

// #6.1
func summ(x, y int) int {
	return x + y
}
func Calculate(numOne, numTwo int, action func(int, int) int) {
	result := action(numOne, numTwo)
	fmt.Println(result)
}

// #6.2
func Execute(value bool, f func(bool)) {
	a := f(value)
	fmt.Println(a)
}

// #6.3
func Apply(number int, fu func(int) int) {
	return fu(number)
}

// #7.1
func Multiplier(factor int) func(numFour int) int {
	return func(numFour int) int {
		return numFour * factor
	}
}

// #7.2
func StringerRepeater(n int) func(str string) string {
	return func(str string) string {
		return stringOne.Repeat(str, n)
	}
}
