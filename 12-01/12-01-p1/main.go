package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// line = "31594   93577"
		fields := strings.Fields(line) // like strings.Split() but auto-removes whitespace
		// fields = "[31594 93577]"
		if len(fields) != 2 {
			continue // skips anything with != 2 fields
		}

		left, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Printf("Error parsing left number: %v\n", left)
			continue
		}

		right, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("Error parsing right number: %v\n", right)
			continue
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDist := 0
	for i := 0; i < len(leftList); i++ {
		dist := abs(leftList[i] - rightList[i])
		totalDist += dist
	}

	fmt.Printf("Total distance: %d\n", totalDist)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
