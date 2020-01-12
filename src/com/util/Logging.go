package util

import (
	"log"
	"os"
)

func LogToFile(path string, message string) {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println(message)

}
