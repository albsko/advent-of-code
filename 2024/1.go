package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func y2024_1() int {
	f, _ := os.Open("1.txt")
	defer f.Close()

	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")

		l, _ := strconv.Atoi(s[0])
		r, _ := strconv.Atoi(s[1])

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	result := 0
	for i := range left {
		result = result + abs(left[i]-right[i])
	}

	return result
}
