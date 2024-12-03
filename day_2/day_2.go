package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//The levels are either all increasing or all decreasing.
//Any two adjacent levels differ by at least one and at most three.

func main() {
	file, err := os.Open("/Users/michaelnewman/git/advent_of_code_2024/day_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var completelySafeCount int64 = 0
	var somewhatSafeCount int64 = 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		integerStrings := strings.Split(scanner.Text(), " ")
		endIndex := len(integerStrings) - 1

		completelySafe, unsafeIndex1, unsafeIndex2 := safe(-1, 0, endIndex, integerStrings)

		if completelySafe == true {
			completelySafeCount += 1
		} else if unsafeIndex1 == endIndex || unsafeIndex2 == endIndex {
			somewhatSafeCount += 1
		} else {
			if unsafeIndex1 == 1 || unsafeIndex2 == 1 {
				somewhatSafe, _, _ := safe(-1, 1.0, endIndex, integerStrings)
				if somewhatSafe == true {
					somewhatSafeCount += 1
					continue
				}
			}

			for i := 1; i < endIndex; i++ {
				somewhatSafe, _, _ := safe(i, 0, endIndex, integerStrings)
				if somewhatSafe == true {
					somewhatSafeCount += 1
					break
				} else if endIndex == i {
					fmt.Printf("%v\n", integerStrings)
				}
			}
		}

	}
	fmt.Printf("completely safe: %d\n", completelySafeCount)
	fmt.Printf("somewhat safe: %d\n", somewhatSafeCount)
	fmt.Printf("total safe: %d\n", completelySafeCount+somewhatSafeCount)
}

func safe(skipIndex int, startIndex int, endIndex int, vals []string) (bool, int, int) {
	isSafe, index := safeWithBounds(skipIndex, startIndex, endIndex, 1.0, 3.0, vals)
	if isSafe == true {
		return true, -1, -1
	} else {
		isSafe2, index2 := safeWithBounds(skipIndex, startIndex, endIndex, -3.0, -1.0, vals)
		return isSafe2, index, index2
	}
}

func safeWithBounds(skipIndex int, startIndex int, endIndex int, min float64, max float64, vals []string) (bool, int) {
	if min == max || len(vals) == 0 {
		return false, -1
	}

	for i := startIndex; i < endIndex; i++ {
		val1, err1 := strconv.ParseFloat(vals[i], 64)

		if (i + 1) == skipIndex {
			i += 1
		}

		if i >= endIndex {
			break
		}

		val2, err2 := strconv.ParseFloat(vals[i+1], 64)

		if err1 != nil || err2 != nil {
			fmt.Println(err1)
			fmt.Println(err2)
			return false, -1
		} else {
			diff := val2 - val1
			if diff < min || diff > max {
				return false, (i + 1)
			}
		}
	}
	return true, 0
}
