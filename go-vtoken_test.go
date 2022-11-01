/**
 * Author: Amin Zamani
 * File: go-vtoken_test.go
 */
package vtoken

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	token, err := New("test", time.Hour)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("Token is empty")
	}
}

func TestVerify(t *testing.T) {
	token, err := New("test", time.Hour)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("Token is empty")
	}
	id, err := Verify(token)
	if err != nil {
		t.Error(err)
	}
	if id.GetIdentifier() != "test" {
		t.Error("Identifier is not equal")
	}
	if id.IsExpired() {
		t.Error("Token is expired")
	}
}
