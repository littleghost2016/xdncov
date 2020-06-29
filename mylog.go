package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var (
	LogPath = "xdncov.log"
)

func init() {
	// 代码参考自 https://github.com/sirupsen/logrus
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, _ := os.OpenFile(LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(io.MultiWriter(file, os.Stdout))

	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)
}

func StandardLog(ID int, message string) {
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"), args)
	log.WithFields(log.Fields{
		"ID":      ID,
		"message": message,
	}).Info("")

}
