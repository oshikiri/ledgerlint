package main

import (
	"fmt"
	"io/ioutil"
)

func readFileContent(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ioutil.ReadFile failed: %v, filePath='%v'\n", err, filePath)
		return "", err
	}
	fileContent := string(bytes)
	return fileContent, nil
}
