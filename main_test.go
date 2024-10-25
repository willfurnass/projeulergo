package main

import (
	"os"
	"testing"
)

func TestProb1(t *testing.T) {
	got := sumDistinctMultsLinear(3, 5, 10-1)
	want := 23
	if got != want {
		t.Errorf("sumDistinctMultsLinear(3, 5, 10-1) = %d; want %d", got, want)
	}

	got = sumDistinctMultsLinear(3, 5, 1000-1)
	want = 233168
	if got != want {
		t.Errorf("sumDistinctMultsLinear(3, 5, 1000-1) = %d; want %d", got, want)
	}
}

func TestSumSeqConst(t *testing.T) {
	got := sumSeqConst(3, 10)
	want := 3 + 6 + 9
	if got != want {
		t.Errorf("sumSeqConst(3, 10-1) = %d; want %d", got, want)
	}
}
func TestSumDistinctMultsConst(t *testing.T) {
	got := sumDistinctMultsConst(3, 5, 10-1)
	want := 3 + 6 + 9 + 5
	if got != want {
		t.Errorf("sumDistinctMultsConst(3, 5, 10-1) = %d; want %d", got, want)
	}

	got = sumDistinctMultsConst(3, 5, 1000-1)
	want = 233168
	if got != want {
		t.Errorf("sumDistinctMultsConst(3, 5, 1000-1) = %d; want %d", got, want)
	}
}

func TestProb2(t *testing.T) {
	got := sumEvenFibs(90)
	want := 44
	if got != want {
		t.Errorf("sumEvenFibs(90) = %d; want %d", got, want)
	}

	got = sumEvenFibs(4000000)
	want = 4613732
	if got != want {
		t.Errorf("sumEvenFibs(4000000) = %d; want %d", got, want)
	}
}

func TestProb3(t *testing.T) {
	got := biggestPrimeFactor(13195)
	want := 29
	if got != want {
		t.Errorf("biggestPrimeFactor(13195) = %d; want %d", got, want)
	}

	got = biggestPrimeFactor(600851475143)
	want = 6857
	if got != want {
		t.Errorf("biggestPrimeFactor(600851475143) = %d; want %d", got, want)
	}
}

func TestIsPalindrome(t *testing.T) {
	got := isPalindrome(9009)
	want := true
	if got != want {
		t.Errorf("isPalindrome(9009) = %t; want %t", got, want)
	}
	got = isPalindrome(9008)
	want = false
	if got != want {
		t.Errorf("isPalindrome(9008) = %t; want %t", got, want)
	}
}
func TestProb4(t *testing.T) {
	got, err := maxPalinProduct2NDigitInts(2)
	want := 9009
	if err != nil {
		t.Errorf("Prob4(2) errored: %v; want %d", err, want)
	}
	if got != want {
		t.Errorf("Prob4(2) = %d; want %d", got, want)
	}

	got, err = maxPalinProduct2NDigitInts(3)
	want = 906609
	if err != nil {
		t.Errorf("Prob4(3) errored: %v; want %d", err, want)
	}
	if got != want {
		t.Errorf("Prob4(3) = %d; want %d", got, want)
	}
}

func TestGCD(t *testing.T) {
	got := gcd(1071, 462)
	want := 21
	if got != want {
		t.Errorf("gcdEuclid(1071, 462) = %d; want %d", got, want)
	}
}
func TestLCMRange(t *testing.T) {
	got := lcmRange(10)
	want := 2520
	if got != want {
		t.Errorf("lcmRange(10) = %d; want %d", got, want)
	}
	got = lcmRange(20)
	want = 232792560
	if got != want {
		t.Errorf("lcmRange(20) = %d; want %d", got, want)
	}
}

func TestDiffSumSqsSqSums(t *testing.T) {
	got := diffSumSqsSqSums(10)
	want := 2640
	if got != want {
		t.Errorf("diffSumSqsSqSums(10) = %d; want %d", got, want)
	}
	got = diffSumSqsSqSums(100)
	want = 25164150
	if got != want {
		t.Errorf("diffSumSqsSqSums(100) = %d; want %d", got, want)
	}
}

func TestNthPrime(t *testing.T) {
	got := nthPrime(6)
	want := 13
	if got != want {
		t.Errorf("nthPrime(6) = %d; want %d", got, want)
	}
	got = nthPrime(10001)
	want = 104743
	if got != want {
		t.Errorf("nthPrime(10001) = %d; want %d", got, want)
	}
}

func TestProb8(t *testing.T) {
	got := prob8(4)
	want := 5832
	if got != want {
		t.Errorf("prob8(4) = %d; want %d", got, want)
	}
	got = prob8(13)
	want = 23514624000
	if got != want {
		t.Errorf("prob8(13) = %d; want %d", got, want)
	}
}

func TestPythagoreanTriplet(t *testing.T) {
	a, b, c, err := pythagoreanTriplet(1000)
	if err != nil {
		t.Errorf("pythagoreanTriplet(1000) unexpectedly errored: %v", err)
	}
	got := a * b * c
	want := 31875000
	if got != want {
		t.Errorf("pythagoreanTriplet(1000) = %d; want %d", got, want)
	}
}

func TestSumPrimesBelowBound(t *testing.T) {
	got := sumPrimesBelowBound(10)
	want := 17
	if got != want {
		t.Errorf("sumPrimesBelowBound(4) = %d; want %d", got, want)
	}
	got = sumPrimesBelowBound(2000000)
	want = 142913828922
	if got != want {
		t.Errorf("sumPrimesBelowBound(2000000) = %d; want %d", got, want)
	}
}

func TestProb11(t *testing.T) {
	got := prob11(4)
	want := 70600674
	if got != want {
		t.Errorf("prob11(4) = %d; want %d", got, want)
	}
}

func TestMinTriNumExcessOfNDivisors(t *testing.T) {
	got := minTriNumExcessOfNDivisors(5)
	want := 28
	if got != want {
		t.Errorf("minTriNumExcessOfNDivisors(5) = %d; want %d", got, want)
	}
	got = minTriNumExcessOfNDivisors(500)
	want = 76576500
	if got != want {
		t.Errorf("minTriNumExcessOfNDivisors(500) = %d; want %d", got, want)
	}
}

func TestProb13BigInt(t *testing.T) {
	got, err := prob13BigInt()
	var want uint64 = 5537376230
	if err != nil {
		t.Errorf("prob13BigInt() unexpectedly errored: %v", err)
	}
	if got != want {
		t.Errorf("prob13BigInt() = %d; want %d", got, want)
	}
}

func TestProb13CustomArbPrec(t *testing.T) {
	got, err := prob13CustomArbPrec()
	var want uint64 = 5537376230
	if err != nil {
		t.Errorf("prob13CustomArbPrec() unexpectedly errored: %v", err)
	}
	if got != want {
		t.Errorf("prob13CustomArbPrec() = %d; want %d", got, want)
	}
}

func TestCollatzSeqLen(t *testing.T) {
	got := collatzSeqLen(13)
	var want uint = 10
	if got != want {
		t.Errorf("collatzSeqLen(13) = %d; want %d", got, want)
	}
}

func TestLongestCollatzSeqLen(t *testing.T) {
	got := longestCollatzSeqLen(1000000)
	var want uint = 837799
	if got != want {
		t.Errorf("LongestCollatzSeqLen(1000000) = %d; want %d", got, want)
	}
}
func TestProb15(t *testing.T) {
	got := prob15(20)
	var want uint64 = 137846528820
	if got != want {
		t.Errorf("prob15(20) = %d; want %d", got, want)
	}
}

func TestDecDigitSum2PowerN(t *testing.T) {
	got := decDigitSum2PowerN(15)
	var want uint = 26
	if got != want {
		t.Errorf("decDigitSum2PowerN(15) = %d; want %d", got, want)
	}

	got = decDigitSum2PowerN(1000)
	want = 1366
	if got != want {
		t.Errorf("decDigitSum2PowerN(1000) = %d; want %d", got, want)
	}
}

func TestProb17(t *testing.T) {
	got := prob17(1, 5)
	var want int = 19
	if got != want {
		t.Errorf("prob17(1, 5) = %d; want %d", got, want)
	}

	got = prob17(1, 1000)
	want = 21124
	if got != want {
		t.Errorf("prob17(1, 1000) = %d; want %d", got, want)
	}
}

func TestProb19(t *testing.T) {
	got := prob19()
	var want int = 171
	if got != want {
		t.Errorf("prob19() = %d; want %d", got, want)
	}
}

func TestProb20(t *testing.T) {
	got := prob20()
	var want int = 648
	if got != want {
		t.Errorf("prob20() = %d; want %d", got, want)
	}
}

func TestProb22(t *testing.T) {
	dat, err := os.ReadFile("0022_names.txt")
	if err != nil {
		panic(err)
	}
	got := sumNameScores(string(dat))
	var want int = 871198282
	if got != want {
		t.Errorf("sumNameScores(someNamesStr) = %d; want %d", got, want)
	}
}
