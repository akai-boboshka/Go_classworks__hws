package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println(ptr)
	var math, physics, chemistry int
	mathPtr := &math
	physicsPtr := &physics
	chemistryPtr := &chemistry

	addGrade(mathPtr, 89)
	addGrade(physicsPtr, 20)
	addGrade(chemistryPtr, 98)
	fmt.Printf("%.2f\n", printAverageGrade(mathPtr, physicsPtr, chemistryPtr))
}

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
func vote(candidatePtr *string, votesPtr *int) {
	*votesPtr++
}

func winner(candidate1VotesPtr *int, candidate2VotesPtr *int) string {
	if *candidate1VotesPtr > *candidate2VotesPtr {
		return "Кандидат 1"
	} else if *candidate2VotesPtr > *candidate1VotesPtr {
		return "Кандидат 2"
	} else {
		return "Ничья"
	}
}

//
func addExpense(totalPtr *float64, amount float64) {
	*totalPtr += amount
}

func printTotalExpense(totalPtr *float64) {
	fmt.Println("Общая сумма расходов: %.2f\n", *totalPtr)
}

//
func addGrade(mathPtr *int, grade int) {
	*mathPtr = grade
}

func printAverageGrade(mathPtr *int, physicsPtr *int, chemistryPtr *int) float64 {
	sum := *mathPtr + *physicsPtr + *chemistryPtr
	return float64(sum) / 3.0
}
