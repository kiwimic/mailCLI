package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func logOperation(from string, to string, attachment string, subject string, gate string, port int, config string, filename string) {

	t := time.Now().Format("2006-01-02 15:04:05")

	if filename == "" {
		filename = "logs/logs.txt"
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	logRecord := t + ";" + from + ";" + to + ";" + attachment + ";" + subject + ";" + gate + ";" + strconv.Itoa(port) + ";" + config + "\n"

	fmt.Fprintf(f, "%s", logRecord)

	f.Close()
}
