package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/nickpancakes/aoc-2017-01/internal"
)

func main() {
	captcha, err := readInputFile()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Captcha:\n%s\n", captcha)

	solved := internal.SolveCaptcha(captcha)

	fmt.Printf("\nSolved:\n%s\n", solved)
}

func readInputFile() (string, error) {
	_, callerFName, _, ok := runtime.Caller(0) //nolint:dogsled
	if !ok {
		return "", fmt.Errorf("Failed to get the caller's filename")
	}

	absPath := path.Join(path.Dir(callerFName), "input.txt")

	data, err := os.ReadFile(absPath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
