package main

import (
	"errors"
	"math"
	"strconv"
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

// Whether x is a palindrome.
func isPalindrome(x int) bool {
	ret := true
	xRunes := []rune(strconv.Itoa(x))
	lenXRunes := len(xRunes)
	for i := range xRunes {
		if xRunes[i] != xRunes[lenXRunes-i-1] {
			ret = false
			break
		}
	}
	return ret
}

// The largest palindrome made from the product of two n-digit numbers.
// Returns a non-nil error if no palindrome found.
// Project Euler prob 4.
func maxProduct2NDigitInts(nDigits int) (int, error) {
	ub := int(math.Pow10(nDigits)) - 1
	lb := int(math.Pow10(nDigits - 1))
	max := 0
	for i := ub; i >= lb; i-- {
		for j := i; j >= lb; j-- {
			p := i * j
			if isPalindrome(p) {
				if p > max {
					max = p
				}
			}
		}
	}
	var err error
	if max == 0 {
		err = errors.New("No palindromes found")
	}
	return max, err
}

func main() {
}
