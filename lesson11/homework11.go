package main

// // •	Реализуйте функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий числа, которые встречаются больше одного раза.

// // func GetSlice(slice []int) []int {
// // 	newSlice := []int{}
// // 	for i := 0; i <= len(slice); i++ {
// // 		for j := i + 1; j <= len(slice); j++ {
// // 			if slice[i] == slice[j] {
// // 				newSlice = append(newSlice, slice[i])
// // 			}
// // 		}
// // 	}
// // 	return newSlice
// // }

// // func main() {
// // 	slice := []int{10, 20, 10, 30, 45, 89, 2, 4, 30}
// // 	duplicates := GetSlice(slice)
// // 	fmt.Println(duplicates)
// // }

// //Реализуйте функцию, которая принимает срез структур Book и возвращает структуру Book, у которой год публикации самый новый.

// type Book struct {
// 	YearOfIssue []int
// }

// func getBook(book Book) (newBook Book) {
// 	// if book[0] > book[1] {

// 	// }
// 	return newBook
// }

// #1

// func sumOfStrings(strOne, strTwo string) (result string) {
// 	return fmt.Sprintf(strOne + strTwo)
// }

// func main() {
// 	result := sumOfStrings("Hello ", "всем")
// 	fmt.Println(result)
// }

// #2
// func stringlength(str string) string {

// 	return fmt.Sprintln("Длина строки: ", len(str))
// }

// func main() {
// 	lenOfStr := stringlength("Hello world")
// 	fmt.Println(lenOfStr)
// }

// #3
// func checkString() {
// 	str := "Привет, как дела?"
// 	fmt.Println(strings.Contains(str, "дела"))
// }

// func main() {
// 	checkString()
// }

// #4

// func FindIndex() {
// 	str := "Hello, World"
// 	findIn := strings.Contains(str, "Hello")
// 	fmt.Println(findIn)
// }

// func main() {
// 	FindIndex()
// }

// #6
// func Delete() {
// 	str := "  Hello,  Ibragim "
// 	delete := strings.Trim(str, " ")
// 	fmt.Println(delete)
// }

// func main() {
// 	Delete()
// }

// #10

// func Compare() {
// 	strOne := "Hello"
// 	strTwo := "hello"

// 	if strOne > strTwo {
// 		fmt.Println(strOne)
// 	} else {
// 		fmt.Println(strTwo)
// 	}
// }

// func main() {
// 	Compare()
// }

// #14

// func deleteLetter() {
// 	str := "Hello, Boboshka"
// 	delete := strings.ReplaceAll(str, "H", "")
// 	fmt.Println(delete)
// }

// func main() {
// 	deleteLetter()
// }

// #20

// func StringToInt() {
// 	str := "10, 20, 30"
// 	stringToInt := strconv.ParseInt(str,)
// }
