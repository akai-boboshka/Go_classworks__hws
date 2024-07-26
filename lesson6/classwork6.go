package main

// func createPointer(x int) *int {
// 	c := new(int)
// 	c = &x
// 	*c = x
// 	fmt.Println(c)
// 	return c
// }

func main() {
	// c1 := createPointer(27)
	// fmt.Println(*c1)
	// fmt.Println(c1)
	// printValue()
	// currentBalance()
}

// var initialBalance int = 100000
// var sum *int = &initialBalance

// func currentBalance() {
// 	fmt.Scan("Введите сумму пополнения: ", &*sum)
// 	result := initialBalance + *sum
// 	fmt.Println(result)
// }

// func printValue(num *int) {
// 	*num = 10
// 	if num == nil {
// 		fmt.Println("Указатель пуст")
// 	} else {
// 		fmt.Println("Указатель, на которое указывает указатель", *num)
// 	}
// }

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
// sayHello()

// defolt value
// var IntPointer *int
// fmt.Printf("%T %#v \n", IntPointer, IntPointer)

// var a int64 = 7
// fmt.Printf("%T %#v \n", a, a)

// var pointerA *int64 = &a
// fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA)

// var newPointer = new(float32)
// fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
// *newPointer = 10
// fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
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

// func printAverageGrade(mathPtr *int, physicsPtr *int, chemistryPtr *int) float64 {
// 	sum := *mathPtr + *physicsPtr + *chemistryPtr
// 	return float64(sum) / 3.0
// }
