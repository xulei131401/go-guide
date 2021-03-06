package s1

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// byte.bufer,stringSearch.buffer,+,fmt.Sprintf,strings.Join
// 结论：使用 buffer 连接字符串的性能比 += 要好很多。
const numbers = 100

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		for i := 0; i < numbers; i++ {
			var s string
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		var builder strings.Builder
		for i := 0; i < numbers; i++ {
			builder.WriteString(strconv.Itoa(i))
		}
	}
	b.StopTimer()
}

func BenchmarkSytesBuf(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		var buf bytes.Buffer
		for i := 0; i < numbers; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}
	}
	b.StopTimer()
}
