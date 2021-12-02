package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Input argument expected - filename")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	handleErr(err)
	defer file.Close()

	var depth int64 = 0
	var position int64 = 0

	for {
		var command string
		var value int64
		_, err := fmt.Fscanf(file, "%s %d\n", &command, &value)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				handleErr(err)
			}
		}

		switch command {
		case "forward":
			position = position + value
		case "down":
			depth = depth + value
		case "up":
			depth = depth - value
		default:
			panic("Unknown command " + command)
		}
	}
	println(depth * position)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
