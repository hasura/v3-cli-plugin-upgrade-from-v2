/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func GetenvBool(name string) bool {
	valU := os.Getenv(name)
	val := strings.ToLower(valU)
	return val != "" && val != "false" && val != "f" && val != "no" && val != "n"
}

func GetenvStringWithDefault(name string, default_value string) string {
	val := os.Getenv(name)
	if val == "" {
		return default_value
	} else {
		return val
	}
}

// Panics if there are any errors
func ReadJSON(filePath string) map[string]interface{} {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error reading file %s: %s", filePath, err))
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling JSON %s: %s", filePath, err))
	}

	return jsonData
}

func FormatJSON(o interface{}) string {
	jsonData, err := json.Marshal(o)
	if err != nil {
		panic(fmt.Sprintf("Error marshalling JSON: %s", err))
	}
	return string(jsonData)
}
