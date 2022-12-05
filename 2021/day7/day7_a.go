package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		panic("Input argument expected - filename")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	handleErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		panic("Could not read the line of numbers")
	}
	crabsLine := scanner.Text()
	crabs := toIntArray(strings.Split(crabsLine, ","))

	positions := make(map[int]int)

	minPos := -1
	maxPos := -1
	for _, pos := range crabs {
		positions[pos]++
		if minPos < 0 || pos < minPos {
			minPos = pos
		}
		if maxPos < 0 || pos > maxPos {
			maxPos = pos
		}
	}

	medianPos := (maxPos - minPos) / 2
	targetPos := medianPos
	minStep := len(crabs) / medianPos
	for pos, count := range positions {
		if pos < targetPos {
			targetPos = targetPos - minStep*count
		}
		fmt.Printf("%2d: %d\n", pos, count)
	}

	fmt.Printf("ws  = %d\n", targetPos)

	fuel := 0
	for _, pos := range crabs {
		fuel = fuel + absInt(targetPos-pos)
	}

// 359474 - too high

	println(fuel)
}

func absInt(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}

func toIntArray(strArray []string) []int {
	var intArray []int
	for _, e := range strArray {
		if len(e) < 1 {
			continue
		}
		num, err := strconv.ParseInt(e, 10, 32)
		handleErr(err)

		intArray = append(intArray, int(num))
	}
	return intArray
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
