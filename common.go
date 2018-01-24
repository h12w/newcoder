package nowcoder

import (
	"bufio"
	"log"
	"os"
	"reflect"
)

// Common code

func Lines(fn func(s string)) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fn(s.Text())
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
