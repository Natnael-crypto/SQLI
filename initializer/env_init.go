package initializer

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"strings"
)

func LoadEnv(fileSystem fs.FS) {
	envFile, err := fileSystem.Open(".env")
	if err != nil {
		log.Fatalf("could not open env file, %v", err)
	}
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		enVar := strings.Split(scanner.Text(), "=")
		os.Setenv(enVar[0], enVar[1])
	}
}
