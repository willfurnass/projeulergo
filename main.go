package main

import (
	"errors"
	"fmt"
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

// Sum of all multiples of d less or equal to ub.
func sumSeqConst(d int, ub int) int {
	n := ub / d
	ret := n * ((2 * d) + ((n - 1) * d)) / 2
	return ret

}

// Sum the distinct multiples of a and b that are less than n.
// Project Euler prob 1.
// Implementation is in constant time (0(1)).
func sumDistinctMultsConst(a, b, n int) int {
	return sumSeqConst(a, n) + sumSeqConst(b, n) - sumSeqConst(a*b, n)
}

// Sum the even Fibonacci numbers that don't exceed max.
// Project Euler prob 2.
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

// Greatest Common Divisor of a and b.
// Implemented using Euclid's algorithm.
func gcd(a, b int) int {
	var ret int
	x, y := a, b
	if b < a {
		x, y = b, a
	}
	for {
		rem := y % x // 171
		if rem == 0 {
			ret = x
			break
		}
		x, y = rem, x
	}
	return ret
}

// Absolute value of x.
func absInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Lowest Common Multiple of a and b.
// Implemented using gcd (Greatest Common Multiple) function.
func lcm(a, b int) int {
	return absInt(a * b) / gcd(a, b)
}

// Lowest Common Multiple of all natural numbers up to and including n.
// Proj Euler prob 5.
func lcmRange(n int) int {
	if n < 3 {
		return n
	}
	accum := 2
	for i := 3; i <= n; i++ {
		accum = lcm(accum, i)
	}
	return accum
}

// a to the power of b.
func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))

}

// Difference between the sum of the squares of the first n natural numbers and the square of the sum of those natural numbers.
// Proj Euler prob 6.
func diffSumSqsSqSums(n int) int {
	sumseq := sumSeqConst(1, n)
	sumSqs := ((2 * n * n * n) + (3 * n * n) + n) / 6
	return (sumseq * sumseq) - sumSqs
}

func main() {
}
