package repl

import (
	"bufio"
	"fmt"
	"gjko/lexer"
	"gjko/token"
	"io"
)

const PROMT = ">> "

//starts lexer
/* line takes line of input text, with l it creates
instance of lexer.New() and print tokens and a copy of
the literal until it encounters the EOF*/
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintln(out, PROMT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
