// wordfreq
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	fmt.Println(os.Args[0])
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		freq[word]++
	}
	for word, n := range freq {
		fmt.Printf("%-50s %d\n", word, n)
	}
}
