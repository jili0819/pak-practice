package main

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func main() {

	resp, err := http.Head("https://wos.develop.meetwhale.com/d_I2PJa7-xFhAV738vXl-")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	parsedURL, _ := url.Parse(resp.Request.URL.String())

	// 解析Content-Disposition头
	fmt.Println(path.Base(parsedURL.Path))
	fmt.Println(strings.TrimPrefix(path.Ext(parsedURL.Path), "."))
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
