package fileops

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func readline() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}

// WriteToFile : To write file at a particular path
func WriteToFile(data, file string) error {

	// Checking dir
	if _, err := os.Stat(file); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(file), 0700)
	}
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("error opening file:", err)
	}
	defer f.Close()
	_, err = f.Write([]byte(data))
	return err
}
func GetFileName(path string) string {
	return strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
}

// ReadFromFile : To read file of specified filename
func ReadFromFile(file string) ([]byte, error) {
	data, readErr := ioutil.ReadFile(file)

	return data, readErr
}
