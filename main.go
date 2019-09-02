package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type field []int

func (f field) String() (s string) {
	for _, number := range f {
		s += fmt.Sprintf(" %02d", number)
	}
	return
}

type ticket struct {
	fields           []field
	remainingNumbers []int
}

func (t ticket) String() (s string) {
	for index, field := range t.fields {
		s += fmt.Sprintf("Field %d: %v\n", index+1, field)
	}
	if len(t.remainingNumbers) > 0 {
		s += "Remaining: "
		for _, number := range t.remainingNumbers {
			s += fmt.Sprintf(" %02d", number)
		}
	}
	return
}

func main() {
	rand.Seed(42)
	fmt.Println(getTicket())
}

func getTicket() (t ticket) {
	allNumbers := [49]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	t.remainingNumbers = allNumbers[:]
	for {
		t.fields = append(t.fields, drawNumbers(&t.remainingNumbers))
		if len(t.remainingNumbers) < 6 {
			break
		}
	}
	return
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
