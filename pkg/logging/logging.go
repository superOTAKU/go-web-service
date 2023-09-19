package logging

import (
	"go-web-service/pkg/config"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
)

var (
	Logger *log.Logger
)

func InitLogger(config *config.LoggingConfig) error {
	dirPath := filepath.Dir(config.Filepath)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}
	file, err := os.OpenFile(config.Filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	writer := io.MultiWriter(os.Stdout, file)
	Logger = log.NewWithOptions(writer, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
	})
	return nil
}
