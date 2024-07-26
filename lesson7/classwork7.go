// package main

// import "fmt"

// func main() {
// 	var myAge Age = 25
// 	fmt.Println(myAge)
// 	var myAdder Adder = add
// 	result := myAdder(2, 3)
// 	fmt.Println(result)
// 	//
// 	p := Person{Name: "Alice", Age: 30}
// 	fmt.Println(p.Name)
// 	fmt.Println(p.Age)
// }

// type Age int

// // Adder
// // Определяем тип функции
// type Adder func(int, int) int

// // Реализуем функцию, соответствующую типу Adder
// func add(a int, b int) int {
// 	return a + b
// }

// // Создание структуры
// type Person struct {
// 	Name string
// 	Age  int
// }

// package main

// import "fmt"

// type Address struct {
// 	Country string
// 	City    string
// }
// type Person struct {
// 	Name    string
// 	Age     int
// 	Address Address
// }

// func main() {
// 	p := Person{
// 		Name: "Ибрагим",
// 		Age:  18,
// 		Address: Address{
// 			Country: "Tajikistan",
// 			City:    "Dushanbe",
// 		},
// 	}
// 	fmt.Println("Ибрагим живёт в", p.Address.City)
// }

// package main

// import "fmt"

// func main() {
// 	aboutCar := Cars{
// 		Model:          "W 211",
// 		YearOfCreation: 2008,
// 		EngineСapacity: 3.5,
// 	}
// 	fmt.Println("Here is about Mercedes car: ", "Модель:", aboutCar.Model, "Год выпуска:", aboutCar.YearOfCreation, "Объём двигателя:", aboutCar.EngineСapacity)
// }

// type Cars struct {
// 	Model          string
// 	YearOfCreation int
// 	EngineСapacity float64
// }

// Определение возраста совершеннолетия
// Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том,
// является ли человек совершеннолетним (возраст 18 лет и старше) или нет.
// package main

// import "fmt"

// type Age int

// func GetAge(age Age) string {
// 	// fmt.Scan("Введите свой возраст:", &age)
// 	if age <= 18 {
// 		message := "Вы несовершеннолетний :("
// 		return message
// 	} else {
// 		messageOne := "Вы совершеннолетний, поздравляю!"
// 		return messageOne
// 	}
// }

// func main() {
// 	checkAge := GetAge(10)
// 	fmt.Println(checkAge)
// }

// Проверка на четность
// Определите тип Number на основе int. Напишите функцию, которая принимает число и возвращает сообщение о том,
// является ли оно четным или нечетным.

// package main

// import (
// 	"fmt"
// )

// type Num int

// func main() {
// 	checkNum := checkNum(13939321)
// 	fmt.Println(checkNum)
// }

// func checkNum(num Num) string {
// 	if num%2 == 0 {
// 		text := "Число чётное"
// 		return text
// 	} else {
// 		textSecond := "Число нечётное"
// 		return textSecond
// 	}
// }

// Арифметические операции
// Определите тип функции Operation, которая принимает два int и возвращает int. Напишите функции для сложения,
// вычитания и умножения. Используйте тип Operation для вызова этих функций

package main

// type Operation func

// Обратный отсчет
// Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает Count и выводит
// обратный отсчет до нуля.

// type Count = int

// func coundown(num Count) {
// 	for i := num; i >= 0; i-- {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	coundown(24)
// }

// Проверка температуры
// Создайте псевдоним для типа float64 под названием Temperature. Напишите функцию, которая принимает
// Temperature и выводит сообщение о состоянии (ниже нуля, выше нуля или ноль).

// type Temperature = float64

// func checkTemperature(temperature Temperature) {
// 	if temperature > 0 {
// 		fmt.Println("Температура на данный момент выше нуля")
// 	} else if temperature < 0 {
// 		fmt.Println("Температура на данный момент ниже нуля")
// 	} else {
// 		fmt.Println("Сейчас температура равна нулю")
// 	}
// }

// func main() {
// 	checkTemperature(10.2)
// }

// Применение скидки
// Создайте псевдоним для типа float64 под названием Price. Напишите функцию, которая принимает Price
// и возвращает новую цену с 10% скидкой.

// type Price = float64

// func PriceWithDiscount(price Price) float64 {
// 	result := price - ((price * 10) / 100)
// 	return result
// }

// func main() {
// 	finalPrice := PriceWithDiscount(2500)
// 	fmt.Println(finalPrice)
// }

// Продукт и его стоимость
// Создайте структуру Product с полями Name (строка) и Price (тип Price). Напишите функцию, которая
// принимает Product и возвращает сообщение о его стоимости.

// func main() {
// 	pr := Product{
// 		Name:  "dfdgg",
// 		Price: 10.10,
// 	}
// 	fmt.Println(pr)
// }

// type Price float64

// type Product struct {
// 	Name  string
// 	Price Price
// }

// func ProductPrice(aboutProduct Product) string {
// 	return fmt.Sprintf("%f", aboutProduct.Price)
// }

//Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том, является ли человек подростком (возраст от 13 до 19 лет включительно) или нет.
// type Age = int

// func checkAge(age Age) (message string) {
// 	if age >= 13 && age <= 19 {
// 		message = "Вы подросток"
// 		return
// 	} else {
// 		message = "Вы не подросток"
// 		return
// 	}
// }

// func main() {
// 	returnFunc := checkAge(10)
// 	fmt.Println(returnFunc)
// }

// Определите тип BMI на основе float64. Напишите функцию, которая принимает значение BMI и возвращает сообщение о том, в какой категории оно находится: "Underweight", "Normal weight", "Overweight", или "Obesity".
// type BMI = float64

// func checkBMI(BMISecond BMI) {
// 	if BMISecond < 18.5 {
// 		fmt.Println("Underweight")
// 	} else if BMISecond >= 18.5 && BMISecond <= 24.9 {
// 		fmt.Println("Normal weight")
// 	} else if BMISecond >= 25 && BMISecond <= 29.9 {
// 		fmt.Println("Overweight")
// 	} else {
// 		fmt.Println("Obesity")
// 	}
// }

// func main() {
// 	checkBMI(189)
// }

//Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функцию для возведения числа в квадрат и используйте тип UnaryOperation для вызова этой функции.
// import "fmt"

// type UnaryOperation func(int) int

// func squaring(num int) int {
// 	return num * num
// }

// func main() {
// 	var multiply UnaryOperation = squaring
// 	result := multiply(100)
// 	fmt.Println(result)
// }

// Определите тип функции Check, которая принимает int и возвращает bool. Напишите функцию для проверки, является ли число положительным, и используйте тип Check для вызова этой функции.
// import "fmt"

// type Check func(int) bool

// func checkNum(num int) bool {
// 	if num > 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func main() {
// 	var CheckFunc Check = checkNum
// 	fmt.Println(CheckFunc(10))
// }

// Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает Count и выводит обратный отсчет до нуля.
// import "fmt"

// type Count = int

// func countdown(num Count) {
// 	for i := num; i >= 0; i-- {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	countdown(14)
// }

// Создайте псевдоним для типа float64 под названием Weight. Напишите функцию, которая принимает Weight и возвращает сообщение о том, в какой категории веса находится объект: "Light", "Medium" или "Heavy".
// import "fmt"

// type Weight = float64

// func checkWeight(weightSecond Weight) {
// 	if weightSecond < 20 {
// 		fmt.Println("Light")
// 	} else if weightSecond > 20 && weightSecond < 85 {
// 		fmt.Println("Medium")
// 	} else {
// 		fmt.Println("Heavy")
// 	}
// }
// func main() {
// 	checkWeight(100)
// }

// Создайте псевдонимы для типов float64 под названиями BMI и BloodPressure. Напишите функцию, которая принимает BMI и BloodPressure, и возвращает сообщение о состоянии здоровья: "Healthy", "At risk" или "Unhealthy".
// import "fmt"

// type BMI = float64
// type BloodPressure = float64

// func checkHealth(num BMI, secondNum BloodPressure) {
// 	if num >= 18.5 && num <= 24.9 && secondNum <= 120 {
// 		fmt.Println("Healthy")
// 	} else if num >= 25 && num <= 29.9 && secondNum < 135 {
// 		fmt.Println("Unhealthy")
// 	} else {
// 		fmt.Println("At risk")
// 	}
// }

// func main() {
// 	checkHealth(20, 100)
// }

// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает Product и выводит его описание.
// import "fmt"

// type Product struct {
// 	Name  string
// 	Price int
// }

// func main() {
// 	AboutProduct := Product{
// 		Name:  "Сок",
// 		Price: 14,
// 	}
// 	fmt.Printf("Название продукта: %s \n Стоимость: %d сомон \n", AboutProduct.Name, AboutProduct.Price)
// }

//Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает два Product и возвращает более дорогой продукт.
// package main
// import "fmt"

// type Product struct {
// 	Name  string
// 	Price int
// }

// func priceComparison(ProductOne Product, ProductSecond Product) Product {
// 	if ProductOne.Price > ProductSecond.Price {
// 		return ProductOne
// 	} else {
// 		return ProductSecond
// 	}
// }

// func main() {
// 	ProductOne := Product{
// 		Name:  "Яблоки",
// 		Price: 10,
// 	}
// 	ProductSecond := Product{
// 		Name:  "Мясо",
// 		Price: 70,
// 	}
// 	result := priceComparison(ProductOne, ProductSecond)
// 	fmt.Println(result)
// }
//
// import "fmt"

// func main() {
// 	var numbers [5]int = [5]int{1, 2, 10}
// 	fmt.Println(numbers[1])
// 	numbers[1] = 100
// 	fmt.Println(numbers[1])
// }

// Создайте структуру Item с полями Name и Quantity. Напишите функцию, которая принимает массив Item и возвращает суммарное количество всех предметов.

// type Item struct {
// 	Name string
// 	Quantity int
// }

// func function(item []int) {
// 	item =
// }

// func main() {
// 	firstItem := Item {
// 		Name: "Мерседес",
// 		Quantity: 10,
// 	}
// 	SecondItem := Item {
// 		Name: "Volkswagen",
// 		Quantity: 37,
// 	}
// 	ThirdItem := Item {
// 		Name: "Buggati",
// 		Quantity: 5,
// 	}

// }

// import "fmt"

// func main() {
// 	// var users []string = []string{"Tom", "Alice", "Kate"}
// 	// fmt.Println(users[2]) // Kate
// 	// users[2] = "Katherine"

// 	// for _, value := range users {
// 	// 	fmt.Println(value)
// 	// }
// 	var numbers []int = []int{1, 20, 23}
// 	// fmt.Println(numbers[1])
// 	sum := 0

// 	for _, num := range numbers {
// 		sum += num
// 	}
// 	fmt.Println(sum)
// }

//Создайте структуру Person с полями Name и Age. Напишите функцию, которая принимает указатель на Person и увеличивает его возраст на 1.
// type Person

// func main() {
// 	users := []string{"Ibragim", "Boboshka", "Sport", "Mercedes", "BMW", "Bobo", "Viktor"}
// 	var n = 3
// 	users = append(users[:n], users[n+3:]...)
// 	fmt.Println(users)
// users1 := users[1:3]
// fmt.Println(users1)

// for _, values := range users {
// 	fmt.Println(values)
// }
// }

// func main() {
// 	num := 1
// 	ptr := new(int)
// 	ptr = &num
// 	fmt.Println(*ptr)
// }

// type person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	user := person{Name: "Bobo", Age: 20}
// 	var userPointer *person = &user
// 	userPointer.Age = 22
// 	fmt.Println(user.Age)
// }
