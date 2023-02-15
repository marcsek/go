package main

import (
	"fmt"
	"math/rand"
	"test/sorts"
	"time"
)

func main() {
	list := []int{}

	for i := 0; i < 10; i++ {
		list = append(list, genRandom(1000, 0))
	}
	fmt.Printf("pre. sort: %v\n", list)

	elapsed, _ := execTime(sorts.BubbleSort, &list)
	fmt.Printf("Exec time: %s\n", elapsed)

	fmt.Printf("pos. sort: %v", list)
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
