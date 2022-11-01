/**
 * Author: Amin Zamani
 * File: id.go
 */

package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type ID struct {
	Identifier string `json:"identifier"`
	Expires    int64  `json:"expires"`
}

// New returns a new ID
func New() *ID {
	return &ID{}
}

// Compare this snippet from pkg/types/id.go:
func (i *ID) IsExpired() bool {
	return i.Expires < time.Now().Unix()
}

// Sets the expiration time:
func (i *ID) SetExpires(expiration time.Duration) {
	i.Expires = time.Now().Add(expiration).Unix()
}

// Sets the identifier
func (i *ID) SetIdentifier(identifier string) {
	i.Identifier = identifier
}

// Returns the identifier
func (i *ID) GetIdentifier() string {
	return i.Identifier
}

// Returns the expiration time
func (i *ID) GetExpires() int64 {
	return i.Expires
}

// Returns the expiration time as a time.Duration
func (i *ID) GetExpiration() time.Duration {
	return time.Until(time.Unix(i.Expires, 0))
}

// Returns the expiration time as a string
func (i *ID) GetExpirationString() string {
	return i.GetExpiration().String()
}

// Returns the ID as a JSON string
func (i *ID) ToJSON() string {
	return fmt.Sprintf("{\"identifier\":\"%s\",\"expires\":%d}", i.Identifier, i.Expires)
}

// Returns the ID from a JSON string
func FromJSON(str string) (*ID, error) {
	var id ID
	err := json.Unmarshal([]byte(str), &id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
