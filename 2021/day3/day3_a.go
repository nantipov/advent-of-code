package main

import (
	"bufio"
	"os"
	"strconv"
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
	lines := 0
	zeroes := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		zeroes = countZeroes(line, zeroes)
		lines++
	}
	binaryGamma := ""
	binaryEpsilon := ""
	for _, count := range zeroes {
		if count > lines/2 {
			binaryGamma = binaryGamma + "0"
			binaryEpsilon = binaryEpsilon + "1"
		} else {
			binaryGamma = binaryGamma + "1"
			binaryEpsilon = binaryEpsilon + "0"
		}
	}
	gamma, err := strconv.ParseInt(binaryGamma, 2, 64)
	handleErr(err)
	epsilon, err := strconv.ParseInt(binaryEpsilon, 2, 64)
	println(gamma * epsilon)
}

func countZeroes(line string, collectedZeroes []int) []int {
	output := collectedZeroes
	if len(collectedZeroes) < len(line) {
		output = make([]int, len(line))
		copy(output, collectedZeroes)
	}
	for i, r := range line {
		if r == '0' {
			output[i] = output[i] + 1
		}
	}
	return output
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
