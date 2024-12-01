package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var lefts, rights = getLeftsAndRights()
	fmt.Printf("distance: %f\n", getDistance(lefts, rights))
	fmt.Printf("similarity: %f\n", getSimilarityScore(lefts, rights))
}

func getSimilarityScore(lefts []float64, rights []float64) float64 {
	var similarity float64 = 0
	var counts = make(map[float64]float64)
	for _, right := range rights {
		counts[right] += 1.0
	}

	for _, left := range lefts {
		similarity += left * counts[left]
	}
	return similarity
}

func getDistance(lefts []float64, rights []float64) float64 {
	var distance float64 = 0
	sort.Slice(lefts, func(i, j int) bool {
		return lefts[i] < lefts[j]
	})
	sort.Slice(rights, func(i, j int) bool {
		return rights[i] < rights[j]
	})
	for i := 0; i < len(lefts); i++ {
		distance += math.Abs(lefts[i] - rights[i])
	}
	return distance
}

func getLeftsAndRights() ([]float64, []float64) {
	file, err := os.Open("/Users/michaelnewman/git/advent_of_code_2024/day_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lefts []float64
	var rights []float64
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "   ")
		left_val, left_err := strconv.ParseFloat(strs[0], 64)
		right_val, right_err := strconv.ParseFloat(strs[1], 64)
		if left_err != nil || right_err != nil {
			fmt.Println(left_err)
			fmt.Println(right_err)
		} else {
			lefts = append(lefts, left_val)
			rights = append(rights, right_val)
		}
	}
	return lefts, rights
}
