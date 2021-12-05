package rkutil

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readLines(filename string, conv func(string) string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, conv(scanner.Text()))
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return result, err
}

func ReadLines(filename string) ([]string, error) {
	return readLines(filename, func(s string) string { return s })
}

func MustLines(filename string) []string {
	result, err := ReadLines(filename)
	if err != nil {
		panic(err)
	}
	return result
}

func ReadTrimmedLines(filename string) ([]string, error) {
	return readLines(filename, func(s string) string { return strings.TrimSpace(s) })
}

func MustTrimmedLines(filename string) []string {
	result, err := ReadTrimmedLines(filename)
	if err != nil {
		panic(err)
	}
	return result
}

func ReadLinesOfNumbers(filename string) ([]int, error) {
	lines, err := ReadTrimmedLines(filename)
	if err != nil {
		return nil, err
	}
	result := make([]int, 0, len(lines))
	for _, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func MustLinesOfNumbers(filename string) []int {
	result, err := ReadLinesOfNumbers(filename)
	if err != nil {
		panic(err)
	}
	return result
}
