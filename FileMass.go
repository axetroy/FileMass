package file_mass

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"

	progressBar "github.com/schollz/progressbar/v3"
)

// 随机生成指定范围内的文件内容
func generateRandomContent(minKB, maxKB int) string {
	// 生成随机内容的大小
	size := rand.Intn(maxKB-minKB+1) + minKB
	// 生成内容，按字节数构建
	content := make([]byte, size*1024)
	for i := range content {
		content[i] = byte(rand.Intn(94) + 32) // 随机字符
	}
	return string(content)
}

// 创建文件并写入随机内容
func createFile(dirPath, fileName string, minSize, maxSize int, wg *sync.WaitGroup, bar *progressBar.ProgressBar) {
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
	content := generateRandomContent(minSize, maxSize)

	// 创建文件并写入内容
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("文件写入失败:", err)
		return
	}

	// 更新进度条
	_ = bar.Add(1)
}

type Config struct {
	Concurrence int
	Output      string
	Depth       int
	MinSize     int
	MaxSize     int
	Dirs        int
	Files       int
	Clean       bool
}

func Mass(config Config) error {

	// 输出根目录
	baseDir := config.Output

	if config.Clean {
		// 清空目录
		err := os.RemoveAll(baseDir)
		if err != nil {
			return fmt.Errorf("清空目录失败: %w", err)
		}
	}

	// 创建进度条
	totalFiles := config.Dirs * config.Files * intPow(config.Depth, config.Dirs)
	bar := progressBar.NewOptions(totalFiles,
		progressBar.OptionSetRenderBlankState(true),
		progressBar.OptionSetWidth(50),
		progressBar.OptionSetDescription("正在生成文件"),
		progressBar.OptionSetPredictTime(true),
		progressBar.OptionShowCount(),
	)

	// 使用 WaitGroup 来等待所有文件生成完毕
	var wg sync.WaitGroup

	// 并发生成文件的限制
	sem := make(chan struct{}, config.Concurrence) // 限制并发数

	// 递归创建目录和文件
	var createFilesRecursively func(currentDir string, currentDepth int)
	createFilesRecursively = func(currentDir string, currentDepth int) {
		if currentDepth > config.Depth {
			return
		}

		for i := 1; i <= config.Dirs; i++ {
			dirPath := filepath.Join(currentDir, fmt.Sprintf("%d-%d", currentDepth, i))
			for j := 1; j <= config.Files; j++ {
				fileName := fmt.Sprintf("%d-%d.txt", currentDepth, j)

				// 使用协程并发创建文件
				wg.Add(1)
				sem <- struct{}{} // 获取一个信号量
				go func(dirPath, fileName string) {
					defer func() { <-sem }() // 释放信号量
					createFile(dirPath, fileName, config.MinSize, config.MaxSize, &wg, bar)
				}(dirPath, fileName)
			}
			createFilesRecursively(dirPath, currentDepth+1)
		}
	}

	// 开始递归创建文件
	createFilesRecursively(baseDir, 1)

	// 等待所有文件完成创建
	wg.Wait()

	// 关闭进度条
	if err := bar.Finish(); err != nil {
		return err
	}

	return nil
}

// 计算整数的幂
func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
