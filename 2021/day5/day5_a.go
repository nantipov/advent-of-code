package main

import (
	"fmt"
	"io"
	"os"
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

	Line struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}

	Interpoint struct {
		x int
		y int
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

	lines := newList()
	intersections := newList()

	var x1, y1, x2, y2 int

	for {
		_, err := fmt.Fscanf(file, "%d,%d -> %d,%d\n", &x1, &y1, &x2, &y2)
		if err == io.EOF {
			break
		} else {
			handleErr(err)
		}

		if !(x1 == x2 || y1 == y2) {
			continue
		}

		line := Line{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}
		markIntersections(lines, &intersections, line)

		lines.add(line)
	}

	println(intersections.count)
}

func markIntersections(lines List, intersections *List, line Line) {
	iterator := lines.iterator()
	for iterator.hasNext() {
		anotherLine := iterator.next().(Line)
		foundInterpoints := findIntersection(line, anotherLine)
		for _, foundInterpoint := range foundInterpoints {
			if !existingIntersection(foundInterpoint, intersections) {
				intersections.add(foundInterpoint)
			}
		}
	}
}

/*
Returns slice of intersection points
*/
func findIntersection(line1, line2 Line) []Interpoint {
	interpoints := make([]Interpoint, 0)
	if isLineHorizontal(line1) {
		for x := minInt(line1.x1, line1.x2); x <= maxInt(line1.x1, line1.x2); x++ {
			if pointBelongsToLine(x, line1.y1, line2) {
				interpoints = append(interpoints, Interpoint{
					x: x,
					y: line1.y1,
				})
			}
		}
	} else {
		for y := minInt(line1.y1, line1.y2); y <= maxInt(line1.y1, line1.y2); y++ {
			if pointBelongsToLine(line1.x1, y, line2) {
				interpoints = append(interpoints, Interpoint{
					x: line1.x1,
					y: y,
				})
			}
		}
	}
	return interpoints
}

func isLineHorizontal(line Line) bool {
	return line.y1 == line.y2
}

func pointBelongsToLine(x, y int, line Line) bool {
	if isLineHorizontal(line) {
		return y == line.y1 && x >= minInt(line.x1, line.x2) && x <= maxInt(line.x1, line.x2)
	} else {
		return x == line.x1 && y >= minInt(line.y1, line.y2) && y <= maxInt(line.y1, line.y2)
	}
}

func minInt(i1, i2 int) int {
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}

func maxInt(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func existingIntersection(interpoint Interpoint, intersections *List) bool {
	iterator := intersections.iterator()
	f := 0
	for iterator.hasNext() {
		f++
		intersection := iterator.next().(Interpoint)
		if interpoint.x == intersection.x && interpoint.y == intersection.y {
			return true
		}
	}
	return false
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

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
