package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	count_files := make([]string, 1)
	files := os.Args[1:]
	if len(files) == 0 {
		count_files = countLines(os.Stdin, counts, count_files)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			count_files = countLines(f, counts, count_files)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	for _, file_name := range count_files {
		fmt.Printf("%s, ", file_name)
	}
	fmt.Printf("%s\n", "")
}

func countLines(file *os.File, counts map[string]int, count_files []string) []string {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			count_files = append(count_files, file.Name())
		}
	}
	return count_files
}
