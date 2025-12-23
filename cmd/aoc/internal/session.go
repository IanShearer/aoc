package internal

import (
	"fmt"
	"os"
	"strings"
)

func LoadSessionCookie() (string, error) {
	content, err := os.ReadFile(".env")
	if err != nil {
		return "", fmt.Errorf("failed to read .env file: %w", err)
	}

	// Parse the .env file to find session=VALUE
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "session=") {
			return strings.TrimPrefix(line, "session="), nil
		}
	}

	return "", fmt.Errorf("session cookie not found in .env file")
}
