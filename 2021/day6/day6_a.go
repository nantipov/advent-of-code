package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	List struct {
		first *ListItem
		last  *ListItem
		count int
	}

	ListItem struct {
		data interface{}
		next *ListItem
	}

	Iterator struct {
		list    *List
		current *ListItem
	}

	Timer struct {
		value int64
	}
)

func main() {
	if len(os.Args) < 3 {
		panic("Input argument expected - filename, days")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	handleErr(err)
	defer file.Close()

	days, err := strconv.ParseInt(os.Args[2], 10, 32)
	handleErr(err)

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		panic("Could not read the line of numbers")
	}
	initialTimersLine := scanner.Text()
	initialTimers := toIntArray(strings.Split(initialTimersLine, ","))

	timers := newList()
	for _, t := range initialTimers {
		timers.add(&Timer{
			value: t,
		})
	}

	//fmt.Printf("Day %2d: ", 0)
	//printTimers(timers)

	for day := 1; day <= int(days); day++ {
		evaluate(&timers)
		//fmt.Printf("Day %2d: [%5d]:\n", day, timers.count)
		//printTimers(timers)
	}

	println(timers.count)
}

func evaluate(timers *List) {
	iterator := timers.iterator()
	timersToAdd := 0
	for iterator.hasNext() {
		t := iterator.next().(*Timer)
		t.value--
		if t.value < 0 {
			t.value = 6
			timersToAdd++
		}
	}
	for timersToAdd > 0 {
		timers.add(&Timer{
			value: 8,
		})
		timersToAdd--
	}
}

func printTimers(timers List) {
	iterator := timers.iterator()
	for iterator.hasNext() {
		t := iterator.next().(*Timer)
		fmt.Printf("%2d ", t.value)
	}
	fmt.Println()
}

func newList() List {
	return List{
		first: nil,
		last:  nil,
		count: 0,
	}
}

func (l *List) add(data interface{}) {
	i := ListItem{
		data: data,
		next: nil,
	}
	if l.first == nil {
		l.first = &i
	} else {
		l.last.next = &i
	}
	l.last = &i
	l.count++
}

func (l List) iterator() Iterator {
	return Iterator{
		list:    &l,
		current: nil,
	}
}

func (it Iterator) hasNext() bool {
	if it.current == nil {
		return it.list.first != nil
	}
	return it.current.next != nil
}

func (it *Iterator) next() interface{} {
	if !it.hasNext() {
		panic("no more items in list")
	}
	if it.current == nil {
		if it.list.first != nil {
			it.current = it.list.first
		}
	} else {
		if it.current.next != nil {
			it.current = it.current.next
		}
	}
	return it.current.data
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

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
