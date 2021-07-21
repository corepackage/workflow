package util

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"path/filepath"
	"regexp"
)

// Validating input file
func ValidateFile(file string) error {
	if file == "" {
		return errors.New("File name is empty")
	} else {
		ext := filepath.Ext(file)
		if ext != ".json" && ext != ".yaml" && ext != ".yml" {
			return errors.New("Invalid file: " + file)
		}
	}
	return nil
}

// FindInArray : to find data from array
func FindInArray(val string, arr []string) (int, bool) {
	for k, v := range arr {
		if v == val {
			return k, true
		}
	}
	return -1, false
}

// ParseData : to parse io reader data to map
func ParseData(body io.ReadCloser) (map[string]interface{}, error) {
	p := make(map[string]interface{})
	err := json.NewDecoder(body).Decode(&p)
	if err != nil {
		log.Println("ParseData: Error parsing body")
		return nil, err
	}
	return p, nil
}

// FindMatchStr : to find substrings matching a regex
func FindMatchStr(regex string, str string) []string {
	re := regexp.MustCompile(regex)
	return re.FindAllString(str, -1)
}
