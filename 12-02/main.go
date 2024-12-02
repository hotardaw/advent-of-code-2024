package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safeReportsCounter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()         // "1 3 4 5 8 10 7"
		fields := strings.Fields(line) // "[1 3 4 5 8 10 7]"

		isSafe := checkIfSafe(fields)
		if isSafe {
			safeReportsCounter += 1
		}
	}
	fmt.Println(safeReportsCounter)
}

/*
reports = rows (6)
levels = columns (5)
7 6 4 2 1 		// safe
1 2 7 8 9 		// unsafe
9 7 6 2 1 		// unsafe
*/

func checkIfSafe(report []string) bool {
	var direction int // 1 for increasing, -1 for decreasing

	for i := range len(report) {
		if i > 0 {
			curr, _ := strconv.Atoi(report[i])
			prev, _ := strconv.Atoi(report[i-1])

			diff := curr - prev

			// First diff observed decides the direction
			if i == 1 {
				if diff > 0 {
					direction = 1
				} else {
					direction = -1
				}
			}

			// Check diff magnitude
			absDiff := abs(diff)
			if absDiff < 1 || absDiff > 3 {
				return false
			}

			// Check direction continuation
			if (direction == 1 && diff <= 0) || (direction == -1 && diff >= 0) {
				return false
			}
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
