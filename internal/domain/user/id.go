package user

import (
	"github.com/kskumgk63/containized-firestore/pkg/uuid"
)

// ID .
type ID struct {
	v string
}

// GenerateID .
func GenerateID() *ID {
	return &ID{uuid.Generate()}
}

// NewIDFromString .
func NewIDFromString(v string) (*ID, error) {
	if err := uuid.Valid(v); err != nil {
		return nil, err
	}
	return &ID{v}, nil
}

// String .
func (id ID) String() string {
	return id.v
}

// IDs .
type IDs []ID

// String .
func (ids IDs) String() []string {
	result := make([]string, len(ids))
	for i, id := range ids {
		result[i] = id.String()
	}
	return result
}
