package main

import "fmt"

func main() {
	// #1
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }
	// #2
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i * i)
	// }
	// #3
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(3 * i)
	// }
	// #4
	// var n, first, second, next int
	// first, second = 0, 1

	// for n = 0; n < 10; n++ {
	// 	if n <= 1 {
	// 		next = n
	// 	} else {
	// 		next = first + second
	// 		first = second
	// 		second = next
	// 	}
	// 	fmt.Println(next)
	// }
	// #5
	// fmt.Println("Введите два числа")
	// var numOne, numTwo int
	// fmt.Scanln(&numOne, &numTwo)

	// fmt.Println(gcd(numOne, numTwo))
	// #6
	// for i := 0; i <= 100; i++ {
	// 	fmt.Println(i)
	// 	if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 	}
	// }
	// #7
	// fmt.Println("Введите число")
	// var num int
	// fmt.Scanln(&num)

	// if num%1 == 0 && num%num == 0 {
	// 	fmt.Println("Это число простое")
	// }
	// #8
	// number := 145

	// for i := 1; i <= number; i++ {
	// 	if number%i == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	// #9
	// var number int
	// fmt.Println("Введите число")
	// fmt.Scanln(&number)

	// sum := 0
	// for number > 0 {
	// 	a := number % 10
	// 	sum += a
	// 	number /= 10
	// }
	// fmt.Println(sum)
	// #10
	var num int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	for num < 0 {
		fmt.Println("Введите следующее число")
	}
	// for i := 0; i <= 20; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

}

// #5
// func gcd(numOne, numTwo int) int {
// 	for numTwo != 0 {
// 		(numOne, numTwo) = (numTwo, numOne%numTwo)
// 	}
// 	return numOne
// }
