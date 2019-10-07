package echo2_test

import (
	"os"
	"strings"
	"testing"

	"github.com/raven-chen/gopl/ch1/echo2"
)

// Result of this benchmark group.  自己循环拼接字符串比 strings.Join 更快. why?
// BenchmarkStringsJoin-8   	20000000	        57.0 ns/op
// BenchmarkExercise1-8     	100000000	        11.6 ns/op

// BenchmarkExercise1AndStringsJoin 验证 exercise1 与 strings.Join 的效率.  benchmark and time package needed
func BenchmarkStringsJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.Join(os.Args, "")
	}
}

// Have to comment out Exercise1 output. otherwise the result cannot be seen.
func BenchmarkExercise1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		echo2.Exercise1()
	}
}
