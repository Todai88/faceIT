package api

import (
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	expectedLength := 2
	actualLength := len(getSlicedUsers())
	if expectedLength != actualLength {
		t.Errorf("Incorrect length")
	}
}
