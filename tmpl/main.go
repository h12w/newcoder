package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "test" {
		test()
		return
	}
	Lines(func(line string) {
	})
}
func test() {
	Check()
}


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
