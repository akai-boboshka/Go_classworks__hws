package main

func main() {
	// Максимальное значение массива
	// var arr []int = []int{10, 20, 1, 2}
	// sort.Ints(arr)
	// fmt.Println(arr[len(arr)-1])

	// max := arr[0]
	// for _, element := range arr {
	// 	if element > max {
	// 		max = element
	// 	}
	// }
	// fmt.Println(max)

	// Минимальное значение массива
	// arr := []int{10, 20, 1, 49}
	// sort.Ints(arr)
	// fmt.Println(arr[0])

	// Подсчитать кол-во положительных элементов в массиве
	// var arr [5]int = [5]int{10, 9, -10, 23, -23}
	// var newArr []int
	// for _, element := range arr {
	// 	if element < 0 {

	// 	} else {
	// 		newArr = append(newArr, element)
	// 		fmt.Println(len(newArr))
	// 	}
	// }

	// Сумма всех элементов массива
	// var arr [8]int = [8]int{99, 20, 43, 40}
	// arrSum := 0
	// for _, element := range arr {
	// 	arrSum += element
	// }
	// fmt.Printf("Сумма элементов массива равна: %d\n", arrSum)

	// Среднее значение всех элементов
	// var arr [8]int = [8]int{99, 20, 43, 40}
	// arrSum := 0
	// for _, element := range arr {
	// 	arrSum += element
	// }
	// fmt.Printf("Среднее значение элементов массива равна: %d\n", arrSum/len(arr))

	// Найти все индексы заданного числа
	// var arr [8]int = [8]int{20, 13, 24, 54, 72, 89, 54}
	// num := 54
	// for index, value := range arr {
	// 	if num == value {
	// 		fmt.Printf("Это число находится под индексом: %d\n", index)
	// 	}
	// }

	// Умножить все элементы массива на заданное число
	// var arr [4]int = [4]int{1, 2, 3, 4}
	// arrMultiply := 2
	// for _, element := range arr {
	// 	arrMultiply *= element
	// }
	// fmt.Println(arrMultiply)

	// Второе максимальное число массива
	// var arr []int = []int{10, 20, 1, 2}
	// sort.Ints(arr)
	// fmt.Println(arr[len(arr)-2])

	// Поменять местами максимальное и минимальное значение
	// var arr []int = []int{10, 20, 1, 2}
	// sort.Ints(arr)
	// arr[len(arr)-1], arr[0] = arr[0], arr[len(arr)-1]
	// fmt.Println(arr)

	// Найти дубликаты и удалить их
	// slice := []int{20, 13, 24, 54, 72, 89, 54}
	// num := 54
	// for index, value := range slice {
	// 	if num == value {
	// 		slice = append(slice[:index], slice[index+1:]...)
	// 		fmt.Println(slice)
	// 	}
	// }

	//
	// slice := []int{1, 2, 3, 10, 34, 290}
	// slice2 := []int{4, 5, 6}
	// slice = append(slice, slice2...)
	// fmt.Println(slice)
	// slice3 := slice[2:5]
	// fmt.Println(slice3)

	// src := []int{1, 2, 3}
	// dst := make([]int, len(src))
	// n := copy(dst, src)
	// fmt.Println(n)
	// fmt.Println(dst)

	// №26 Найти максимальное произведение трех чисел
	// arr := []int{20, 39, 18, 47, 89, 32, 56, 100}
	// sort.Ints(arr)
	// fmt.Println(arr)
	// fmt.Println(arr[len(arr)-1] * arr[len(arr)-2] * arr[len(arr)-3])

	// #27
	// arr := []int{20, 39, 18, 47, 89, 32, 56, 100}
	// sort.Ints(arr)
	// slice := arr[len(arr)-4 : len(arr)]
	// sumOfSliceElements := 0
	// for _, element := range slice {
	// 	sumOfSliceElements += element
	// }
	// fmt.Println(sumOfSliceElements)

}
