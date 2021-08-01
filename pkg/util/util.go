package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
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

// FindValue : to find value from map or array using nested keys
func FindValue(bodyJson interface{}, keys []string) (interface{}, error) {
	itrMap := bodyJson
	for i := 0; i < len(keys); i++ {
		index, err := strconv.Atoi(keys[i])
		if err != nil {
			mapObj, ok := itrMap.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid key %v for the input body", keys[i])
			}
			itrMap = mapObj[keys[i]]
		} else {
			arrObj, ok := itrMap.([]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid key %v for the input body", keys[i])
			}
			itrMap = arrObj[index]
		}
	}
	return itrMap, nil

}

// ToTime : to convert string to time
func ToTime(str string) (time.Duration, error) {
	t, err := time.ParseDuration(str)
	if err != nil {
		return 0, fmt.Errorf("time is not valid")
	}
	return t, nil
}
