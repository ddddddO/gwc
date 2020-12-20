package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ddddddO/gwc/wc"
)

func main() {
	var (
		isHeader bool // show header

		isBytes         bool // wc -c
		isChars         bool // wc -m
		isLines         bool // wc -l
		isWords         bool // wc -w
		isMaxLineLength bool // wc -L
	)
	flag.BoolVar(&isHeader, "h", false, "print the wc options header")
	flag.BoolVar(&isBytes, "c", false, "print the byte counts")
	flag.BoolVar(&isChars, "m", false, "print the character counts")
	flag.BoolVar(&isLines, "l", false, "print the newline counts")
	flag.BoolVar(&isWords, "w", false, "print the word counts")
	flag.BoolVar(&isMaxLineLength, "L", false, "print the maximum display width")
	flag.Parse()

	opts := wc.Options{
		IsHeader:        isHeader,
		IsBytes:         isBytes,
		IsChars:         isChars,
		IsLines:         isLines,
		IsWords:         isWords,
		IsMaxLineLength: isMaxLineLength,
	}

	files := flag.Args()
	if len(files) == 0 {
		os.Exit(1)
	}

	if err := wc.Wc(opts, files); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
