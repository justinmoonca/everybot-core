// Package file 文件操作辅助函数
package file

import (
	"fmt"
	"github.com/justinmoonca/everybot-core/pkg/app"
	"github.com/justinmoonca/everybot-core/pkg/helpers"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// Put 将数据存入文件
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func SaveUploadFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	// 确保目录存在，不存在创建
	dirName := fmt.Sprintf("storage/uploads/files/%s/", app.TimenowInTimezone().Format("2006/01/02"))
	helpers.IsNotExistMkDir(dirName)

	// 保存文件
	fileName := randomNameFromUploadFile(file)
	filePath := dirName + fileName
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}
