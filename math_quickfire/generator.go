package main

import (
	"fmt"
	"math/rand"
)

type OperationType int

const (
	Addition       OperationType = iota // 0
	Subtraction                         // 1
	Multiplication                      // 2
	Division                            // 3
)

func promptUser() {

	var n int
	fmt.Println("Enter the number of problems you wish to solve")
	fmt.Scan(&n)

	if n == 1 {
		fmt.Printf("You will have %d problem to solve \n", n)
	} else {
		fmt.Printf("You will have %d problems to solve \n", n)
	}

	printQuestions(n)

}

func generateRandomMultiplicationQuestion() (int, int, int) {

	x := rand.Intn(15)
	y := rand.Intn(15)

	fmt.Printf("%d * %d = ", x, y)

	return x, y, int(Multiplication)

}

func generateRandomAdditionQuestion() (int, int, int) {

	x := rand.Intn(200)
	y := rand.Intn(200)

	fmt.Printf("%d + %d = ", x, y)

	return x, y, int(Addition)

}

func generateRandomSubtractionQuestion() (int, int, int) {

	x := rand.Intn(200)
	y := rand.Intn(200)

	fmt.Printf("%d - %d = ", x, y)

	return x, y, int(Subtraction)

}

func generateRandomDivisionQuestion() (int, int, int) {

	y := rand.Intn(14) + 1
	answer := rand.Intn(15)

	x := y * answer

	fmt.Printf("%d / %d = ", x, y)

	return x, y, int(Division)

}

func chooseRandomGeneratorOperation() (int, int, int) {
	n := rand.Intn(4)
	switch n {
	case 0:

		return generateRandomAdditionQuestion()

	case 1:
		return generateRandomSubtractionQuestion()
	case 2:
		return generateRandomMultiplicationQuestion()
	case 3:
		return generateRandomDivisionQuestion()
	}
	return 0, 0, 0

}

func validateAnswer(x, y, n, answer int) bool {
	switch n {
	case 0:
		return x+y == answer
	case 1:
		return x-y == answer
	case 2:
		return x*y == answer
	case 3:
		return x/y == answer
	}
	return false
}

func printQuestions(n int) {
	correctCount := 0
	var isCorrect bool
	for i := 0; i < n; i++ {
		fmt.Printf("Q%d): ", i + 1)
		x, y, opCode := chooseRandomGeneratorOperation()
		var userAnswer int
		fmt.Scan(&userAnswer)
		isCorrect = validateAnswer(x, y, opCode, userAnswer)

		if isCorrect {
			fmt.Println("Correct!")
			correctCount++
		} else {
			fmt.Println("Incorrect!")
		}

	}

	fmt.Printf("\nGame Over! You got %d out of %d questions correct!\n", correctCount, n)
}
