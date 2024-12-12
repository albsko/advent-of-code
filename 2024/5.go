package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func y2024_5() int {
	result := 0

	f, _ := os.Open("5_1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rules := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "|")

		l, _ := strconv.Atoi(s[0])
		r, _ := strconv.Atoi(s[1])

		rules[l] = append(rules[l], r)
	}

	f, _ = os.Open("5_2.txt")
	defer f.Close()

	scanner = bufio.NewScanner(f)

	data := make([][]int, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")

		sInt := make([]int, len(s))
		for i := 0; i < len(s); i++ {
			el, _ := strconv.Atoi(s[i])
			sInt[i] = el
		}

		data = append(data, sInt)
	}

	results := make([][]int, 0, 1024)

	for _, arr := range data {
		if validateOrder(arr, rules) {
			results = append(results, arr)
		}
	}

	for _, arr := range results {
		result += arr[len(arr)/2]
	}

	return result
}

func validateOrder(data []int, rules map[int][]int) bool {
	positions := make(map[int]int, len(data))
	for i, page := range data {
		positions[page] = i
	}

	for i, page := range data {
		for beforePage, afterPages := range rules {
			for _, afterPage := range afterPages {
				if afterPage == page {
					if pos, ok := positions[beforePage]; ok {
						if pos >= i {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
