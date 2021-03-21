package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Generate generates the logs with given options
func Generate(option *Option) error {
	var (
		created  = time.Now()
		interval time.Duration
		delay    time.Duration
	)

	delay = 30
	interval = 10

	logFileName := option.Output
	writer, err := NewWriter("log", logFileName)
	if err != nil {
		return err
	}

	// Generates the logs until the certain number of lines is reached
	for line := 0; line < option.Number; line++ {
		time.Sleep(delay)
		log := NewCommonLogFormat(created)
		_, _ = writer.Write([]byte(log + "\n"))
		created = created.Add(interval)
	}
	_ = writer.Close()
	fmt.Println(logFileName, "is created.")

	return nil
}

// NewWriter returns a closeable writer corresponding to given log type
func NewWriter(logType string, logFileName string) (io.WriteCloser, error) {
	switch logType {
	case "stdout":
		return os.Stdout, nil
	case "log":
		logFile, err := os.Create(logFileName)
		if err != nil {
			return nil, err
		}
		return logFile, nil
	case "gz":
		logFile, err := os.Create(logFileName)
		if err != nil {
			return nil, err
		}
		return gzip.NewWriter(logFile), nil
	default:
		return nil, nil
	}
}

// NewSplitFileName creates a new file path with split count
func NewSplitFileName(path string, count int) string {
	logFileNameExt := filepath.Ext(path)
	pathWithoutExt := strings.TrimSuffix(path, logFileNameExt)
	return pathWithoutExt + strconv.Itoa(count) + logFileNameExt
}
