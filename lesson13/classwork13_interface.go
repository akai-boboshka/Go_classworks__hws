package main

import "fmt"

// type User struct {
// 	Name string
// 	Age  int
// }

// func (u *User) GetUserInfo() string {
// 	return fmt.Sprintf("Name: %s, Age: %d", u.Name, u.Age)
// }

// func main() {
// 	user := User{
// 		Name: "Ibragim",
// 		Age:  18,
// 	}

// 	fmt.Println(user.GetUserInfo())
// }

// type Vehicle interface {
// 	move()
// }

// // структура "Автомобиль"
// type Car struct{}

// // структура "Самолет"
// type Aircraft struct{}

// func (c Car) move() {
// 	fmt.Println("Автомобиль едет")
// }
// func (a Aircraft) move() {
// 	fmt.Println("Самолет летит")
// }

// func main() {

// 	var tesla Vehicle = Car{}
// 	var boing Vehicle = Aircraft{}
// 	tesla.move()
// 	boing.move()
// }

// type Food interface {
// 	meal()
// }

// type Fruit struct{}
// type Meat struct{}

// func (f Fruit) meal() {
// 	fmt.Println("Apple, Pear")
// }

// func (m Meat) meal() {
// 	fmt.Println("beef, chicken")
// }

// func main() {
// 	var fruits Food = Fruit{}
// 	var meat Food = Meat{}

// 	fruits.meal()
// 	meat.meal()
// }
// type Animal interface {
// 	Speak() string
// }

// type Dog struct{}

// func (d Dog) Speak() string {
// 	return "Woof"
// }

// type Cat struct{}

// func (c Cat) Speak() string {
// 	return "Meow"
// }

// // Функция для работы с любым Animal
// func Describe(a Animal) {
// 	fmt.Println(a.Speak())
// }

// func duck() {
// 	dog := Dog{}
// 	cat := Cat{}

// 	Describe(dog)
// 	Describe(cat)
// }

// func main() {
// 	duck()
// }

// func main() {
// 	// Срез значений любого типа
// 	slice := []interface{}{42, "hello", true}
// 	for _, v := range slice {
// 		fmt.Println(v)
// 	}

// 	// Карта значений любого типа
// 	m := map[string]interface{}{
// 		"age":       30,
// 		"name":      "Alice",
// 		"isStudent": false,
// 	}
// 	for k, v := range m {
// 		fmt.Printf("%s: %v\n", k, v)
// 	}
// }

// type Voice interface {
// 	AnimalsVoice() string
// }

// type Dog struct{}

// func (d Dog) AnimalsVoice() string {
// 	return "Woof"
// }

// func Describe(v Voice) {
// 	fmt.Println(v.AnimalsVoice())
// }

// func duck() {
// 	dog := Dog{}
// 	Describe(dog)
// }

// func main() {
// 	duck()
// }

type Voice interface {
	AnimalsVoice() string
}

type Cow struct{}
type Tiger struct{}

func (c Cow) AnimalsVoice() string {
	return "Moo"
}

func (t Tiger) AnimalsVoice() string {
	return "Wrrr"
}

// func Describe(v Voice) {
// 	fmt.Println(v.AnimalsVoice())
// }

// func duck() {
// 	cow := Cow{}
// 	tiger := Tiger{}
// 	Describe(cow)
// 	Describe(tiger)
// }

func main() {
	// duck()
	// var v Voice
	// v = Cow{}
	// fmt.Println(v.AnimalsVoice())

	GetInterface(1213)
}

func GetInterface(i interface{}) {
	fmt.Println(i)
}
