package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// env vars
var (
	port      = os.Getenv("PORT")
	projectID = os.Getenv("PROJECT_ID")
)

// Port .
func Port() (int, error) {
	p, err := strconv.Atoi(port)
	if err != nil {
		return 0, fmt.Errorf("Failed to get PORT: %v", err)
	}
	return p, nil
}

// ProjectID .
func ProjectID() (string, error) {
	if projectID == "" {
		return "", errors.New("PROJECT_ID must be set")
	}
	return projectID, nil
}
