package main

import (
	"bufio"
	"fmt"
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
		fmt.Println(reverse(line))
	})
}
func test() {
	Check(
		reverse(""), "",
		reverse("a"), "a",
		reverse("a b"), "b a",
		reverse("hello xiao mi"), "mi xiao hello",
	)
}

func reverse(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words)/2; i++ {
		j := len(words) - i - 1
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
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
