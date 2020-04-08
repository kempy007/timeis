package main

import (
	"io"
	"log"
	"os"
	"time"
)

type writer struct {
	io.Writer
	timeFormat string
}

func (w writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.timeFormat)), b...))
}

func main(){
	for {
		logger := log.New(&writer{os.Stdout, "2006/01/02 15:04:05 "}, "[info]", 0)
		logger.Println("( The Time is, " + time.Now().Format("15:04:05") + " )")
		<-time.After(time.Second * 10)
	}
}