package main

import (
	"fmt"
	"regexp"
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
	// 正则表达式：匹配以字母/数字开头、以字母/数字结尾，中间允许0+个横线的子串
	pattern := `[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?`
	re := regexp.MustCompile(pattern)

	input := "234-few、氛围服务"

	// 提取所有匹配项
	matches := re.FindAllString(input, -1)
	fmt.Println("提取结果:")
	for i, s := range matches {
		fmt.Printf("%d: %q\n", i+1, s)
	}
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
