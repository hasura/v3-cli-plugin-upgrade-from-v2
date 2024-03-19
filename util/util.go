/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package util

import (
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
