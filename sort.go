package sort

import (
	"math/rand"
)

func bubbleSort(arrayIn []int) {
	for i := 1; i <= len(arrayIn); i++ {
		for j := 1; j < len(arrayIn)+1-i; j++ {
			if arrayIn[j-1] > arrayIn[j] {
				arrayIn[j-1], arrayIn[j] = arrayIn[j], arrayIn[j-1]
			}
		}
	}
}

func selectionSort(arrayIn []int) {
	for i := 0; i < len(arrayIn); i++ {
		min_i, _ := min(arrayIn[i:])
		arrayIn[i], arrayIn[min_i+i] = arrayIn[min_i+i], arrayIn[i]
	}
}

func min(arrayIn []int) (index int, min int) {
	index = 0
	min = arrayIn[0]
	for i := 1; i < len(arrayIn); i++ {
		if arrayIn[i] < min {
			min = arrayIn[i]
			index = i
		}
	}
	return
}

func insertionSort(arrayIn []int) {
	for i := 1; i < len(arrayIn); i++ {
		for j := i; j > 0; j-- {
			if arrayIn[j] < arrayIn[j-1] {
				arrayIn[j], arrayIn[j-1] = arrayIn[j-1], arrayIn[j]
			}
		}
	}
}

func mergeSort(arrayIn []int) []int {
	if len(arrayIn) == 1 {
		return arrayIn
	} else if len(arrayIn) > 1 {
		arrayLeft, arrayRight := splitArray(arrayIn)
		return merge(mergeSort(arrayLeft), mergeSort(arrayRight))
	}
	return nil
}

func splitArray(arrayIn []int) ([]int, []int) {
	return arrayIn[:len(arrayIn)/2], arrayIn[len(arrayIn)/2:]
}

func merge(arrayLeft, arrayRight []int) (arrayOut []int) {
	var el int
	for len(arrayLeft) > 0 && len(arrayRight) > 0 {
		if arrayLeft[0] <= arrayRight[0] {
			el, arrayLeft = arrayLeft[0], arrayLeft[1:]
		} else if arrayRight[0] < arrayLeft[0] {
			el, arrayRight = arrayRight[0], arrayRight[1:]
		}
		arrayOut = append(arrayOut, el)
	}
	arrayOut = append(arrayOut, arrayLeft...)
	arrayOut = append(arrayOut, arrayRight...)
	return
}

func quickSort(arrayIn []int) (arrayOut []int) {
	if len(arrayIn) == 1 {
		arrayOut = arrayIn
		return
	} else if len(arrayIn) > 1 {
		p := rand.Intn(len(arrayIn))
		pivot := arrayIn[p]
		left := 0
		arrayIn[len(arrayIn)-1], arrayIn[p] = arrayIn[p], arrayIn[len(arrayIn)-1]
		for i := 0; i < len(arrayIn); i++ {
			if arrayIn[i] < pivot {
				arrayIn[i], arrayIn[left] = arrayIn[left], arrayIn[i]
				left++
			}
		}
		arrayIn[left], arrayIn[len(arrayIn)-1] = arrayIn[len(arrayIn)-1], arrayIn[left]
		arrayOut = append(arrayOut, quickSort(arrayIn[:p])...)
		arrayOut = append(arrayOut, quickSort(arrayIn[p:])...)
	}
	return
}
