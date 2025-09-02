package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func TruncateTo8MBFast(s string) string {
	maxBytes := 8 * 2 // 16b
	//maxBytes := 8 * 1024 * 1024
	if len(s) <= maxBytes {
		return s
	}

	truncated := s[:maxBytes]
	// 处理尾部可能被截断的多字节字符
	for i := 0; i < utf8.MaxRune; i++ {
		if len(truncated)-i <= 0 {
			break
		}
		if utf8.FullRune([]byte(truncated[:len(truncated)-i])) {
			return truncated[:len(truncated)-i]
		}
	}
	return ""
}

func main() {
	ch := make(chan struct{}, 2)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		go func(s int) {
			fmt.Println(fmt.Sprintf("hello world %d", s))
			time.Sleep(2 * time.Second)
			<-ch
		}(i)
	}

	// 使用sts token创建 imm client
	/*
	 */
	//oss.New("oss-cn-hangzhou.aliyuncs.com", "LTAI5tMCNyg9zqxhAZzn9gQn", "<KEY>")

}

func SplitRuneChunks(s string, chunkSize int) [][]rune {
	runes := []rune(s)
	if chunkSize <= 0 || len(runes) == 0 {
		return nil
	}

	// 预分配内存优化
	total := (len(runes) + chunkSize - 1) / chunkSize
	chunks := make([][]rune, 0, total)
	for i := 0; i < len(runes); i += chunkSize {
		end := i + chunkSize
		if end > len(runes) {
			end = len(runes)
		}
		// 创建独立副本避免共享底层数组
		chunk := make([]rune, end-i)
		copy(chunk, runes[i:end])
		chunks = append(chunks, chunk)
	}
	return chunks
}
