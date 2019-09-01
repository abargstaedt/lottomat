package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type field []int

func (f field) String() string {
	var str string
	for _, number := range f {
		str += fmt.Sprintf(" %02d", number)
	}
	return str
}

func main() {
	allNumbers := [49]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	remainingNumbers := allNumbers[:]
	for index, field := range getFields(&remainingNumbers) {
		fmt.Println("Feld", index+1, ":", field)
	}
	fmt.Println("Übrig :", remainingNumbers)
}

func getFields(remainingNumbers *[]int) []field {
	fields := make([]field, 8)
	for i := 0; i < 8; i++ {
		fields[i] = drawNumbers(remainingNumbers)
	}
	return fields
}

func drawNumbers(remainingNumbers *[]int) []int {
	numbers := make([]int, 6)
	for i := 0; i < 6; i++ {
		n := rand.Intn(len(*remainingNumbers))
		numbers[i] = (*remainingNumbers)[n]
		*remainingNumbers = append((*remainingNumbers)[:n], (*remainingNumbers)[n+1:]...)
	}
	sort.Ints(numbers)
	return numbers
}
