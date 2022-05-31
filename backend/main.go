package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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
	now := time.Now()
	rand.Seed(now.UnixNano())
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
	availableNumbers := getNumbers(49)
	return func() (field, bool) {
		if len(availableNumbers) < 6 {
			return nil, false
		}
		field := make([]int, 6)
		for i := range field {
			n := rand.Intn(len(availableNumbers))
			field[i] = (availableNumbers)[n]
			availableNumbers = append((availableNumbers)[:n], (availableNumbers)[n+1:]...)
		}
		sort.Ints(field)
		return field, true
	}
}

func getNumbers(len int) []int {
	numbers := make([]int, len)
	for i := range numbers {
		numbers[i] = i + 1
	}
	return numbers
}
