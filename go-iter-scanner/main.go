package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"iter"
	"unicode"
)

type Token struct {
	kind    string
	literal []rune
}

func (t Token) String() string {
	return fmt.Sprintf("{%v, %q}", t.kind, string(t.literal))
}

func runes(r io.Reader) iter.Seq[rune] {
	buf := bufio.NewReaderSize(r, 100)

	return func(yield func(rune) bool) {
		for {
			r, _, err := buf.ReadRune()
			if err != nil || !yield(r) {
				return
			}
		}
	}
}

func tokens(runes iter.Seq[rune]) iter.Seq[Token] {
	return func(yield func(Token) bool) {
		next, _ := iter.Pull(runes)
		buf := make([]rune, 0, 512) // reusable

		for {
			r, keepgoing := next()
			if !keepgoing {
				return
			}
			switch {
			case r == '=':
				if !yield(Token{"EQ", []rune{r}}) {
					return
				}
			case unicode.IsDigit(r):
				buf = append(buf, r)
				for r, ok := next(); ok && unicode.IsDigit(r) || r == '.'; r, ok = next() {
					buf = append(buf, r)
				}
				if !yield(Token{"NUM", buf}) {
					return
				}
			case unicode.IsLetter(r):
				buf = append(buf, r)
				for r, ok := next(); ok && unicode.IsLetter(r); r, ok = next() {
					buf = append(buf, r)
				}
				if !yield(Token{"SYM", buf}) {
					return
				}
			default:
				// whitespace
			}
			// reset slice but keep the underlying array
			// setting to nil will cause a new allocation
			buf = buf[:0]
		}
	}
}

var code = []byte(`if true then
	1
else
	2
end		
x = 10
π = 3.1416
`)

func main() {
	println("-- code")
	print(string(code))

	b := bytes.NewReader(code)
	println("-- runes")
	for r := range runes(b) {
		fmt.Printf("[%q]", r)
	}
	println()

	println("-- tokens func")
	b = bytes.NewReader(code)
	for t := range tokens(runes(b)) {
		fmt.Printf("tok: %+v\n", t)
	}
}
