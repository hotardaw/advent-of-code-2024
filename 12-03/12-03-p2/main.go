package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var result int
	do := true

	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		pattern := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
		instructions := pattern.FindAllString(scanner.Text(), -1) // []string
		for _, instruction := range instructions {
			if strings.Compare("do()", instruction) == 0 {
				do = true
			} else if strings.Compare("don't()", instruction) == 0 {
				do = false
			} else {
				if do {
					result += extractAndMultiply(instruction)
				}
			}
		}
	}

	fmt.Println(result)
}

func extractAndMultiply(mul string) int {
	pattern := regexp.MustCompile(`\d+`)
	var numbers []int
	for _, s := range pattern.FindAllString(mul, -1) {
		num, _ := strconv.Atoi(s)

		numbers = append(numbers, num)
	}
	return numbers[0] * numbers[1]
}
