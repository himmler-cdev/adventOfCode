package src

import (
	"strconv"
	"strings"
)

func Level2Task1() int {
	data, err := ReadFileAsLines("level2.txt")
	saveAmount := 0

	if err != nil {
		return -1
	}

	for _, entry := range data {
		nums := strings.Split(entry, " ")
		intList := make([]int, len(nums))
		for i, v := range nums {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			intList[i] = num
		}

		if validate(intList) {
			saveAmount++
		}
	}

	return saveAmount
}

func validate(intList []int) bool {
	desc, incr, max, min := true, true, true, true

	for i := 1; i < len(intList); i++ {
		diff := absValue(intList[i] - intList[i-1])
		if intList[i] > intList[i-1] {
			desc = false
		}
		if intList[i] < intList[i-1] {
			incr = false
		}
		if diff > 3 {
			max = false
		}
		if diff < 1 {
			min = false
		}
		if !(desc || incr) && (!max || !min) {
			return false
		}
	}

	return (desc || incr) && max && min
}

func absValue(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Level2Task2() int {
	data, err := ReadFileAsLines("level2.txt")
	saveAmount := 0

	if err != nil {
		return -1
	}

	for _, entry := range data {
		nums := strings.Split(entry, " ")
		intList := make([]int, len(nums))
		for i, v := range nums {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			intList[i] = num
		}

		if validateTask2(intList) {
			saveAmount++
		}
	}

	return saveAmount
}

func validateTask2(intList []int) bool {
	for i := 0; i < len(intList); i++ {
		data1 := append([]int{}, intList[:i]...)
		data1 = append(data1, intList[i+1:]...)
		if validate(data1) {
			return true
		}
	}
	return false
}
