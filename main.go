package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
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

// Sieve of Eratosthenes (O(log n)) to find primes up to n.
func sieve(n int) *[]bool {
	notPrimeMask := make([]bool, n+1) // elements default to false
	for i := 2; i <= n; i++ {
		if notPrimeMask[i] {
			continue
		}
		for j := i * i; j <= n; j += i {
			notPrimeMask[j] = i != j
		}
	}
	return &notPrimeMask
}

// Biggest prime factor of n.
// Proj Euler prob 3.
func biggestPrimeFactor(n int) int {
	var bpf int

	floorSqrt := int(math.Floor(math.Sqrt(float64(n))))
	for val, isNotPrime := range *sieve(floorSqrt) {
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
func maxPalinProduct2NDigitInts(nDigits int) (int, error) {
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
	return absInt(a*b) / gcd(a, b)
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
	// The sums of squares of sequence of integers are 'pyramid' numbers,
	// for which there a 0(1) way of calculating the nth:
	sumSqs := ((2 * n * n * n) + (3 * n * n) + n) / 6
	return (sumseq * sumseq) - sumSqs
}

// Nth prime number.
// Find by iteratively apply Sieve of Eratosthenes over increasingly larger ranges.
// Proj Euler prob 7.
func nthPrime(n int) int {
	ub := 2 * n // upper bound of sieving
	for {
		nPrimes := 0
		for val, isNotPrime := range *sieve(ub) {
			if val >= 2 && !isNotPrime {
				nPrimes++
				if nPrimes == n {
					return val
				}
			}
		}
		ub += n
	}
}

var prob8Runes = []rune(
	"73167176531330624919225119674426574742355349194934" +
		"96983520312774506326239578318016984801869478851843" +
		"85861560789112949495459501737958331952853208805511" +
		"12540698747158523863050715693290963295227443043557" +
		"66896648950445244523161731856403098711121722383113" +
		"62229893423380308135336276614282806444486645238749" +
		"30358907296290491560440772390713810515859307960866" +
		"70172427121883998797908792274921901699720888093776" +
		"65727333001053367881220235421809751254540594752243" +
		"52584907711670556013604839586446706324415722155397" +
		"53697817977846174064955149290862569321978468622482" +
		"83972241375657056057490261407972968652414535100474" +
		"82166370484403199890008895243450658541227588666881" +
		"16427171479924442928230863465674813919123162824586" +
		"17866458359124566529476545682848912883142607690042" +
		"24219022671055626321111109370544217506941658960408" +
		"07198403850962455444362981230987879927244284909188" +
		"84580156166097919133875499200524063689912560717606" +
		"05886116467109405077541002256983155200055935729725" +
		"71636269561882670428252483600823257530420752963450")

// Find largest product of a window of n digits in prob8runes.
func prob8(winSize int) int {
	maxProduct := 0
	for i := 0; i < len(prob8Runes)-winSize; i++ {
		winProduct := 1
		for j := i; j < i+winSize; j++ {
			// Quick and dirty way of converting rune to corresponding integer
			winProduct *= int(prob8Runes[j] - '0')
			if winProduct > maxProduct {
				maxProduct = winProduct
			}
		}
	}
	return maxProduct
}

// Find a Pythagorean triplet where sum equals n.
// Proj Euler prob 9.
func pythagoreanTriplet(n int) (int, int, int, error) {
	for a := 1; a < n; a++ {
		for b := a + 1; b < n; b++ {
			c := n - a - b
			if a*a+b*b == c*c {
				return a, b, c, nil
			}
		}
	}
	return 0, 0, 0, errors.New("no satisfactory Pythagorean triplet found")
}

// Sum of all primes below n.
func sumPrimesBelowBound(n int) int {
	var sum int
	for val, isNotPrime := range *sieve(n) {
		if val >= 2 && !isNotPrime {
			sum += val
		}
	}
	return sum
}

var prob11GridRows = []string{
	"08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08",
	"49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00",
	"81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65",
	"52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91",
	"22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80",
	"24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50",
	"32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70",
	"67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21",
	"24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72",
	"21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95",
	"78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92",
	"16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57",
	"86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58",
	"19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40",
	"04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66",
	"88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69",
	"04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36",
	"20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16",
	"20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54",
	"01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48"}

// Find greatest product of winSize adjacent numbers in the same direction (up,
// down, left, right, or diagonally) in the prob11GridRows grid.
// Proj Euler prob 11.
func prob11(winSize int) int {
	ncols := len(strings.Split(prob11GridRows[0], " "))
	nrows := len(prob11GridRows)

	// Convert grid from slice of strings (one per row) to
	// 2D slice of integers (slice of slices)
	grid := make([][]int, ncols)
	for c := range grid {
		grid[c] = make([]int, nrows)
	}
	for y, row := range prob11GridRows {
		for x, val := range strings.Split(row, " ") {
			grid[x][y], _ = strconv.Atoi(val)
		}
	}

	var max int
	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if c < ncols-winSize {
				// Find product for runs of length winSize along rows
				prod := 1
				for i := 0; i < winSize; i++ {
					prod *= grid[r][c+i]
				}
				if prod > max {
					max = prod
				}
			}
			if r < nrows-winSize {
				// Find product for runs of length winSize down columns
				prod := 1
				for i := 0; i < winSize; i++ {
					prod *= grid[r+i][c]
				}
				if prod > max {
					max = prod
				}
			}
			if r < nrows-winSize && c < ncols-winSize {
				// Find product for diagonal runs of length winSize
				prod := 1
				for i := 0; i < winSize; i++ {
					prod *= grid[r+i][c+i]
				}
				if prod > max {
					max = prod
				}
			}
			if r > winSize - 1 && c < ncols-winSize {
				// Find product for runs of length winSize along other diagonal
				prod := 1
				for i := 0; i < winSize; i++ {
					prod *= grid[r-i][c+i]
				}
				if prod > max {
					max = prod
				}
			}
		}
	}
	return max
}


func main() {
}
