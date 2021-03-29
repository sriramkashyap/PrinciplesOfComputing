package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Creates an array of given size, with random numbers
// between range [min, max).
func create_array(size int, min int, max int) []int {
	rand.Seed(int64(time.Now().Second()))
	var numbers []int
	for i := 0; i < size; i++ {
		numbers = append(numbers, min+rand.Intn(max-min))
	}
	return numbers
}

// Pretty-prints the array.
func print_array(numbers []int) string {
	var result strings.Builder
	result.WriteByte('[')
	if len(numbers) > 0 {
		fmt.Fprint(&result, numbers[0])
		for i := 1; i < len(numbers); i++ {
			fmt.Fprintf(&result, ",%v", numbers[i])
		}
	}
	result.WriteByte(']')
	return result.String()
}

// Test basic array creation.
func create_array_test() {
	fmt.Printf("CREATE ARRAY TEST:\n")
	fmt.Printf("Array: %s\n", print_array(create_array(0, 0, 9)))
	fmt.Printf("Array: %s\n", print_array(create_array(5, 0, 9)))
	fmt.Printf("Array: %s\n", print_array(create_array(10, 0, 9)))
	fmt.Printf("Array: %s\n", print_array(create_array(10, -10, 9)))
	fmt.Printf("Array: %s\n", print_array(create_array(10, 0, 1)))
	fmt.Printf("\n")
}

// Swap two elements in an array. No validation.
func swap(input []int, index_a int, index_b int) {
	tmp := input[index_a]
	input[index_a] = input[index_b]
	input[index_b] = tmp
}

// Make a copy of an array.
func dup(input []int) []int {
	var retval []int
	for i := 0; i < len(input); i++ {
		retval = append(retval, input[i])
	}
	return retval
}

// Recursive permutation of input array. Returns array of arrays.
func permute_array(input []int, offset int) [][]int {
	var result [][]int
	if offset == len(input) {
		return [][]int{dup(input)}
	}
	for i := offset; i < len(input); i++ {
		swap(input, offset, i)
		result = append(result, permute_array(input, offset+1)...)
		swap(input, offset, i)
	}
	return result
}

// Test permutation.
func permute_test() {
	fmt.Printf("PERMUTE TEST:\n")
	test_arrays := [][]int{{}, {1}, {1, 2}, {1, 2, 3}}
	for _, test_array := range test_arrays {
		fmt.Printf("Input: %s\n", print_array(test_array))
		my_permutations := permute_array(test_array, 0)
		for i := 0; i < len(my_permutations); i++ {
			fmt.Printf(" P[%d]: %s\n", i, print_array(my_permutations[i]))
		}
	}
	fmt.Printf("\n")
}

func main() {
	create_array_test()
	permute_test()
}
