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

	var (
		// mu    sync.Mutex
		total int
	)
	didError := false
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
	ch, errch := counter.CountFlies(filenames)

	for {
		select {
		case res, open := <-ch:
			if !open {
				ch = nil
				break
			}
			total += res.Words
			res.Counts.Print(os.Stdout, s)

		case err, open := <-errch:
			if !open {
				errch = nil
				break
			}
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
		}
		if ch == nil || errch == nil {
			break
		}
	}
	if len(filenames) == 0 {
		wordCount := counter.GetCounts(os.Stdin)
		wordCount.Print(os.Stdout, s)
	}
	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}
	if didError {
		os.Exit(1)
	}

}
