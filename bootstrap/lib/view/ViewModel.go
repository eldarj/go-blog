package view

import (
	"fmt"
	m "go-blog/model"
	"io/ioutil"
	"reflect"
	"strings"
)

func ViewModel(viewName string, model m.BlogPageModel) string {
	readFile, err := ioutil.ReadFile("./www/" + viewName + ".html")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	result := string(readFile)
	data := reflect.ValueOf(model)
	dataTypes := data.Type()

	for i := 0; i< data.NumField(); i++ {
		fieldName := dataTypes.Field(i).Name
		fieldValue := data.Field(i).Interface()
		result = strings.Replace(result, "{{model." + fieldName + "}}", fieldValue.(string), -1)
	}

	return result
}