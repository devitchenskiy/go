// echo вывод аргументов командной строки
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + fmt.Sprint(i) + ". " + os.Args[i]
		sep = "\n"
		//fmt.Println(fmt.Sprint(i) + ". " + os.Args[i])
	}
	fmt.Println(s)
}
