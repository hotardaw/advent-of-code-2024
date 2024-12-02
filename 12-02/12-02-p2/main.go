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
		line := scanner.Text()
		fields := strings.Fields(line)

		nums := make([]int, len(fields))
		for i, field := range fields {
			nums[i], _ = strconv.Atoi(field)
		}

		// Check if safe originally or after removing one number
		if isSequenceSafe(nums) || canBeMadeSafe(nums) {
			safeReportsCounter++
		}
	}
	fmt.Printf("\nTotal safe reports: %d\n", safeReportsCounter)
}

func isSequenceSafe(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	isIncreasing := true
	if nums[1]-nums[0] < 0 {
		isIncreasing = false
	}

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}

		// Check if direction matches initial direction
		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}
	}
	return true
}

func canBeMadeSafe(nums []int) bool {
	// Try removing each number once
	for i := range nums {
		withoutI := make([]int, 0, len(nums)-1)
		withoutI = append(withoutI, nums[:i]...)
		withoutI = append(withoutI, nums[i+1:]...)

		if isSequenceSafe(withoutI) {
			return true
		}
	}
	return false
}
