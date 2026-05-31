package main

import (
	"flag"
	"fmt"
	"go-cli/counter"
	"log"
	"os"
	"sync"
)

func main() {
	log.SetFlags(0)

	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		total int
	)
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
		wg.Add(1)
		go func() {
			defer wg.Done()
			wordCount, err := counter.CountWordsInFile(file)
			if err != nil {
				fmt.Fprintln(os.Stderr, "counter:", err)
				return
			}
			mu.Lock()
			defer mu.Unlock()
			total = total + wordCount.Words
			fmt.Printf("%s\t", file)
			wordCount.Print(os.Stdout, s)
		}()
	}
	if len(filenames) == 0 {
		wordCount := counter.GetCounts(os.Stdin)
		wordCount.Print(os.Stdout, s)
	}
	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

}
