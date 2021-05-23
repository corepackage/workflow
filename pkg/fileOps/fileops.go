package fileops

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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
	writeErr := ioutil.WriteFile(file, []byte(data), 777)
	return writeErr
}
func GetFileName(path string) string {
	return strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
}

// ReadFromFile : To read file of specified filename
func ReadFromFile(file string) ([]byte, error) {
	data, readErr := ioutil.ReadFile(file)

	return data, readErr
}
