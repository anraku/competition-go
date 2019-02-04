package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input1 := strings.Fields(scanner.Text())
	log.SetFlags(log.Lshortfile)
	k, _ := strconv.Atoi(input1[1])

	scanner.Scan()
	input2 := strings.Fields(scanner.Text())
	var a []int
	for _, v := range input2 {
		i, _ := strconv.Atoi(v)
		a = append(a, i)
	}

	var max int
	for i := k; i >= k/2; i-- {
		var v int
		for _, x := range a {
			v += x ^ i
		}
		if max < v {
			max = v
		}
	}

	//var v int
	//for _, x := range a {
	//	v += x ^ k
	//}
	//if max < v {
	//	max = v
	//}
	fmt.Println(max)

}

func max(a ...int) int {
	r := a[0]
	for i := 0; i < len(a); i++ {
		if r < a[i] {
			r = a[i]
		}
	}
	return r
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

var nextReader func() string

func NewScanner() func() string {
	r := bufio.NewScanner(os.Stdin)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)
	return func() string {
		r.Scan()
		return r.Text()
	}
}
func nextString() string {
	return nextReader()
}
func nextInt64() int64 {
	v, _ := strconv.ParseInt(nextReader(), 10, 64)
	return v
}
func nextInt() int {
	v, _ := strconv.Atoi(nextReader())
	return v
}
func nextInts(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = nextInt()
	}
	return r
}
func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(nextReader(), 64)
	return f
}
