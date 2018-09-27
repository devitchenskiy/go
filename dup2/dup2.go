// dup2
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {

			fmt.Println(arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v/n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
func countLines(f *os.File, counts map[string]int, s string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		//line := s + " " + input.Text()
		//counts[line] = counts[line] + 1
		counts[s+" "+input.Text()]++

	}
}
