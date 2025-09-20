package benchinput

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var nameRegex = regexp.MustCompile(`[^a-zA-Z0-9_ ]`)

func Bool(name string, default_ bool) bool {
	if nameRegex.MatchString(name) {
		panic("input name must not contain any special characters except _ or spaces")
	}
	str, err := getEnv(name)
	if err != nil {
		return default_
	}
	return str == "true"
}

func Int(name string, default_, min, max int) int {
	if nameRegex.MatchString(name) {
		panic("input name must not contain any special characters except _ or spaces")
	}
	str, err := getEnv(name)
	if err != nil {
		return default_
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return default_
	}
	return val
}

func getEnv(name string) (string, error) {
	if nameRegex.MatchString(name) {
		panic("input name must not contain any special characters except _ or spaces")
	}
	val, ok := os.LookupEnv(EnvVarName(name))
	if !ok {
		return "", fmt.Errorf("env var %s not set", EnvVarName(name))
	}
	return val, nil
}

func EnvVarName(name string) string {
	if nameRegex.MatchString(name) {
		panic("input name must not contain any special characters except _ or spaces")
	}
	return "BENCHSPOTTER_" + strings.ReplaceAll(name, " ", "_")
}
