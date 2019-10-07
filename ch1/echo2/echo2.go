// Package echo2 os.Args 是通过空格分隔不同参数的.
package echo2

import (
	"fmt"
	"os"
)

// Exercise1 print commandline command
func Exercise1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// 打印命令自身
func exercise1p1() {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// 打印命令 index + 命令.
func exercise1p2() {
	for i, arg := range os.Args {
		fmt.Println(fmt.Sprintf("%d: %s", i, arg))
	}
}
