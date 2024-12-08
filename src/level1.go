package src

import (
	"sort"
	"strconv"
	"strings"
)

func Level1() int {
	data, err := ReadFileAsLines("level1.txt")
	leftList := []int{}
	rightList := []int{}
	sum := 0

	if err != nil {
		return -1
	}

	for _, line := range data {
		leftAndRight := strings.Split(line, "   ")
		leftInt, _ := strconv.Atoi(leftAndRight[0])
		leftList = append(leftList, leftInt)
		rightInt, _ := strconv.Atoi(leftAndRight[1])
		rightList = append(rightList, rightInt)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	for idx, left := range leftList {
		if left < rightList[idx] {
			sum += rightList[idx] - left
		} else {
			sum += left - rightList[idx]
		}
	}

	return sum
}

func Level1Task2() int {
	data, err := ReadFileAsLines("level1.txt")
	leftList := []int{}
	rightList := []int{}
	sum := 0

	if err != nil {
		return -1
	}

	for _, line := range data {
		leftAndRight := strings.Split(line, "   ")
		leftInt, _ := strconv.Atoi(leftAndRight[0])
		leftList = append(leftList, leftInt)
		rightInt, _ := strconv.Atoi(leftAndRight[1])
		rightList = append(rightList, rightInt)
	}

	for _, left := range leftList {
		cnt := 0
		for _, right := range rightList {

			if left == right {
				cnt++
			}
		}
		sum += left * cnt
	}

	return sum
}
