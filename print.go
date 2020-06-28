package main

import (
	"fmt"
	"time"
)

func MyPrintln(args ...interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), args)
}

func MyPrint(args ...interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), args)
}
