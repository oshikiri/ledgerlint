package main

import (
	"fmt"
	"io/ioutil"
)

// FIXME?: The current implementation requires O(n)
//         (but the array is usually small)
func contains(array []string, query string) bool {
	for _, item := range array {
		if query == item {
			return true
		}
	}
	return false
}

func readFileContent(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		// FIXME: show usage if filePath is empty
		fmt.Printf("ioutil.ReadFile failed: %v, filePath='%v'\n", err, filePath)
		return "", err
	}
	fileContent := string(bytes)
	return fileContent, nil
}
