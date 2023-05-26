package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}

type Number interface {
	int64 | float64
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as type for map values
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func SumNums[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func main() {

	// initialize a map for int values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// initialize a map for float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("Generic Sums with Constraints: %v and %v\n", SumNums(ints), SumNums(floats))
}
