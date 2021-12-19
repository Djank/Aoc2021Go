package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileToString(filePath string) string {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(fileData)
}

func InputToStrings(str string) []string {
	return strings.Split(strings.TrimSuffix(str, "\n"), "\n")
}

func InputToIntegers(str string) []int {
	result := []int{}
	for lineNum, s := range InputToStrings(str) {
		s = strings.Trim(s, "\r\n")
		number, err := strconv.Atoi(s)
		if err != nil {
			log.Printf("Input error on line %d", lineNum)
			log.Panic(err)
		}
		result = append(result, number)
	}
	return result
}

func SplitLineToIntegers(line string, sep string) []int {
	ints := strings.Split(line, sep)
	var numbers []int
	for _, v := range ints {
		v = strings.Trim(v, " ")
		number, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("Can't convert string %s to int. Line %s", v, line)
			log.Panic(err)
		}
		numbers = append(numbers, number)	
	}
	return numbers
}
