package main

import (
	"fmt"
	"sort"
)

func main() {
	// 	// var arr [5]int = [5]int{1, 20, 13, 20, 100}
	// 	// fmt.Println(arr[2])

	// 	// for _, values := range arr {
	// 	// 	fmt.Println(values)
	// 	// }
	// 	// slice := arr[1:3]
	// 	// fmt.Println(slice)

	// 	// slice := []int{20, 29, 10, 24}
	// 	// slice = append(slice, )

	// 	// slice1 := []int{1, 2, 3}
	// 	// slice2 := []int{5, 10}
	// 	// slice3 := append(slice1, slice2...)
	// 	// fmt.Println(len(slice3), cap(slice3))
	// 	// slice := []int{1, 2, 3, 5}
	// 	// index := 3 // индекс для вставки
	// 	// value := 4
	// 	// slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
	// 	// fmt.Println(slice) // [1 2 3 4 5]

	// 	// var numbers [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}

	// 	// for _, val := range numbers {
	// 	// 	fmt.Println(val)
	// 	// }

	// 	// Создаем массив из 7 целых чисел
	// 	// var arr [7]int

	// 	// Инициализируем массив значениями от 1 до 7
	// 	// for i := 0; i < 7; i++ {
	// 	// 	arr[i] = i + 1
	// 	// }

	// 	// fmt.Println(arr)

	// }

	// import {
	// 	"fmt"
	// }

	// func main() {
	// arrays()
	// slices()
	// variadicFunction()
	// convertToArrayPointer()
	// passToFunction()
	// sliceWithNew()
	// }

	// func arrays() {
	// 	type Person struct {
	// 		Age  int
	// 		Name string
	// 	}
	// 	var intArr [5]int
	// fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)

	// intArr[0] = 10
	// intArr[1] = 20
	// fmt.Println(intArr)

	// people := [2]Person{
	// 	{
	// 		Age:  10,
	// 		Name: "Nicka",
	// 	},
	// 	{
	// 		Age:  23,
	// 		Name: "Vika",
	// 	},
	// }
	// fmt.Println(people)

	// stringsArr := [...]string{"First", "Second", "Third", "Fourth"}
	// fmt.Printf("Length: %d Capacity: %d\n", len(stringsArr), cap(stringsArr))

	// var arrOne [4]int = [4]int{10, 20, 17, 278}
	// fmt.Printf("Length: %d Capacity: %d\n", len(arrOne), cap(arrOne))
	// for i := 0; i < len(arrOne); i++ {
	// 	fmt.Println(arrOne[i])
	// }

	// for index, value := range arrOne {
	// 	fmt.Printf("Index: %d Value: %d\n", index, value)
	// }

	// newArrOne := changeArray(&arrOne)
	// fmt.Println(arrOne)
	// fmt.Println(newArrOne)
	// }

	// func changeArray(arr *[4]int) [4]int {
	// 	arr[2] = 99
	// 	return *arr
	// }

	// func slices() {
	// var defaultSlice []int
	// fmt.Printf("Length: %d Capacity: %d\n", len(defaultSlice), cap(defaultSlice))

	// stringSliceLiteral := []string{"Vika", "Andrey", "Sergey"}
	// fmt.Printf("Length: %d Capacity: %d\n", len(stringSliceLiteral), cap(stringSliceLiteral))

	// slice := make([]int, 0, 5)
	// fmt.Printf("Value: %#v\n", slice)
	// fmt.Printf("Length: %d Capacity: %d\n", len(slice), cap(slice))
	// slice = append(slice, 1, 2, 3, 4, 5)
	// fmt.Printf("Value: %#v\n", slice)
	// fmt.Printf("Length: %d Capacity: %d\n", len(slice), cap(slice))
	// slice = append(slice, 6)
	// fmt.Printf("Value: %#v\n", slice)
	// fmt.Printf("Length: %d Capacity: %d\n", len(slice), cap(slice))
	// }

	// Функция с неограниченным числом параметров
	// func variadicFunction() {
	// 	showAllElements(1, 20, 30, 99)

	// 	firstSlice := []int{10, 20, 30}
	// 	secondSlice := []int{40, 50, 60, 70, 80}
	// 	// showAllElements(firstSlice...)

	// 	newSlice := append(firstSlice, secondSlice...)
	// 	fmt.Printf("Value: %#v\n", newSlice)
	// }

	// func showAllElements(values ...int) {
	// 	for _, val := range values {
	// 		fmt.Println(val)
	// 	}
	// 	fmt.Println()
	// }

	// Конвертация слайса в указатель на массив. Нужно, чтобы совпадала длина

	// func convertToArrayPointer() {
	// 	initialSlice := []int{1, 2}
	// 	fmt.Printf("Type: %T Value: %#v\n", initialSlice, initialSlice)
	// 	fmt.Printf("Length: %d Capacity: %d\n\n", len(initialSlice), cap(initialSlice))

	// 	intArray := (*[2]int)(initialSlice)
	// 	fmt.Printf("Type: %T Value: %#v\n", intArray, intArray)
	// 	fmt.Printf("Length: %d Capacity: %d\n", len(intArray), cap(intArray))
	// }

	// Передача слайса в функцию

	// func passToFunction() {
	// 	slice := []int{1, 2} // тут слайс ссылатеся на массив состоящий из двух элементов и который вмещает 2 элемента
	// 	fmt.Printf("Type: %T Value: %#v\n", slice, slice)
	// 	fmt.Printf("Length: %d Capacity: %d\n\n", len(slice), cap(slice))
	// 	changeValue(slice)
	// 	fmt.Printf("Type: %T Value: %#v\n", slice, slice)
	// 	fmt.Printf("Length: %d Capacity: %d\n\n", len(slice), cap(slice))

	// 	newSlice := append(slice, 10) // тут уже, слайс ссылается на новый массив который уже вмещает 4 элемента
	// 	fmt.Printf("Type: %T Value: %#v\n", newSlice, newSlice)
	// 	fmt.Printf("Length: %d Capacity: %d\n\n", len(newSlice), cap(newSlice))

	// 	newSliceTwo := appendValue(newSlice)
	// 	fmt.Printf("Type:%T Value: %#v\n", newSliceTwo, newSliceTwo)
	// 	fmt.Printf("Length: %d Capacity: %d\n\n", len(newSliceTwo), cap(newSliceTwo))
	// }

	// func changeValue(sliceToChange []int) {
	// 	sliceToChange[0] = 20
	// }

	// func appendValue(sliceOne []int) []int {
	// 	sliceOne = append(sliceOne, 25, 40)
	// 	fmt.Printf("Type: %T Value: %#v\n", sliceOne, sliceOne)
	// 	fmt.Printf("Length: %d Capacity: %d\n\n", len(sliceOne), cap(sliceOne))

	// 	return sliceOne
	// }

	// Создание слайса через new

	// func sliceWithNew() {
	// 	slicePointer := new([]int)

	// 	fmt.Printf("Type: %T Value: %#v\n", slicePointer, slicePointer)
	// 	fmt.Printf("Length: %d Capacity: %d\n\n", len(*slicePointer), cap(*slicePointer))
	// }

	//

	// func main() {
	// sliceFloat64 := []float64{4.4, 1.2, -2.8, 19.09}
	// fmt.Println(sort.Float64sAreSorted(sliceFloat64)) // Проверяем, является ли слайс отсортированным
	// sort.Float64s(sliceFloat64)
	// fmt.Println(sliceFloat64, sort.Float64sAreSorted(sliceFloat64))
	// fmt.Println(sort.SearchFloat64s(sliceFloat64, 1.2)) // Ищем значение в слайсе, выводит индекс этого значения

	// s := []string{"one", "two", "three"}
	// sort.Strings(s)
	// fmt.Println(s, sort.StringsAreSorted(s))

	// Сортировка string
	peopleInfo := []struct {
		Name string
		Age  int
	}{
		{"Nick", 20},
		{"Bobo", 18},
		{"Jamal", 45},
		{"Boboshka", 19},
	}
	sort.Slice(peopleInfo, func(i, j int) bool { return peopleInfo[i].Age < peopleInfo[j].Age })
	fmt.Println(peopleInfo)
}
