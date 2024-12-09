package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func y2024_3() int {
	f, _ := os.Open("3.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	text := scanner.Text()
	if scanner.Scan() {
		text = scanner.Text()
	}

	pattern := `mul\(([0-9]+),([0-9]+)\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(text, -1)

	sum := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}

	return sum
}
