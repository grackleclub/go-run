package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	timeoutDefault int    = 60
	timeoutVar     string = "GORUN_TIMEOUT"
)

var (
	testFilePath = path.Join("example", "static", "test.txt")
)

func main() {
	timeoutEnv, ok := os.LookupEnv(timeoutVar)
	timeout, err := strconv.Atoi(timeoutEnv)
	if err != nil || !ok {
		timeout = timeoutDefault
	}
	fmt.Printf(
		"👋 I'm a sample go program that runs for %s=%d seconds!\n",
		timeoutVar, timeout,
	)
	fmt.Printf("👀 watching file: %s\n", testFilePath)

	for i := range timeout {
		b, err := os.ReadFile(testFilePath)
		if err != nil {
			panic(fmt.Errorf("read %q: %w", testFilePath, err))
		}
		fmt.Printf("(%v) 📁 test.txt: %s\n", i+1, strings.TrimSpace(string(b)))
		time.Sleep(1 * time.Second)
	}
	os.Exit(4)
}
