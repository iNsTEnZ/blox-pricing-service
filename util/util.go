package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetEnv(variable string, defaultValue string) string {
	value := os.Getenv(variable)

	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func IntEnv(variable string, defaultValue int) int {
	value := os.Getenv(variable)

	if len(value) == 0 {
		return defaultValue
	}

	if num, err := strconv.Atoi(value); err != nil {
		log.Printf("error: unable to convert value of '%s' to a number", variable)
		return defaultValue
	} else {
		return num
	}
}

func GetListEnv(variable string, defaultValue string) []string {
	value := os.Getenv(variable)

	if len(value) == 0 {
		value = defaultValue
	}

	return strings.Split(value, ",")
}
