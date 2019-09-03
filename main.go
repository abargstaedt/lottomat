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

type ticket []field

func (t ticket) String() (s string) {
	for index, field := range t {
		s += fmt.Sprintf("Field %d: %v\n", index+1, field)
	}
	return
}

func main() {
	rand.Seed( /* put your lucky number here: */ 13)
	fmt.Println(getTicket())
}

func getTicket() (t ticket) {
	nextField := fieldGenerator()
	for {
		field, ok := nextField()
		if !ok {
			break
		}
		t = append(t, field)
	}
	return
}

func fieldGenerator() func() (field, bool) {
	availableNumbers := make([]int, 49)
	for i := range availableNumbers {
		availableNumbers[i] = i + 1
	}
	return func() (field, bool) {
		if len(availableNumbers) < 6 {
			return nil, false
		}
		numbers := make([]int, 6)
		for i := range numbers {
			n := rand.Intn(len(availableNumbers))
			numbers[i] = (availableNumbers)[n]
			availableNumbers = append((availableNumbers)[:n], (availableNumbers)[n+1:]...)
		}
		sort.Ints(numbers)
		return numbers, true
	}
}
