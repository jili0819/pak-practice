package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

//这个示例代码使用了Gin框架实现了文件断点续传上传文件的功能，具体实现过程如下：
//定义了两个结构体UploadInfo和ChunkInfo，分别用于存储上传文件信息和上传文件分片信息。
//实现了上传文件分片处理函数UploadChunk，该函数接收上传文件分片数据，并将其保存到本地文件系统中。
//实现了合并上传文件分片处理函数MergeChunks，该函数检查上传文件分片是否全部上传完成，如果是，则将所有分片合并为一个完整的文件，并删除上传文件分片。
//初始化上传文件目录UploadDir，用于存储上传文件和上传文件分片。
//注册上传文件分片处理路由/upload/chunk，用于接收上传文件分片数据。
//注册合并上传文件分片处理路由/merge/chunks，用于合并上传文件分片。
//启动Gin引擎，监听HTTP请求。

const (
	// 文件保存目录
	UploadDir = "./uploads"
)

// 上传文件信息
type UploadInfo struct {
	FileName string // 文件名
	ChunkNum int    // 总分片数
	ChunkIdx int    // 当前分片索引
}

// 上传文件分片信息
type ChunkInfo struct {
	ChunkIdx int    // 分片索引
	Chunk    []byte // 分片数据
}

// 上传文件分片处理函数
func UploadChunk(c *gin.Context) {
	// 获取上传文件信息
	fileName := c.PostForm("fileName")
	//chunkNum, _ := strconv.Atoi(c.PostForm("chunkNum"))
	chunkIdx, _ := strconv.Atoi(c.PostForm("chunkIdx"))

	// 获取上传文件分片信息
	chunk := c.Request.Body
	chunkData, err := io.ReadAll(chunk)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "读取分片数据失败",
		})
		return
	}

	// 保存上传文件分片
	chunkPath := filepath.Join(UploadDir, fileName, fmt.Sprintf("%d", chunkIdx))
	err = os.WriteFile(chunkPath, chunkData, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "保存分片数据失败",
		})
		return
	}

	// 返回上传文件分片处理结果
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "上传分片成功",
	})
}

// 合并上传文件分片处理函数
func MergeChunks(c *gin.Context) {
	// 获取上传文件信息
	fileName := c.Query("fileName")
	chunkNum, _ := strconv.Atoi(c.Query("chunkNum"))

	// 检查上传文件分片是否全部上传完成
	chunkDir := filepath.Join(UploadDir, fileName)
	chunkFiles, err := os.ReadDir(chunkDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "读取分片文件列表失败",
		})
		return
	}
	if len(chunkFiles) != chunkNum {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "分片文件数量不正确",
		})
		return
	}

	// 合并上传文件分片
	file, err := os.Create(filepath.Join(UploadDir, fileName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "创建文件失败",
		})
		return
	}
	defer file.Close()

	for i := 0; i < chunkNum; i++ {
		chunkPath := filepath.Join(chunkDir, fmt.Sprintf("%d", i))
		chunkData, err := os.ReadFile(chunkPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "读取分片数据失败",
			})
			return
		}

		_, err = file.Write(chunkData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "写入文件数据失败",
			})
			return
		}
	}

	// 删除上传文件分片
	err = os.RemoveAll(chunkDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "删除分片文件失败",
		})
		return
	}

	// 返回合并上传文件分片处理结果
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "上传文件成功",
	})
}

// 初始化上传文件目录
func InitUploadDir() {
	err := os.MkdirAll(UploadDir, 0755)
	if err != nil {
		panic(err)
	}
}

func main() {
	// 初始化上传文件目录
	InitUploadDir()

	// 创建Gin引擎
	r := gin.Default()

	// 注册上传文件分片处理路由
	r.POST("/upload/chunk", UploadChunk)

	// 注册合并上传文件分片处理路由
	r.GET("/merge/chunks", MergeChunks)

	// 启动Gin引擎
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
