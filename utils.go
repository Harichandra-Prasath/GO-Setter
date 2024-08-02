package main

import (
	"fmt"
	"strings"
)

func getMod(origin string) (string, error) {

	if strings.HasPrefix(origin, "https://") {
		origin = strings.TrimPrefix(origin, "https://")
		origin = strings.TrimSuffix(origin, ".git")
		return origin, nil
	} else if strings.HasPrefix(origin, "git@") {
		origin = strings.TrimPrefix(origin, "git@")
		origin = strings.TrimSuffix(origin, ".git")
		origin = strings.Replace(origin, ":", "/", 1)
		return origin, nil
	}
	return "", fmt.Errorf("sorry, cannot extract from origin. Please enter manually")

}

func updateBool(input string, field *bool) {
	switch input {
	case "yes":
		*(field) = true
	case "no":
		*(field) = false
	default:
		panic("Unknown Input")
	}
}

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
