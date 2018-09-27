// rotate
package main

import (
	"fmt"
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	rotate(&a, 3)
	fmt.Println(a)
}
func rotate(s *[10]int, n int) {
	t := make([]int, n-1)
	copy(t, s[0:n-1])
	copy(s[:], s[n-1:])
	copy(s[len(s)-n+1:], t)

}
