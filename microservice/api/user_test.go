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

func TestingDoesUserExistShouldReturnTrue(t *testing.T) {
	expected := true
	actual := doesUserExist(1)
	if expected != actual {
		t.Errorf("User does not exists")
	}
}

func TestingDoesUserExist_ShouldReturnFalse(t *testing.T) {
	expected := false
	actual := doesUserExist(3)
	if expected != actual {
		t.Errorf("User does exist")
	}
}

func TestingValidateUser_ShouldReturnTrue(t *testing.T) {
	expected := true
	user := User{FirstName: "Joakim", LastName: "Bajoul", NickName: "Todai", ID: 1, Country: "United Kingdom", Email: "Joabaj88@gmail.com", Password: "qwerty123"}

	actual := user.validate()
	if expected != actual {
		t.Errorf("Validation failed")
	}
}

func TestingValidateUser_ShouldReturnFalse(t *testing.T) {
	expected := false
	user := User{FirstName: "", LastName: "Bajoul", NickName: "Todai", ID: 1, Country: "United Kingdom", Email: "Joabaj88@gmail.com", Password: "qwerty123"}

	actual := user.validate()
	if expected != actual {
		t.Errorf("Validation succeeded")
	}
}

func TestingCompareEqual_ShouldReturnTrue(t *testing.T) {
	expected := true
	actual := compareEquals("HELLO", "hello")

	if expected != actual {
		t.Errorf("Compare failed")
	}
}

func TestingFilterUsers_ShouldReturnOneUser(t *testing.T) {
	expected := 1
	actual := filterUsers(map[string][]string{"Country": []string{"USA"}})

	if expected != len(actual) {
		t.Errorf("Incorrect length")
	}
}
