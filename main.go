package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func newField(remainingBalls *[]int) []int {
	field := make([]int, 6)
	for i := 0; i < 6; i++ {
		n := rand.Intn(len(*remainingBalls))
		field[i] = (*remainingBalls)[n]
		*remainingBalls = append((*remainingBalls)[:n], (*remainingBalls)[n+1:]...)
	}
	sort.Ints(field)
	return field
}

func main() {
	allBalls := [49]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	remainingBalls := allBalls[:]
	for f := 0; f < 8; f++ {
		fmt.Println("Feld", f+1, ":", newField(&remainingBalls))
	}
	fmt.Println("Übrig :", remainingBalls)
}
