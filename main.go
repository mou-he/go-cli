package main

import (
	"flag"
	"fmt"
	"go-cli/counter"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	total := 0
	bytevar := false
	wordvar := false
	linevar := false
	flag.BoolVar(&bytevar, "b", false, "count bytes")
	flag.BoolVar(&wordvar, "w", false, "count words")
	flag.BoolVar(&linevar, "l", false, "count lines")
	flag.Parse()
	s := counter.Settings{
		Byte: bytevar,
		Word: wordvar,
		Line: linevar,
	}
	filenames := flag.Args()
	for _, file := range filenames {
		wordCount, err := counter.CountWordsInFile(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}
		total = total + wordCount.Words
		fmt.Printf("%s\t", file)
		wordCount.Print(os.Stdout, s)
	}
	if len(filenames) == 0 {
		wordCount := counter.GetCounts(os.Stdin)
		wordCount.Print(os.Stdout, s)
	}
	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

}
