package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"translate"
)

func main() {
	sourceLang := flag.String("from", "ro", "language to translate from")
	destLang := flag.String("to", "en", "language to translate to")
	flag.Parse()

	text := readInputText()

	translation, err := translate.Translate(text, *sourceLang, *destLang)
	if err != nil {
		fmt.Printf("could not translate %q from %q to %q: %v", text, *sourceLang, *destLang, err)
		os.Exit(1)
	}

	fmt.Println(translation)
}

func readInputText() string {
	if hasPipedInput() {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("could not read input from stdin: %v", err)
			os.Exit(1)
		}

		return strings.TrimSpace(string(b))
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Print("you must translate something, don't you?")
		os.Exit(1)
	}

	return strings.Join(args, " ")
}

func hasPipedInput() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	return stat.Mode()&os.ModeCharDevice == 0
}
