package utils

import (
	"log"
	"strconv"
	"strings"
)

func TypeCheck(vType string, value string) {
	switch vType {
	case "int":
		if strings.HasPrefix(value, "\"") || strings.HasSuffix(value, "\"") {
			log.Fatal("Invalid int value: ", value)
		}
		if _, err := strconv.Atoi(value); err != nil {
			log.Fatal("Invalid int value: ", value)
		}
	case "str":
		if !strings.HasPrefix(value, "\"") || !strings.HasSuffix(value, "\"") {
			log.Fatal("Invalid str value: ", value)
		}
	}

}
