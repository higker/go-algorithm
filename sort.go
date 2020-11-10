package main

import (
	"math/rand"
	"time"
)

func main() {

	arr := generateNumbers(10)
	//fmt.Println(arr)
	mergeSort(arr) // 0.752 seconds

}

func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func selection(arr []int) {
	var min int
	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}

func insertion(arr []int) {
	var pervIndex, current int
	for i := 1; i < len(arr); i++ {
		current = arr[i]
		pervIndex = i - 1
		for pervIndex >= 0 && current < arr[pervIndex] {
			arr[pervIndex+1] = arr[pervIndex]
			pervIndex--
		}
		arr[pervIndex+1] = current
	}
}

func shell(arr []int) {
	var pervIndex, current int
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			current = arr[i]
			pervIndex = i - gap
			for pervIndex >= 0 && current < arr[pervIndex] {
				arr[pervIndex+gap] = arr[pervIndex]
				pervIndex -= gap
			}
			arr[pervIndex+gap] = current
		}
	}
}

// https://zh.wikipedia.org/wiki/%E9%9A%8E%E4%B9%98
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// func mergeSort(arr []int) []int {
// 	var result []int
// 	if len(arr) < 2 {
// 		return arr
// 	}
// 	middle := len(arr) / 2
// 	// 注意这是切片 切取的时候是包左 不包右
// 	left := arr[:middle]
// 	right := arr[middle:]
// 	return func(left, right []int) []int {
// 		// index
// 		l, r := 0, 0
// 		for l < len(left) && r < len(right) {
// 			if left[l] < right[r] {
// 				result = append(result, left[l])
// 				l++
// 			} else {
// 				result = append(result, right[r])
// 				r++
// 			}
// 		}
// 		result = append(result, left[l:]...)
// 		result = append(result, right[r:]...)
// 		return result
// 	}(mergeSort(left), mergeSort(right))
// }

func mergeSort(arr []int) []int {
	var result []int
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	// 注意这是切片 切取的时候是包左 不包右
	left := arr[:middle]
	right := arr[middle:]
	return func(left, right []int) []int {
		// 分组不能保证左右各组数据个数是一样的
		for len(left) != 0 && len(right) != 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		}
		if len(left) != 0 {
			result = append(result, left[:]...)

		}
		if len(right) != 0 {
			result = append(result, right[:]...)

		}
		return result
	}(mergeSort(left), mergeSort(right))
}

func generateNumbers(size int) []int {
	numbers := make([]int, 0, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		numbers = append(numbers, (rand.Intn(99999) - 9999))
	}
	return numbers
}
