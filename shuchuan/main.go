package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "test" {
		test()
		return
	}
	i := 0
	Lines(func(line string) {
		if i%2 == 0 {
			nums := strings.Split(line, " ")
			fmt.Println(maxNumString(nums...))
		}
		i++
	})
}
func test() {
	Check(
		maxNumString("12", "123"), "12312",
	)
}

type byNum []string

func (a byNum) Len() int      { return len(a) }
func (a byNum) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byNum) Less(i, j int) bool {
	x, y := a[i], a[j]
	return x+y < y+x // key point
}

func maxNumString(nums ...string) string {
	sort.Sort(sort.Reverse(byNum(nums)))
	return strings.Join(nums, "")
}

// Common code

func Lines(fn func(s string)) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fn(strings.TrimSpace(s.Text()))
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}

func Check(s ...interface{}) {
	if len(s)%2 != 0 {
		log.Fatal("checked variables should be paired")
	}
	for i := 0; i < len(s)/2; i++ {
		a, b := s[i*2], s[i*2+1]
		if !reflect.DeepEqual(a, b) {
			log.Fatalf("expect\n%v\ngot\n%v", b, a)
		}
	}
}
