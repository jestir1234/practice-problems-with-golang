package main

import (
	"fmt"
	// "reflect"
)

// Find the most frequent integer in an array (fixed array length)
func mostFrequentInt(arr [6]int) int {
	var nums = make(map[int]int)
	var largest int = arr[0]
	largestCount := 0
	for i := 0; i < len(arr); i++ {
		nums[arr[i]] += 1
		if nums[arr[i]] > largestCount {
			largest = arr[i]
			largestCount = nums[arr[i]]
		}
	}

	fmt.Println(largest)
	return largest
}

// O(n) time complexity
// O(n) space complexity

// Find the most frequent integer in an array/slice (varying length)
func mostFrequentIntSlice(arr []int) int {
	var nums = make(map[int]int)
	var largest int = arr[0]
	largestCount := 0

	for i := 0; i < len(arr); i++ {
		nums[arr[i]] += 1
		if nums[arr[i]] > largestCount {
			largestCount = nums[arr[i]]
			largest = arr[i]
		}
	}
	fmt.Println(largest)
	return largest
}

// Find pairs in an integer array whose sum is equal to 10 (bonus: do it in linear time)

func sumTo10(arr []int) [2]int {
	var nums = make(map[int]int)
	var result [2]int
	for i := 0; i < len(arr); i++ {
		if nums[10-arr[i]] == arr[i] {
			fmt.Println("FOUND IT")
			result = [2]int{nums[10-arr[i]], arr[i]}
		} else {
			nums[10-arr[i]] = arr[i]
		}
	}
	return result
}

// Find the first non-repeated character in a String

func firstNonRepeatCharacter(sampleString string) string {

	var stringHash = make(map[string][2]int)
	for i := 0; i < len(sampleString); i++ {
		var currentChar = string(sampleString[i])
		if _, char := stringHash[currentChar]; char {
			stringHash[currentChar] = [2]int{2, i}
		} else {
			stringHash[currentChar] = [2]int{1, i}
		}
	}

	lowest := len(sampleString)
	resultChar := ""

	fmt.Println(lowest)
	fmt.Println(stringHash)
	for k, v := range stringHash {
		if v[0] == 1 && v[1] < lowest {
			lowest = v[1]
			resultChar = k
		}
	}

	fmt.Println(resultChar, lowest)

	return resultChar
}

func main() {
	// numbers := []int{1, 2, 2, 5, 5, 5}
	// mostFrequentInt(numbers)
	// sliceNumbers := make([]int, 10)
	// sliceNumbers = []int{6, 65, 47, 65, 99}
	// mostFrequentIntSlice(sliceNumbers)
	// sumTo10(numbers)
	var sample string = "disaster"
	firstNonRepeatCharacter(sample)
}
