package dup2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Exercise1p4 根据文件名读取文件内容并打印有重复行的文件名
func Exercise1p4() {
	counts := make(map[string]int)
	files := os.Args[1:]
	filesHasDups := []string{}

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			if countLines(f, counts) {
				filesHasDups = append(filesHasDups, arg)
			}

			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	fmt.Printf("Files that has dup lines %v", strings.Join(filesHasDups, ","))
}

func countLines(f *os.File, counts map[string]int) (hasDupLines bool) {
	input := bufio.NewScanner(f)
	currentFileCount := map[string]int{}
	for input.Scan() {
		counts[input.Text()]++
		currentFileCount[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()

	// Determine file which has dup lines.
	for _, n := range currentFileCount {
		if n > 1 {
			hasDupLines = true
			break
		}
	}

	return
}
