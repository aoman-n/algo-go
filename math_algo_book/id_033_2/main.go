package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var reader io.Reader
var writer io.Writer

const MOD = 1000000007

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

type point struct {
	x, y float64
}

func newPoint(x, y int) *point {
	return &point{
		x: float64(x),
		y: float64(y),
	}
}

func (p *point) getVector(other *point) *vector {
	return &vector{
		x: other.x - p.x,
		y: other.y - p.y,
	}
}

type vector struct {
	x, y float64
}

func (v *vector) len() float64{
	return math.Abs(
		math.Sqrt((v.x * v.x) + (v.y * v.y)),
	)
}

func (v *vector) inner(other *vector) float64{
	return v.x*other.x + v.y*other.y
}

func (v *vector) outer(other *vector) float64{
	return math.Abs(v.x*other.y - v.y*other.x)
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	pointA := newPoint(ni2(sc))
	pointB := newPoint(ni2(sc))
	pointC := newPoint(ni2(sc))

	vectorBA := pointB.getVector(pointA)
	vectorBC := pointB.getVector(pointC)

	vectorCB := pointC.getVector(pointB)
	vectorCA := pointC.getVector(pointA)

	if vectorBA.inner(vectorBC) <= 0 {
		fmt.Fprintf(writer, "%.12f", vectorBA.len())
		return
	}
	
	if vectorCB.inner(vectorCA) <= 0 {
		fmt.Fprintf(writer, "%.12f", vectorCA.len())
		return
	}
	
	answer := vectorBA.outer(vectorBC) / vectorBC.len()
	fmt.Fprintf(writer, "%.12f", answer)
}

// ==================================================
// io
// ==================================================

func ni(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func ni2(sc *bufio.Scanner) (int, int) {
	return ni(sc), ni(sc)
}