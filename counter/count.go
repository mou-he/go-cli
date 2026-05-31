package counter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Counts struct {
	Words int
	Bytes int
	Lines int
}

//	func CountWordsInFile(filename string) (int, error) {
//		file, err := os.Open(filename)
//		if err != nil {
//			log.Fatal("failed to read file", err)
//		}
//		return CountWords(file), nil
//	}
type SeekerReader interface {
	io.Seeker
	io.Reader
}
type Settings struct {
	Byte bool
	Word bool
	Line bool
}
type FileCounts struct {
	Counts
	FileName string
}

func (c Counts) Print(w io.Writer, s Settings) {
	// 右对齐，固定宽度 10
	const format = "%10s"

	if !s.Byte && !s.Word && !s.Line {
		fmt.Fprintf(w, format+format+format+"\n",
			fmt.Sprintf("words: %d", c.Words),
			fmt.Sprintf("bytes: %d", c.Bytes),
			fmt.Sprintf("lines: %d", c.Lines),
		)
		return
	}

	var parts []string
	if s.Word {
		parts = append(parts, fmt.Sprintf(format, fmt.Sprintf("words: %d", c.Words)))
	}
	if s.Byte {
		parts = append(parts, fmt.Sprintf(format, fmt.Sprintf("bytes: %d", c.Bytes)))
	}
	if s.Line {
		parts = append(parts, fmt.Sprintf(format, fmt.Sprintf("lines: %d", c.Lines)))
	}

	fmt.Fprintln(w, strings.Join(parts, ""))
}

func GetCounts(file io.ReadSeeker) Counts {
	res := Counts{}
	isInsideWord := false
	reader := bufio.NewReader(file)
	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			break
		}
		res.Bytes += size
		if r == '\n' {
			res.Lines++
		}
		isSpace := unicode.IsSpace(r)
		if !isSpace && !isInsideWord {
			res.Words++
		}
		isInsideWord = !isSpace

	}
	// a book
	return res
}
func CountWordsInFile(filename string) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{0, 0, 0}, err
	}
	defer file.Close()
	counts := GetCounts(file)
	return counts, nil
}

func CountWords(file io.Reader) int {
	wordCounts := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordCounts++
	}
	return wordCounts
}

func CountBytes(file io.Reader) int {
	count := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		count++
	}
	return count
}
func CountLines(file io.Reader) int {
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}
	return count
}
func CountFlies(filenames []string) (<-chan FileCounts, <-chan error) {
	ch := make(chan FileCounts)
	errch := make(chan error)
	wg := sync.WaitGroup{}
	wg.Add(len(filenames))
	for _, filename := range filenames {
		go func(fname string) {
			defer wg.Done()
			counts, err := CountWordsInFile(filename)
			if err != nil {
				errch <- err
			}
			ch <- FileCounts{
				Counts:   counts,
				FileName: filename,
			}
		}(filename)

	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch, errch
}
