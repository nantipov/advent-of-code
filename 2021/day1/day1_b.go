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

	const measures = 3
	m := make([]int64, measures+1)
	lagIndex := 0
	increaseCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		depth, err := strconv.ParseInt(line, 10, 32)
		handleErr(err)

		m[lagIndex] = depth
		lagIndex++
		if lagIndex == len(m) {
			if isDiffCounted(m) {
				increaseCounter++
			}
			shiftM(m)
			lagIndex = len(m) - 1
		}
	}
	println(increaseCounter)
}

func isDiffCounted(m []int64) bool {
	var s1 int64 = 0
	var s2 int64 = 0
	for i := 0; i < len(m)-1; i++ {
		s1 = s1 + m[i]
		s2 = s2 + m[i+1]
	}
	return s2-s1 > 0
}

func shiftM(m []int64) {
	for i := 1; i < len(m); i++ {
		m[i-1] = m[i]
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
