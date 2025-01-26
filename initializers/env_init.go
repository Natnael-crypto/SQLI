package initializers

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
		log.Fatalf("could not open env file: %v", err)
	}

	envFile := bytes.NewReader(envBytes)
	scanner := bufio.NewScanner(envFile)

	for scanner.Scan() {
		enVar := strings.SplitN(scanner.Text(), "=", 2) // Use SplitN to avoid index out-of-range errors
		if len(enVar) != 2 {
			log.Printf("invalid environment variable format: %s", scanner.Text())
			continue
		}

		err := os.Setenv(strings.TrimSpace(enVar[0]), strings.TrimSpace(enVar[1]))
		if err != nil {
			log.Printf("failed to set environment variable %s: %v", enVar[0], err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading env file: %v", err)
	}
}
