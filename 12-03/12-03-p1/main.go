package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mulPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	mulRegex := regexp.MustCompile(mulPattern)

	var totalSum int64

	for scanner.Scan() {
		chunk := scanner.Text()

		matches := mulRegex.FindAllString(chunk, -1)
		for _, match := range matches {
			submatch := mulRegex.FindStringSubmatch(match)
			if len(submatch) == 3 { // Full match & two capture groups
				x, _ := strconv.Atoi(submatch[1])
				y, _ := strconv.Atoi(submatch[2])
				product := x * y
				totalSum += int64(product)
			}
		}
	}

	fmt.Println(totalSum)
}
