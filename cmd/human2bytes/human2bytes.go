package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"unicode"

	"github.com/dustin/go-humanize"
)

var sloppy = flag.Bool("sloppy", false, "continue past parse errors")

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s [NUMBER..]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s <FILE\n", os.Args[0])
	flag.PrintDefaults()
}

func convertLine(line string) (string, error) {
	split := strings.IndexFunc(line, unicode.IsSpace)
	if split == -1 {
		// no whitespace, interpret whole line as number
		split = len(line)
	}
	human := line[:split]
	rest := line[split:]

	number, err := humanize.ParseBigBytes(human)
	if err != nil {
		return line, err
	}
	robot := number.String()
	return robot + rest, nil
}

func main() {
	prog := path.Base(os.Args[0])
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = Usage
	flag.Parse()

	if flag.NArg() == 0 {
		// read stdin, interpret first word of each line
		in := bufio.NewReader(os.Stdin)
		for {
			line, err := in.ReadString('\n')
			if line != "" {
				out, err := convertLine(line)
				if !*sloppy && err != nil {
					log.Fatalf("cannot convert line: %s", err)
				}
				_, err = os.Stdout.WriteString(out)
				if err != nil {
					log.Fatalf("cannot write to standard output: %s", err)
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("cannot read standard input: %s", err)
			}
		}
	} else {
		for _, line := range flag.Args() {
			out, err := convertLine(line + "\n")
			if !*sloppy && err != nil {
				log.Fatalf("cannot convert line: %s", err)
			}
			_, err = os.Stdout.WriteString(out)
			if err != nil {
				log.Fatalf("cannot write to standard output: %s", err)
			}
		}
	}
}
