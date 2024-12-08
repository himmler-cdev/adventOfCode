package src

import (
	"regexp"
	"strconv"
	"strings"
)

func Level3Task1() int {
	data, err := ReadFileAsLines("level3.txt")
	sum := 0

	if err != nil {
		return -1
	}

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, entry := range data {
		matches := re.FindAllStringSubmatch(entry, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	return sum
}

func Level3Task2() int {
	data, err := ReadFileAsLines("level3.txt")
	stringFile := ""
	sum := 0

	if err != nil {
		return -1
	}

	for _, entry := range data {
		stringFile += entry
	}

	entry := removeDoDont(stringFile)

	for _, mul := range findMul(entry) {
		num1, _ := strconv.Atoi(mul[1])
		num2, _ := strconv.Atoi(mul[2])
		sum += num1 * num2
	}

	return sum
}

func findMul(line string) [][]string {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	return re.FindAllStringSubmatch(line, len(line))
}

func removeDoDont(line string) string {
	result := ""
	enabled := true
	dontOffset := len("don't()")
	doOffset := len("do()")

	for len(line) > 0 {
		dontIndex := strings.Index(line, "don't()")
		doIndex := strings.Index(line, "do()")

		if dontIndex == -1 && doIndex == -1 {
			if enabled {
				result += line
			}
			break
		}

		if dontIndex != -1 && (doIndex == -1 || dontIndex < doIndex) {
			if enabled {
				result += line[:dontIndex]
			}
			enabled = false
			line = line[dontIndex+dontOffset:]
		} else {
			if enabled {
				result += line[:doIndex]
			}
			enabled = true
			line = line[doIndex+doOffset:]
		}
	}

	return result
}
