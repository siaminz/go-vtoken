package types

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	id := New()
	if id == nil {
		t.Error("ID is nil")
	}
}

func TestIsExpired(t *testing.T) {
	id := New()
	id.SetExpires(time.Hour)
	if id.IsExpired() {
		t.Error("ID is expired")
	}
}

func TestSetExpires(t *testing.T) {
	id := New()
	id.SetExpires(time.Hour)
	if id.IsExpired() {
		t.Error("ID is expired")
	}
}

func TestSetIdentifier(t *testing.T) {
	id := New()
	id.SetIdentifier("test")
	if id.GetIdentifier() != "test" {
		t.Error("Identifier is not equal")
	}
}

func TestGetIdentifier(t *testing.T) {
	id := New()
	id.SetIdentifier("test")
	if id.GetIdentifier() != "test" {
		t.Error("Identifier is not equal")
	}
}

func TestGetExpires(t *testing.T) {
	id := New()
	id.SetExpires(time.Hour)
	if id.GetExpires() < time.Now().Unix() {
		t.Error("ID is expired")
	}
}

func TestGetExpiration(t *testing.T) {
	id := New()
	id.SetExpires(time.Hour)
	if id.GetExpiration() >= time.Hour {
		t.Error("Expiration is not equal")
	}
}

func TestFromJSON(t *testing.T) {
	id := New()
	id.SetIdentifier("test")
	id.SetExpires(time.Hour)
	id2, err := FromJSON(id.ToJSON())
	if err != nil {
		t.Error("Error is not nil")
	}
	if id2.GetIdentifier() != "test" {
		t.Error("Identifier is not equal")
	}
	if id2.GetExpires() < time.Now().Unix() {
		t.Error("ID is expired")
	}
}

func TestFromJSONError(t *testing.T) {
	_, err := FromJSON("test")
	if err == nil {
		t.Error("Error is nil")
	}
}
