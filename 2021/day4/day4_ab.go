package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	boardWidth  = 5
	boardHeight = 5
)

type (
	board struct {
		numbers  []int64
		cols     []int
		rows     []int
		totalSum int64
	}
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
	lineDrawnNumbers := scanner.Text()
	drawnNumbers := toIntArray(strings.Split(lineDrawnNumbers, ","))

	currentBoard := newBoard()

	col := 0
	row := 0

	leastTurns := -1
	var leastTurnsScore int64 = 0
	mostTurns := -1
	var mostTurnsScore int64 = 0

	for scanner.Scan() {
		boardLine := scanner.Text()
		if len(boardLine) < 1 {
			continue
		}

		boardElements := strings.Split(boardLine, " ")
		boardNumbersRow := toIntArray(boardElements)

		for _, boardNumber := range boardNumbersRow {
			currentBoard.totalSum = currentBoard.totalSum + boardNumber
			currentBoard.numbers[col*boardWidth+row] = boardNumber
			col++
			if col == boardWidth {
				col = 0
				row++
			}
			if row == boardHeight {
				isWon, score, turn := play(drawnNumbers, currentBoard)
				//printBoard(currentBoard)
				if isWon {
					if leastTurns < 0 || turn < leastTurns {
						leastTurns = turn
						leastTurnsScore = score
					}
					if mostTurns < 0 || turn > mostTurns {
						mostTurns = turn
						mostTurnsScore = score
					}
				}
				col = 0
				row = 0
				currentBoard = newBoard()
			}
		}

	}

	println("least turns score: ", leastTurnsScore)
	println("most turns score:  ", mostTurnsScore)
}

/*
Plays and returns: isWon, score, turn
*/
func play(drawnNumbers []int64, b board) (bool, int64, int) {
	var markedSum int64 = 0
	for turn, drawnNumber := range drawnNumbers {
		isFound, col, row := searchInBoard(drawnNumber, b)
		if isFound {
			markedSum = markedSum + drawnNumber
			b.cols[col] = b.cols[col] + 1
			b.rows[row] = b.rows[row] + 1
			if b.cols[col] == boardWidth || b.rows[row] == boardHeight {
				return true, (b.totalSum - markedSum) * drawnNumber, turn
			}
		}
	}
	return false, 0, 0
}

func searchInBoard(n int64, b board) (bool, int, int) {
	for i, bn := range b.numbers {
		if bn == n {
			return true, i % boardWidth, i / boardWidth
		}
	}
	return false, 0, 0
}

func newBoard() board {
	return board{
		numbers:  make([]int64, boardWidth*boardHeight),
		cols:     make([]int, boardWidth),
		rows:     make([]int, boardHeight),
		totalSum: 0,
	}
}

func toIntArray(strArray []string) []int64 {
	var intArray []int64
	for _, e := range strArray {
		if len(e) < 1 {
			continue
		}
		num, err := strconv.ParseInt(e, 10, 8)
		handleErr(err)

		intArray = append(intArray, num)
	}
	return intArray
}

func printBoard(b board) {
	fmt.Printf("b=%v; cols=%v; rows=%v; totalSum=%d\n", b.numbers, b.cols, b.rows, b.totalSum)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
