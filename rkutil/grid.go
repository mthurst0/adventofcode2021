package rkutil

import "strconv"

func ReadGrid(filename string) [][]int {
	lines := MustLines(filename)
	result := make([][]int, len(lines))
	for j, line := range lines {
		for i := 0; i < len(line); i++ {
			c, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			result[j] = append(result[j], c)
		}
	}
	return result
}
