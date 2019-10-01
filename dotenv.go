package dotenv

import (
	"bufio"
	"os"
	"strings"
)

// fileName is a default filename.
const fileName string = ".env"

// LoadEnv load env variables from .env.
func LoadEnv(filenames ...string) error {
	for _, filename := range setFilename(filenames) {
		lines, err := readFile(filename)
		if err != nil {
			return err
		}

		setEnv(lines)
	}

	return nil
}

// setFilename set default filename if filenames has no strings.
func setFilename(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{fileName}
	}

	return filenames
}

// readFile read a .env.
func readFile(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// setEnv set env.
func setEnv(lines []string) {
	for _, l := range lines {
		pair := strings.Split(l, "=")
		os.Setenv(pair[0], pair[1])
	}
}
