package uuid

import "github.com/gofrs/uuid"

// Valid .
func Valid(v string) error {
	_, err := uuid.FromString(v)
	if err != nil {
		return err
	}
	return nil
}

// Generate .
func Generate() string {
	id, _ := uuid.NewV4()
	return id.String()
}
