package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var directory = "log"

func init() {
	if dir := os.Getenv("DIR_LOG"); dir != "" {
		directory = dir
	}
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, 0755)
	}
}

func print(msg string) {
	fmt.Print("\n========\n\n")
	fmt.Printf(msg)
	fmt.Print("\n========\n\n")
}

func save(file, data string) {
	file += "_" + time.Now().Format("2006_01_02")
	f, err := os.OpenFile(directory+"/"+file+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err = f.WriteString(data); err != nil {
		log.Fatal(err)
	}
}

// Error the errors are recorded in the log/error.log file
func Error(tag string, err error) {
	if err != nil {
		msg := fmt.Sprintf("[%s] ERROR/%s: %s\n",
			time.Now().Format("2006/01/02 15:04:05"), tag, err.Error())
		print(msg)
		save("error", msg)
	}
}

// ErrorFatal the errors are recorded in the log/error.log file and exit
func ErrorFatal(tag string, err error) {
	if err != nil {
		tag = "[FATAL] " + tag
		Error(tag, err)
		os.Exit(1)
	}
}
