package main

import "fmt"

// Find the matching element in the two arrays
func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{5, 6, 7, 8, 9}

	solution1(arr1, arr2)

	solution2(arr1, arr2)
}

// The simple solution of iterating over the two maps creates a runtime of
// O(n^2)
func solution1(arr1, arr2 []int) {
	for _, i := range arr1 {
		for _, j := range arr2 {
			if i == j {
				fmt.Println("match:", j)
			}
		}
	}
}

// Using a map allows us to reach an O(n) runtime
func solution2(arr1, arr2 []int) {
	arr1map := make(map[int]struct{})
	for _, i := range arr1 {
		arr1map[i] = struct{}{}
	}
	for _, j := range arr2 {
		_, ok := arr1map[j]
		if ok {
			fmt.Println("match:", j)
		}
	}
}
