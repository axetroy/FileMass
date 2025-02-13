package file_mass

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/schollz/progressbar/v3"
)

// 随机生成指定范围内的文件内容
func generateRandomContent(minKB, maxKB int) string {
	// 生成随机内容的大小
	size := rand.Intn(maxKB-minKB+1) + minKB
	// 生成内容，按字节数构建
	content := make([]byte, size*1024)
	for i := 0; i < len(content); i++ {
		content[i] = byte(rand.Intn(94) + 32) // 随机字符
	}
	return string(content)
}

// 创建文件并写入随机内容
func createFile(dirPath, fileName string, wg *sync.WaitGroup, bar *progressbar.ProgressBar) {
	defer wg.Done()

	// 确保目录存在
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Println("目录创建失败:", err)
		return
	}

	// 文件路径
	filePath := filepath.Join(dirPath, fileName)

	// 生成随机内容
	content := generateRandomContent(1, 100)

	// 创建文件并写入内容
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("文件写入失败:", err)
		return
	}

	// 更新进度条
	bar.Add(1)
}

func Mass() {

	// 输出根目录
	baseDir := "./output"

	// 创建进度条
	totalFiles := 1000 * 100 // 1000 个目录，每个目录 100 个文件
	bar := progressbar.NewOptions(totalFiles,
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetDescription("正在生成文件"),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionShowCount(),
	)

	// 使用 WaitGroup 来等待所有文件生成完毕
	var wg sync.WaitGroup

	// 并发生成文件的限制
	concurrencyLimit := 10
	sem := make(chan struct{}, concurrencyLimit) // 限制并发数

	// 循环创建文件
	for i := 1; i <= 1000; i++ {
		dirPath := filepath.Join(baseDir, strconv.Itoa(i))
		for j := 1; j <= 100; j++ {
			fileName := strconv.Itoa(j) + ".txt"

			// 使用协程并发创建文件
			wg.Add(1)
			sem <- struct{}{} // 获取一个信号量
			go func() {
				defer func() { <-sem }() // 释放信号量
				createFile(dirPath, fileName, &wg, bar)
			}()
		}
	}

	// 等待所有文件完成创建
	wg.Wait()

	// 关闭进度条
	bar.Finish()

	fmt.Println("目录和文件创建完成！")
}
