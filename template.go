package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile)
	nextReader = NewScanner()
	line1 := nextInt()
	line2 := nextInt()
	sort.Sort(IntSorter(line2))
	line3 := nextFloat64()
	sort.Sort(Float64Sorter(line3))
	x := line2

	fmt.Printf("line1= %+v\nx = %+v\nline3=%+v\n", line1, x, line3)
}

// ------ Mathライブラリ ---------------------------------//
func max(a ...int) int {
	r := a[0]
	for i := 0; i < len(a); i++ {
		if r < a[i] {
			r = a[i]
		}
	}
	return r
}

func maxIdx(a ...int) int {
	r := a[0]
	index := 0
	for i := 0; i < len(a); i++ {
		if r < a[i] {
			r = a[i]
			index = i
		}
	}
	return index
}

func min(a ...int) int {
	r := a[0]
	for i := 0; i < len(a); i++ {
		if r > a[i] {
			r = a[i]
		}
	}
	return r
}

func minIdx(a ...int) int {
	r := a[0]
	index := 0
	for i := 0; i < len(a); i++ {
		if r > a[i] {
			r = a[i]
			index = i
		}
	}
	return index
}

func sum(a []int) (r int) {
	for i := range a {
		r += a[i]
	}
	return r
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// ------ Mathライブラリ ---------------------------------//

// ----- 標準入力用の関数 ----------------------------------//
var nextReader func() []string

func NewScanner() func() []string {
	r := bufio.NewScanner(os.Stdin)
	r.Buffer(make([]byte, 1024), int(1e+11))
	return func() []string {
		r.Scan()
		return strings.Fields(r.Text())
	}
}

func nextString() []string {
	return nextReader()
}

func nextInt() []int {
	var intArray []int
	strArray := nextReader()
	for _, v := range strArray {
		i, _ := strconv.Atoi(v)
		intArray = append(intArray, i)
	}
	return intArray
}

func nextFloat64() []float64 {
	var floatArray []float64
	strArray := nextReader()
	for _, v := range strArray {
		f, _ := strconv.ParseFloat(v, 64)
		floatArray = append(floatArray, f)
	}
	return floatArray
}

// ------ 標準入力用の関数 ---------------------------------//

// ------ ソートライブラリ ---------------------------------//
type IntSorter []int

func (s IntSorter) Len() int           { return len(s) }
func (s IntSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSorter) Less(i, j int) bool { return s[i] < s[j] }

type Float64Sorter []float64

func (s Float64Sorter) Len() int           { return len(s) }
func (s Float64Sorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Float64Sorter) Less(i, j int) bool { return s[i] < s[j] }

// ------ ソートライブラリ ---------------------------------//

// ------ あまり使わない -----------------------------------//
func nextInt64() []int64 {
	var int64Array []int64
	strArray := nextReader()
	for _, v := range strArray {
		i, _ := strconv.ParseInt(v, 10, 64)
		intArray = append(intArray, i)
	}
	return intArray
}

// ------ あまり使わない -----------------------------------//
