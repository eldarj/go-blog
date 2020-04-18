package view

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ViewData(viewName string, data map[string]string) string {
	readFile, err := ioutil.ReadFile("./www/views/" + viewName + ".html")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	result := string(readFile)
	for k, v := range data {
		result = strings.Replace(result, "{{" + k + "}}", v, -1)
	}
	fmt.Println("Contents of file:", result)
	return result
}