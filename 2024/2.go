package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func y2024_2() int {
	f, _ := os.Open("2.txt")
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Fields(line)

		x := make([]int, 0, 6)
		for _, v := range s {
			vint, _ := strconv.Atoi(v)
			x = append(x, vint)
		}

		if isChangingInSameDirection(x, 3) {
			count++
		}
	}

	return count
}

func isChangingInSameDirection(s []int, maxChange int) bool {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if len(s) <= 1 {
		return false
	}

	initDiff := s[1] - s[0]
	if initDiff == 0 {
		return false
	}

	isIncreasing := initDiff > 0

	for i := 1; i < len(s); i++ {
		diff := s[i] - s[i-1]

		if abs(diff) > maxChange {
			return false
		}

		if isIncreasing && diff <= 0 {
			return false
		}
		if !isIncreasing && diff >= 0 {
			return false
		}
	}

	return true
}
