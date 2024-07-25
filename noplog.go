package noplog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var log_data_info string
var log_data_warning string
var log_data_error string

func Log_info(service string, msg string) {
	log_data_info = log_data_info + time.Now().Format("15:04:05.0000") + "\t" + service + "-" + msg + "\n"
}
func Log_warning(service string, msg string) {
	log_data_warning = log_data_warning + time.Now().Format("15:04:05.0000") + "\t" + service + "\t" + msg + "\n"
}
func Log_error(service string, msg string) {
	log_data_error = log_data_error + time.Now().Format("15:04:05.0000") + "\t" + service + "\t" + msg + "\n"
}

func Logs_Write(log_dir string) {

	var time_delay int = 5

	for {
		if len(log_data_info) > 0 {
			currentDate := time.Now().Format("20060102")
			fname := log_dir + "Log-" + currentDate + "-info.log"
			file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			content := log_data_info
			if err != nil {
				fmt.Println("Error open file:", err)
			} else {
				_, err = file.WriteString(content)
				if err != nil {
					fmt.Println("Error writing to file:", err)
				} else {
					log_data_info = strings.Replace(log_data_info, content, "", -1)
				}
			}
		}

		if len(log_data_warning) > 0 {
			currentDate := time.Now().Format("20060102")
			fname := log_dir + "Log-" + currentDate + "-warning.log"
			file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			content := log_data_warning
			if err != nil {
				fmt.Println("Error open file:", err)
			} else {
				_, err = file.WriteString(content)
				if err != nil {
					fmt.Println("Error writing to file:", err)
				} else {
					log_data_warning = strings.Replace(log_data_warning, content, "", -1)
				}
			}
		}

		if len(log_data_error) > 0 {
			currentDate := time.Now().Format("20060102")
			fname := log_dir + "Log-" + currentDate + "-error.log"
			file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			content := log_data_error
			if err != nil {
				fmt.Println("Error open file:", err)
			} else {
				_, err = file.WriteString(content)
				if err != nil {
					fmt.Println("Error writing to file:", err)
				} else {
					log_data_error = strings.Replace(log_data_error, content, "", -1)
				}
			}
		}
		time.Sleep(time.Duration(time_delay) * 1000 * time.Millisecond)
	}
}
