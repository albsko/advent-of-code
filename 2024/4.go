package main

import (
	"bufio"
	"os"
)

func y2024_4() int {
	f, _ := os.Open("4.txt")
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)

	m := make([][]rune, 0, 1024)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))

		for i, r := range line {
			row[i] = r
		}

		m = append(m, row)
	}

	word := "XMAS"

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	r, c := len(m), len(m[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if searchWord(m, word, i, j, dx, dy) {
					count++
				}
			}
		}
	}

	return count
}

func searchWord(matrix [][]rune, word string, startX, startY, dx, dy int) bool {
	rows, cols := len(matrix), len(matrix[0])

	for k := 0; k < len(word); k++ {
		x, y := startX+k*dx, startY+k*dy

		if x < 0 || x >= rows || y < 0 || y >= cols || matrix[x][y] != rune(word[k]) {
			return false
		}
	}
	return true
}
