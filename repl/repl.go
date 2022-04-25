package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tehmantra/monkey/lexer"
	"github.com/tehmantra/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Fprintf(out, "%+v\n", tok)
		// }

		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			for _, e := range p.Errors() {
				fmt.Fprint(out, e)
			}
		} else {
			fmt.Fprint(out, program.String())
		}

		fmt.Fprintln(out)
	}
}
