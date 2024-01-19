// echo выводит аргументы командной строки
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Ex3
func main() {
	s, sep := "", ""
	start_time := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	end_time := time.Now()
	elapsed_time := end_time.Sub(start_time)
	fmt.Println("\n", elapsed_time)
	fmt.Println("-----------")
	start_time = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end_time = time.Now()
	fmt.Println("\n", end_time.Sub(start_time))
}

// Ex1
// func main() {
// 	var s, sep string
// 	for i := 0; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	var s, sep string
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// Ex2

// func main() {
// 	s := ""
// 	for i, arg := range os.Args[1:] {
// 		s += strconv.Itoa(i) + " " + arg + "\n"
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// }
