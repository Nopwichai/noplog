package noplog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var log_data string

func Log_info(service string, fn string, msg string) {
	log_data = log_data + time.Now().Format("15:04:05.0000") + "\tI\t" + service + "-" + fn + "\t" + msg + "\n"
}
func Log_warning(service string, fn string, msg string) {
	log_data = log_data + time.Now().Format("15:04:05.0000") + "\tW\t" + service + "-" + fn + "\t" + msg + "\n"
}
func Log_error(service string, fn string, msg string) {
	log_data = log_data + time.Now().Format("15:04:05.0000") + "\tE\t" + service + "-" + fn + "\t" + msg + "\n"
}

func Logs_write(log_dir string, time_delay int) {

	for {
		if len(log_data) > 0 {
			currentDate := time.Now().Format("2006-01-02_15")
			fname := log_dir + "Log-" + currentDate + ".log"
			file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			content := log_data
			if err != nil {
				fmt.Println("Error open file:", err)
			} else {
				_, err = file.WriteString(content)
				if err != nil {
					fmt.Println("Error writing to file:", err)
				} else {
					log_data = strings.Replace(log_data, content, "", -1)
				}
			}
		}
		time.Sleep(time.Duration(time_delay) * 1000 * time.Millisecond)
	}
}
