package main

import (
	"fmt"
	"os"
)

type (
	Stack []int64
)

func (s *Stack) Push(i int64) {
	(*s) = append((*s), i)
}

func (s *Stack) Pop() int64 {
	var i int64 = (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return i
}

func (s *Stack) Remove(i int) {
	var right []int64
	if (len(*s)-1) > i {
		right = (*s)[i+1:]
	}
	(*s) = (*s)[:i]
	(*s) = append((*s), right...)
}

func prime(primes []int64) int64 {
	// determine next prime
	var i int64
primer:
	for i = (int64)(primes[len(primes)-1]); ; i++ {
		var prime = true
	loopy:
		for j := 0; j < len(primes); j++ {
			if i%primes[j] == 0 {
				prime = false
				break loopy
			}
		}
		if prime == true {
			break primer
		}
	}
	return i
}

func main() {
	var i int64 // counter
	var primes []int64 = []int64{2}
	var n int64                  // number of plates
	var q int64                  // "queries"
	fmt.Scanf("%d %d", &n, &q)   // parse n & q
	var plates Stack = Stack{}   // setup holder for all plates
	var noplates Stack = Stack{} // setup second stack
	for i = 0; i < n; i++ {
		var z int64
		_, err := fmt.Scanf("%d", &z)
		if err != nil {
			break
		}
		plates.Push(z)
	}
	for i = 0; i < n; i++ {
		noplates.Push(plates.Pop()) // reverse
	}
	plates = noplates
	noplates = Stack{}
	var piles []int64       // setup final pile
	for i = 0; i < q; i++ { // for all q
		primes = append(primes, prime(primes)) // get new prime
		for j := 0; 0 < len(plates); j++ {     // for all plates
			now := plates.Pop() // get the plate
			if now%primes[i] == 0 { // if plate is divisible by the current prime
				piles = append(piles, now) // add plate to pile
			} else {
				noplates.Push(now) // add plate back in
			}
		}
		plates = noplates  // reset
		noplates = Stack{} // clear
	}
	for j := 0; j < len(piles); j++ {
		fmt.Fprintf(os.Stdout, "%d\n", piles[j])
	}
	for j := 0; j < len(plates); j++ {
		fmt.Fprintf(os.Stdout, "%d\n", plates[j])
	}
}
