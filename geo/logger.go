package geo

import (
	"io"
	"log"
	"os"
)

// simple logger

const (
	prefix = "[go-geo] "
	flag   = log.LstdFlags | log.Lshortfile
)

var (
	info *log.Logger
	warn *log.Logger
	erro *log.Logger
)

func init() {
	var file *os.File
	var err error
	if errorOutputPath := os.Getenv("GO_GEO_ERROR_LOG_PATH"); errorOutputPath != "" {
		file, err = os.OpenFile(errorOutputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open error log file: ", err)
		}
	}

	info = log.New(os.Stdout, "INFO "+prefix, flag)
	warn = log.New(os.Stdout, "WARN "+prefix, flag)
	erro = log.New(io.MultiWriter(os.Stdout, file), "ERRO "+prefix, flag)
}
