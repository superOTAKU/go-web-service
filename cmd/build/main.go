package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const (
	targetDir = "dist"
	confDir   = "conf"
	appName   = "web-service"
)

// 构建web应用
func main() {
	// 如果dist已存在，先删除
	_, err := os.Stat(targetDir)
	if !os.IsNotExist(err) {
		if err := os.RemoveAll(targetDir); err != nil {
			log.Fatalf("delete dist error: %v\n", err)
		} else {
			log.Println("dist removed")
		}
	} else {
		log.Println("dist not exists")
	}
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		log.Fatalf("Make dist error: %v\n", err)
	}
	// 执行构建
	osName := runtime.GOOS
	execFileName := appName
	if osName == "windows" {
		execFileName += ".exe"
	}
	buildCmd := exec.Command("go", "build", "-o", filepath.Join(targetDir, execFileName), "./cmd/web/main.go")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		log.Fatalf("Build error: %v\n", err)
		return
	}
	// 复制配置文件
	copyDir(confDir, fmt.Sprintf("%s/%s", targetDir, confDir))
}

func copyDir(sourceDir, destDir string) error {
	_, err := os.Stat(destDir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return err
		}
	}
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		toPath := filepath.Join(destDir, path[len(sourceDir):])
		if info.IsDir() {
			if err := os.MkdirAll(toPath, info.Mode()); err != nil {
				return err
			}
		} else {
			if err := copyFile(path, toPath); err != nil {
				return err
			}
		}
		return nil
	})
}

// 复制文件
func copyFile(sourcePath, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
