package main

import "fmt"

// func main() {
// definition()
// var lib library = library{"Book1", "Book2", "Book3", "Book4"}
// lib.print()
// var Boboshka = person{Name: "Бобошка", Age: 18}
// Boboshka.print()

// }

// type Square struct {
// 	Side int
// }

// func (s Square) Perimetr() {
// 	fmt.Printf("%T, %#v \n", s, s)
// 	fmt.Printf("Периметр фигуры: %d \n", s.Side*4)
// }

// func definition() {
// 	square := Square{Side: 13}
// 	// ptrSquare := &square

// 	square.Perimetr()
// }

// type library []string

// func (l library) print() {
// 	for _, values := range l {
// 		fmt.Println(values)
// 	}
// }

// type person struct {
// 	Name string
// 	Age  int
// }

// func (p person) print() {
// 	fmt.Println("Имя:", p.Name)
// 	fmt.Println("Возраст:", p.Age)
// }

// type Rectangle struct {
// 	Width, Height int
// }

// func (r Rectangle) Area() int {
// 	return r.Width * r.Height
// }

// func main() {
// 	rect := Rectangle{Width: 10, Height: 33}
// 	area := rect.Area()
// 	fmt.Printf("Площадь прямоугольника равна = %d", area)

// 	num := MyInt(4)
// 	fmt.Println(num.IsEven)
// }

// func (r *Rectangle) scale(factor int) {
// 	r.Width *= factor
// 	r.Height *= factor
// }

// type MyInt interface

// func (m MyInt) isEven() bool {
// 	return m%2 == 0
// }

/*
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func main() {
	tempC := Celsius(32)
	tempF := tempC.ToFahrenheit()
	fmt.Println(tempF)
}
*/
// Определение структуры Person
// type Person struct {
// 	Name string
// 	Age  int
// }

// Функция NewPerson действует как конструктор для структуры Person
// Она инициализирует и возвращает новый экземпляр Person
// func NewPerson(name string, age int) *Person {
// 	return &Person{
// 		Name: name,
// 		Age:  age,
// 	}
// }

// func main() {
// 	person := NewPerson("Ibragim", 18)
// 	fmt.Println(person)
// }

// Вызываем анонимную функцию внутри другой функции
// func give_me_a_func() func(string) {
// 	return func(message string) {
// 		fmt.Println(message)
// 	}
// }

// func main() {
// 	sendFunction := give_me_a_func()
// 	sendFunction("Yes, We did it! :)")
// }

// func SummOfNum() func(int, int) {
// 	return func(numOne, numTwo int) {
// 		fmt.Println(numOne * numTwo)
// 	}
// }

// func main() {
// 	summ := SummOfNum()
// 	summ(10, 20)
// }

// Замыкание
// func incrementor() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

// func main() {
// 	// next := incrementor()
// 	// fmt.Println(next())
// 	// fmt.Println(next())
// 	// fmt.Println(next())
// 	// fmt.Println(next())
// 	// fmt.Println(next())

// 	nextOne := inc()
// 	fmt.Println(nextOne())
// 	fmt.Println(nextOne())
// 	fmt.Println(nextOne())
// 	fmt.Println(nextOne())
// 	fmt.Println(nextOne())
// }

// func inc() func() int {
// 	num := 0
// 	return func() int {
// 		num++
// 		return num
// 	}
// }

// Структуры

type User struct {
	Name     string
	Age      int
	Password string
	Score    []int
}

// func (u User) isElder() bool {
// 	a := u.Age
// 	isTrue := false

// 	if a >= 18 {
// 		isTrue = true
// 	} else if a < 18 {
// 		isTrue = false
// 	}

// 	return isTrue
// }

// func (u User) getHighScore() int {
// 	highScore := 0
// 	for _, value := range u.Score {
// 		if highScore < value {
// 			highScore = value
// 		}
// 	}
// 	return highScore
// }

func (u User) GetTheHighestScore() int {
	highScore := 0
	for _, value := range u.Score {
		if highScore < value {
			highScore = value
		}
	}
	return highScore
}

func main() {
	user := User{"John", 18, "Tojikiston ba Pesh!", []int{1000, 67, 100, 10, 299}}
	// fmt.Println(user.getHighScore())
	fmt.Println(user.GetTheHighestScore())
	// fmt.Println(user.isElder())
	// if user.isElder() {
	// 	fmt.Println("Заходи!")
	// } else {
	// 	fmt.Println("Упс, вход запрещён!")
	// }
}

// func (u User) CheckAge() bool {
// 	age := u.Age
// 	isTrue := false

// 	if age >= 18 {
// 		isTrue = true
// 	} else if age < 18 {
// 		isTrue = false
// 	}

// 	return isTrue
// }

// func main() {
// 	user := User{"Ibragim", 18, "Ibragim1234"}
// 	fmt.Println(user.CheckAge())
// }
