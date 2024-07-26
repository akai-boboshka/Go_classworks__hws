/*package main

import "fmt"

func main() {
	// #1
	PreentGreeting("Вечер")
	// #2
	PrintWeather("облачно")
	// #3
	PrintTrafficLight("зелёный")
	// #4
	GetGrade(2)
	// #5
	GetDiscount(1200)
	// #6
	GetTemperatureDescription(9)
	// #7
	CheckNumber(10)
	// #8
	CheckAge(18)
	// #9
	CheckPassword(1234)
	// #10
	summ := Add
	fmt.Println(summ(10, 20))
	// #11
	CompareStrings("Привет", "Привет")
	// #12
	Max(10, 100)
	// #13
	var operation func(int, int) int = func(numOne, numTwo int) (difference int) {
		if numOne < 0 {
			fmt.Println("Используйте числа больше нуля")
		} else if numTwo < 0 {
			fmt.Println("Используйте числа больше нуля")
		}
		difference = numOne - numTwo
		return
	}
	fmt.Println(operation(-100, 200))
	// #14
	var concat func(string, string) string = func(stringOne, stringTwo string) string {
		var num string = "19090"
		var stringSumm string
		if stringSumm == num {
			fmt.Println("Нельзя использовать числа!")
		} else {
			stringSumm = stringOne + " " + stringTwo
			fmt.Println(stringSumm)
		}
		return stringSumm
	}
	fmt.Println(concat("Hello", "Hello"))
	// #15
	var multiply func(int, int) int = func(numOne, numTwo int) (multiplyOfNums int) {
		if numOne == 0 {
			fmt.Println("Нельзя умножать на ноль!")
		} else if numTwo == 0 {
			fmt.Println("Нельзя умножать на ноль!")
		}
		multiplyOfNums = numOne * numTwo
		return
	}
	fmt.Println(multiply(19, 0))
	// #16
	var ApplyOperation func(int, int) int = func(numOne, numTwo int) (difference int) {
		if numOne < 0 {
			fmt.Println("Используйте числа больше нуля")
		} else if numTwo < 0 {
			fmt.Println("Используйте числа больше нуля")
		}
		difference = numOne / numTwo
		return
	}
	fmt.Println(ApplyOperation(1000, 200))
	// #17
	CheckCondition(20, checkNum)

}

// #1
func PreentGreeting(dayType string) {
	var morning string = "Утро"
	var evening string = "Вечер"
	if dayType == morning {
		fmt.Println("Доброе утро!")
	} else if dayType == evening {
		fmt.Println("Добрый вечер!")
	}
}

// #2
func PrintWeather(weatherType string) {
	var sunny string = "солнечно"
	var cloudy string = "облачно"

	if weatherType == sunny {
		fmt.Println("Погода солнечная")
	} else if weatherType == cloudy {
		fmt.Println("Погода дождливая")
	}

}

// #3
func PrintTrafficLight(lightColor string) {
	var redColorOfTrafficLight = "красный"
	var yellowColorOfTrafficLight = "жёлтый"
	var greenColorOfTrafficLight = "зелёный"

	if lightColor == redColorOfTrafficLight {
		fmt.Println("Стоп")
	} else if lightColor == yellowColorOfTrafficLight {
		fmt.Println("Внимание")
	} else if lightColor == greenColorOfTrafficLight {
		fmt.Println("Идите")
	}
}

// #4
func GetGrade(score int) {
	if score == 1 {
		fmt.Println("A")
	} else if score == 2 {
		fmt.Println("B")
	} else if score == 3 {
		fmt.Println("C")
	}
}

// #5
func GetDiscount(amount int) {
	if amount > 1000 {
		fmt.Println("Скидка 10%")
	} else if amount > 500 && amount <= 100 {
		fmt.Println("Скидка 5%")
	} else if amount <= 500 {
		fmt.Print("Скидка 0%")
	}
}

// #6
func GetTemperatureDescription(temperature int) {
	if temperature < 10 {
		fmt.Println("На улице холодно")
	} else if temperature > 10 && temperature <= 25 {
		fmt.Println("На улице тепло")
	} else {
		fmt.Println("На улице жарко")
	}
}

// #7
func CheckNumber(num int) {
	if num > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число отрицательное")
	}
}

// #8
func CheckAge(age int) {
	if age < 18 {
		fmt.Println("Вы несовершеннолетний")
	} else {
		fmt.Println("Вы совершеннолетний")
	}
}

// #9
func CheckPassword(password int) {
	if password == 1234 {
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Пароль неверный")
	}
}

// #10
func Add(numOne, numTwo int) (sum int) {
	sum = numOne + numTwo
	return
}

// #11
func CompareStrings(stringOne, stringTwo string) {
	if stringOne == stringTwo {
		fmt.Println("Строки равны")
	} else {
		fmt.Println("Строки не равны")
	}
}

// #12
func Max(numOne, numTwo int) {
	if numOne > numTwo {
		fmt.Println(numOne)
	} else {
		fmt.Println(numTwo)
	}
}

// #17
func CheckCondition(num int, f func(int) bool) {
	if f(num) {
		fmt.Println("Условие выполнено")
	} else {
		fmt.Println("Условие не выполнено")
	}
}

func checkNum(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}*/

// package main

// import "fmt"

// func main() {
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
// var num int
// fmt.Println("Введите число")
// fmt.Scanln(&num)

// for i := 0; i <= 20; i++ {
// 	if i%2 == 1 {
// 		continue
// 	}
// 	fmt.Println(i)
// }

//}

// #5
// func gcd(numOne, numTwo int) int {
// 	for numTwo != 0 {
// 		(numOne, numTwo) = (numTwo, numOne%numTwo)
// 	}
// 	return numOne
// }

// package main

// import "fmt"

// func main() {
// 	var x int = 10
// 	var y *int // Объявление указателя

// 	y = &x // присвоение адреса переменной x указателю y
// 	fmt.Println(x)
// 	fmt.Println(&x)
// 	fmt.Println(y)
// 	fmt.Println(*y)
// }

// package main

// func main() {
// var ptr *int
// fmt.Println(ptr)
// var math, physics, chemistry int
// mathPtr := &math
// physicsPtr := &physics
// chemistryPtr := &chemistry

// addGrade(mathPtr, 89)
// addGrade(physicsPtr, 20)
// addGrade(chemistryPtr, 98)
// fmt.Printf("%.2f\n", printAverageGrade(mathPtr, physicsPtr, chemistryPtr))
// }

// #1
// func updateValue(p *int) {
// 	*p = *p + 10
// }

// #2
// func swap(x, y int) {
// 	*x, *y = *y, *x
// }

// #3
// func printValue(ptr *int) {
// 	if ptr == nil {
// 		fmt.Println("Указатель пуст")
// 	} else {
// 		fmt.Println("Значение, на которое указывает указатель", *ptr)
// 	}
// }

// Задача уровня мидл
// func vote(candidatePtr *string, votesPtr *int) {
// 	*votesPtr++
// }

// func winner(candidate1VotesPtr *int, candidate2VotesPtr *int) string {
// 	if *candidate1VotesPtr > *candidate2VotesPtr {
// 		return "Кандидат 1"
// 	} else if *candidate2VotesPtr > *candidate1VotesPtr {
// 		return "Кандидат 2"
// 	} else {
// 		return "Ничья"
// 	}
// }

// //
// func addExpense(totalPtr *float64, amount float64) {
// 	*totalPtr += amount
// }

// func printTotalExpense(totalPtr *float64) {
// 	fmt.Println("Общая сумма расходов: %.2f\n", *totalPtr)
// }

// //
// func addGrade(mathPtr *int, grade int) {
// 	*mathPtr = grade
// }

// func printAverageGrade(mathPtr *int, physicsPtr *int, chemistryPtr *int) float {
// 	sum := *mathPtr + *physicsPtr + *chemistryPtr
// 	return float64(sum) / 3.0
// }

package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 25
	fmt.Println(*p)
}
