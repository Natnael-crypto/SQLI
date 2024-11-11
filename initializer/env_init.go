package initializer

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

func LoadEnv() {
	envBytes, err := os.ReadFile(".env")
	if err != nil {
		log.Fatalf("could not open env file, %v", err)
	}
	envFile := bytes.NewReader(envBytes)
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		enVar := strings.Split(scanner.Text(), "=")
		os.Setenv(enVar[0], enVar[1])
	}
}
