package main

import (
	"math"
)

// Sum the distinct multiples of a and b that are less than n.
// Project Euler prob 1.
// Implementation is in linear time (0(n)).
func sumDistinctMultsLinear(a, b, n int) int {
	sum := 0
	for i := a; i <= n; i += a {
		sum += i
	}
	for i := b; i <= n; i += b {
		sum += i
	}
	for i := a * b; i <= n; i += a * b {
		sum -= i
	}
	return sum

}

// Sum the even Fibonacci numbers that don't exceed max.
func sumEvenFibs(max int) int {
	prevprev, prev := 1, 2

	sum := prev
	for {
		cur := prevprev + prev
		if cur > max {
			break
		}
		if cur%2 == 0 {
			sum += cur
		}
		prevprev, prev = prev, cur
	}
	return sum
}

// Biggest prime factor of n.
// Proj Euler prob 3.
func biggestPrimeFactor(n int) int {
	var bpf int

	// Sieve of Eratosthenes (O(log n)) to find primes up to sqrt(n)
	floorSqrt := int(math.Floor(math.Sqrt(float64(n))))
	notPrimeMask := make([]bool, floorSqrt+1) // elements default to false
	for i := 2; i <= floorSqrt; i++ {
		if notPrimeMask[i] {
			continue
		}
		for j := i; j <= floorSqrt; j += i {
			notPrimeMask[j] = i != j
		}
	}

	for val, isNotPrime := range notPrimeMask {
		if val < 2 || isNotPrime {
			continue
		}
		// Do we have a prime factor?
		if n%val == 0 {
			// Yes
			n /= val
			bpf = val
			if n == 1 {
				// Fully factorised
				break
			}
		}
	}

	return bpf
}

func main() {
}
