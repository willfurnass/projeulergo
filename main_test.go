package main

import "testing"

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
