package main

import (
	"fmt"
	"math/rand"
	"time"
)

type iRange []int

func main() {

	// 1. Create a slice of ints from 0 through 10
	numberRange := newRange()
	numberRange.print()

	//2. Check if each element is even or odd
	numberRange.oddOrEven()
}

func newRange() iRange {
	var numberRange iRange
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Millisecond) //Sleep for 1ms to generate different numbers
		numberRange = append(numberRange, randomNumber(999999))
	}
	return numberRange
}

func (ir iRange) oddOrEven() {
	for _, number := range ir {
		if number%2 == 0 {
			fmt.Println(number, "is even.")
		} else {
			fmt.Println(number, "is odd.")
		}
	}
}

func (ir iRange) print() {
	fmt.Println("Number range initialized to:", ir)
}

func randomNumber(maxLimit int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(maxLimit)
}
