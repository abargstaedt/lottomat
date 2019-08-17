package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	allBalls := [49]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	remainingBalls := allBalls[:]
	for f := 0; f < 8; f++ {
		var field [6]int
		for k := 0; k < 6; k++ {
			n := rand.Intn(len(remainingBalls))
			field[k] = remainingBalls[n]
			remainingBalls = append(remainingBalls[:n], remainingBalls[n+1:]...)
		}
		sorted := field[:]
		sort.Ints(sorted)
		fmt.Println("Feld", f+1, ":", sorted)
	}
	fmt.Println("Übrig :", remainingBalls)
}
