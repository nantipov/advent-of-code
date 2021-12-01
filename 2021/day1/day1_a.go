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
	var prevDepth int64 = -1
	increaseCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		depth, err := strconv.ParseInt(line, 10, 32)
		handleErr(err)
		if prevDepth > -1 && depth > prevDepth {
			increaseCounter++
		}
		prevDepth = depth
	}
	println(increaseCounter)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
