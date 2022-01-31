package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func printFibonacci(id int) {
	numbers := []uint64{1, 1}
	for i := 0; i < 50; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		n := numbers[0] + numbers[1]
		numbers[0] = numbers[1]
		numbers[1] = n

		fmt.Printf("Number %d: %d\n", i+1, n)
	}

	log.Fatalf("Routine %d finished first.\n", id)
}

/*func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go printFibonacci(1)
	go printFibonacci(2)

	wg.Wait()
}*/
