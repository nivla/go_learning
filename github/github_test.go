package github

import (
	"bytes"
	"testing"
)

func TestGetUserByName(t *testing.T) {
	client := New()
	user, err := client.GetUserByName("nivla")

	if err != nil {
		t.Errorf("Error found")
	}

	expected := "nivla"
	if !bytes.Equal([]byte(user.Login), []byte(expected)) {
		t.Errorf("Expected value `%s` but instad go `%s`", expected, user.Login)
	}

}
