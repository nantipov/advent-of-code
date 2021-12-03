package main

import (
	"bufio"
	"os"
	"strconv"
)

type (
	node struct {
		id             int
		quanity        int
		quantityZeroes int
		position       int
		node0          *node
		node1          *node
		value          string
	}
)

var (
	nodeIds = 0
)

func main() {
	if len(os.Args) < 2 {
		panic("Input argument expected - filename")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	handleErr(err)
	defer file.Close()

	nodeIds++
	root := node{
		id:             nodeIds,
		quanity:        0,
		quantityZeroes: 0,
		position:       1,
		node0:          nil,
		node1:          nil,
		value:          "",
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		add(&root, 1, line)
	}
	binaryOxygen := search(&root, fnOxygen).value
	binaryCO2 := search(&root, fnCO2).value
	oxygen, err := strconv.ParseInt(binaryOxygen, 2, 64)
	handleErr(err)
	co2, err := strconv.ParseInt(binaryCO2, 2, 64)
	handleErr(err)

	println("oxygen ", binaryOxygen, " ", oxygen)
	println("co2    ", binaryCO2, " ", co2)

	println(oxygen * co2)
}

func add(currentNode *node, targetPosition int, value string) {
	isZero := false
	if value[targetPosition-1] == '0' {
		isZero = true
	}

	if currentNode.position == targetPosition {
		currentNode.quanity++
		currentNode.value = value

		if isZero {
			currentNode.quantityZeroes++
		}
		//print("added ", " (", value, ") ")
		//	printNode(currentNode)
	}

	if targetPosition < len(value) {
		if isZero {
			currentNode.node0 = tryNode(currentNode.node0, targetPosition+1)
			add(currentNode.node0, targetPosition+1, value)
		} else {
			currentNode.node1 = tryNode(currentNode.node1, targetPosition+1)
			add(currentNode.node1, targetPosition+1, value)
		}
	}
}

func tryNode(n *node, position int) *node {
	if n != nil {
		return n
	} else {
		nodeIds++
		return &node{
			id:             nodeIds,
			quanity:        0,
			quantityZeroes: 0,
			position:       position,
			node0:          nil,
			node1:          nil,
			value:          "",
		}
	}
}

func search(n *node, searchFn func(*node) *node) *node {
	if n == nil {
		panic("Cannot search within the nil node")
	}

	print("[current ]")
	printNode(n)

	nextNode := searchFn(n)

	print("[next    ]")
	printNode(nextNode)

	if nextNode == nil || nextNode.quanity == 0 {
		return n
	}
	if nextNode.quanity == 1 {
		return nextNode
	}
	return search(nextNode, searchFn)
}

func fnOxygen(n *node) *node {
	if n.quantityZeroes > n.quanity/2 {
		println("OXY follow to 0 (majority) ")
		return n.node0
	} else {
		println("OXY follow to 1 (majority or equal) ")
		return n.node1
	}
}

func fnCO2(n *node) *node {
	if n.quantityZeroes > n.quanity/2 {
		println("CO2 follow to 1 (minority) ")
		return n.node1
	} else {
		println("CO2 follow to 0 (minority or equal) ")
		return n.node0
	}
}

func printNode(n *node) {
	if n == nil {
		println()
		return
	}
	println("id=", n.id, ", q=", n.quanity, ", q0=", n.quantityZeroes, ", q1=", n.quanity-n.quantityZeroes, ", pos=", n.position, ", value=", n.value)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
