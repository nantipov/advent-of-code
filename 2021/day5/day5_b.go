package main

import (
	"fmt"
	"io"
	"math"
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
		id         int
		x1         int
		y1         int
		x2         int
		y2         int
		isVertical bool
		k          float64
		b          float64
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

	lineId := 0
	for {
		_, err := fmt.Fscanf(file, "%d,%d -> %d,%d\n", &x1, &y1, &x2, &y2)
		if err == io.EOF {
			break
		} else {
			handleErr(err)
		}

		lineId++
		line := line(lineId, x1, y1, x2, y2)

		//fmt.Printf("line=%+v\n", line)

		if math.Abs(line.k) != 0.0 && math.Abs(line.k) != 1.0 {
			// exclude non-45 degrees multiplies lines
			continue
		}

		markIntersections(lines, &intersections, line)

		lines.add(line)
	}

	println(intersections.count)
}

func line(id, x1, y1, x2, y2 int) Line {
	isVertical := x1 == x2

	k := 0.0
	b := 0.0
	if !isVertical {
		k = (float64(y2) - float64(y1)) / (float64(x2) - float64(x1))
		b = float64(y1) - k*float64(x1)
	}
	return Line{
		id:         id,
		x1:         x1,
		y1:         y1,
		x2:         x2,
		y2:         y2,
		isVertical: isVertical,
		k:          k,
		b:          b,
	}
}

func markIntersections(lines List, intersections *List, line Line) {
	iterator := lines.iterator()
	for iterator.hasNext() {
		anotherLine := iterator.next().(Line)
		foundInterpoints := findIntersection(line, anotherLine)
		for _, foundInterpoint := range foundInterpoints {
			if !existingIntersection(foundInterpoint, intersections) {
				//fmt.Printf("--- %+v === %+v\n", foundInterpoint, anotherLine)
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
	if line1.isVertical {
		for y := minInt(line1.y1, line1.y2); y <= maxInt(line1.y1, line1.y2); y++ {
			//fmt.Printf("*** %d, %d\n", line1.x1, y)
			if pointBelongsToLine(float64(line1.x1), float64(y), line2) {
				interpoints = append(interpoints, Interpoint{
					x: line1.x1,
					y: y,
				})
			}
		}
	} else {
		for x := minInt(line1.x1, line1.x2); x <= maxInt(line1.x1, line1.x2); x++ {
			y := line1.k*float64(x) + line1.b
			//fmt.Printf("*** %d, %f\n", x, y)
			if pointBelongsToLine(float64(x), y, line2) {
				interpoints = append(interpoints, Interpoint{
					x: x,
					y: int(y),
				})
			}
		}
		//fmt.Println()
	}
	return interpoints
}

func pointBelongsToLine(x, y float64, line Line) bool {
	if line.isVertical {
		return int(x) == line.x1 && int(y) >= minInt(line.y1, line.y2) && int(y) <= maxInt(line.y1, line.y2)
	} else {
		return int(x) >= minInt(line.x1, line.x2) && int(x) <= maxInt(line.x1, line.x2) && line.k*x+line.b == y
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
