package main

import "fmt"

func main() {
	arr := []int{2, 3, 9, 10, 20, 1}
	shellSort(arr)
	fmt.Println(arr)
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
	for i := 0; i < len(arr)-1; i++ {
		mid := i
		for j := i + 1; j < len(arr); j++ {
			if arr[mid] > arr[j] {
				arr[mid], arr[j] = arr[j], arr[mid]
			}
		}
	}
}

func insertion(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		var value, index = arr[i+1], i
		for index >= 0 && value < arr[i] {
			arr[index+1] = arr[index]
			index--
		}
		arr[index+1] = value
	}
}

func shell(arr []int) {
	// 算出每次分组的元素个数和步数
	for pace := len(arr) / 2; pace > 0; pace /= 2 {
		// i = index 是下标 每次都在移动下标然后减去步数间隔
		for i := pace; i < len(arr); i++ {
			for j := i - pace; j >= 0; j -= pace {
				if arr[j] > arr[j+pace] {
					arr[j], arr[j+pace] = arr[j+pace], arr[j]
				}
			}
		}
	}
}

func shellSort(row []int) {
	var increment int = 1
	// 取第一次的分组 步长
	for increment < len(row)/2 {
		increment = increment*2 + 1
	}
	for increment >= 1 {
		for i := increment; i < len(row); i++ {
			for j := i; j >= increment; j-- {
				if row[j] < row[j-increment] {
					row[j], row[j-increment] = row[j-increment], row[j]
				}
			}
		}
		increment /= 2
	}
}
