package logger

import (
	"io"
	"log"
	"os"
	"strings"
)

const logOutputPath = "debug.log"

var logger struct {
	*log.Logger
	level int
}

type multiWriter struct {
	FileWriter, StdoutWriter io.Writer

	n   int
	err error
}

func (c *multiWriter) Write(p []byte) (n int, err error) {
	c.n, c.err = c.StdoutWriter.Write(p)
	if err != nil {
		return 0, err
	}
	c.n, c.err = c.FileWriter.Write(p)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func NewLogger(level string) {
	var levelNum int
	switch strings.ToLower(level) {
	case "debug":
		levelNum = DEBUG
	case "info":
		levelNum = INFO
	case "warning", "warn":
		levelNum = INFO
	case "error":
		levelNum = ERROR
	case "fatal":
		levelNum = FATAL
	case "panic":
		levelNum = PANIC
	default:
		levelNum = INFO
	}

	if logger.Logger != nil {
		logger.Fatalf("NewLogger: 已经存在一个日志器了")
		return
	}
	f, err := os.OpenFile(logOutputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("Open logFile Error: %v", err)
		return
	}
	mw := &multiWriter{FileWriter: f, StdoutWriter: os.Stdout}

	logger.Logger = log.New(mw, "", log.LstdFlags|log.Lshortfile)
	logger.level = levelNum
}
