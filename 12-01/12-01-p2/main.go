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

// figure out how often each num in left list appears in right list
// calc the similarity score by adding each num in the left list after
// multiplying it by the num of times it appears in the right list.

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

	// Count occurrences in the right list
	rightCountMap := make(map[int]int)
	for _, num := range rightList {
		rightCountMap[num]++
	}

	// Calculate similarity score
	similarityScore := 0
	for _, leftNum := range leftList {
		similarityScore += leftNum * rightCountMap[leftNum]
	}

	fmt.Println("Similarity score: ", similarityScore)
}

/*
Nil maps - will panic if written to:
	var m map[string]int
	m := map[string]int(nil)

Usable maps:
	m := make(map[string]int)
	m := map[string]int{} // empty map literal
	m := map[string]int{"a": 1} // map with init values
*/
