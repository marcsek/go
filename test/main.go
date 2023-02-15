package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := []int{}

	for i := 0; i < 10; i++ {
		list = append(list, genRandom(10, 0))
	}
	fmt.Printf("pre. sort: %v\n", list)

	elapsed, _ := execTime(bubbleSort, &list)
	fmt.Printf("Exec time: %s\n", elapsed)

	fmt.Printf("pos. sort: %v", list)
}

func bubbleSort(listPtr *[]int) bool {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(*listPtr)-1; i++ {
			if (*listPtr)[i] > (*listPtr)[i+1] {
				swap(&(*listPtr)[i], &(*listPtr)[i+1])
				swapped = true
			}
		}
	}
	return true
}

func execTime[In any, Out any](x func(In) Out, p In) (time.Duration, Out) {
	start := time.Now()
	x(p)
	return time.Since(start), x(p)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func genRandom(max int, min int) int {
	return rand.Intn(max-min+1) + min
}
