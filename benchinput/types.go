package benchinput

import (
	"fmt"
	"os"
	"strconv"

	"github.com/iancoleman/strcase"
)

func Bool(name string, default_ bool) bool {
	str, err := getEnv(name)
	if err != nil {
		return default_
	}
	return str == "true"
}

func Int(name string, default_, min, max int) int {
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
	val, ok := os.LookupEnv(EnvVarName(name))
	if !ok {
		return "", fmt.Errorf("env var %s not set", EnvVarName(name))
	}
	return val, nil
}

func EnvVarName(name string) string {
	return "BENCHSPOTTER_" + strcase.ToScreamingSnake(name)
}
