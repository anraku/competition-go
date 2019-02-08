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

var maxDiff int
var diffs []int
var sumNum int

func main() {
	log.SetFlags(log.Lshortfile)
	// 入力値を取得
	nextReader = NewScanner()
	line1 := nextInt()
	line2 := nextInt()
	sort.Ints(line2)

	koma_count := line1[0]
	place_count := line1[1]
	x := line2

	// コマの数と地点の数が同じなら移動回数は0になる
	if len(x) <= koma_count {
		fmt.Println(0)
		return
	}

	// それぞれの地点の差を取得
	// x1-x2, x1-x3, x1-x4, ...
	for i := 0; i < len(x)-1; i++ {
		diffs = append(diffs, x[i+1]-x[i])
	}

	sort.Ints(diffs)

	var ans int
	for i := 0; i < place_count-koma_count; i++ {
		ans += diffs[i]
	}

	fmt.Println(ans)
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

// ------ あまり使わない -----------------------------------//
func nextInt64() []int64 {
	var int64Array []int64
	strArray := nextReader()
	for _, v := range strArray {
		i, _ := strconv.ParseInt(v, 10, 64)
		int64Array = append(int64Array, i)
	}
	return int64Array
}

// ------ あまり使わない -----------------------------------//
