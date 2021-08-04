package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p(s.Contains("test", "es"))        // true
	p(s.Count("test", "t"))            // 2
	p(s.HasPrefix("test", "te"))       // true
	p(s.HasSuffix("test", "st"))       // true
	p(s.Index("test", "t"))            // 0
	p(s.LastIndex("test", "t"))        // 3
	p(s.Join([]string{"a", "b"}, "-")) // a-b
	p(s.Repeat("a", 5))                // aaaaa
	p(s.Replace("fooo", "o", "O", -1)) // fOOO
	p(s.Replace("fooo", "o", "O", 2))  // fOOo
	p(s.Split("a-b-c", "-"))           // [a b c]
	p(s.ToLower("TEST"))               // test
	p(s.ToUpper("test"))               // TEST
	p(len("hello"))                    // 5
	p("hello"[1])                      // 1
}
