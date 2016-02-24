package main

import (
	"errors"
	"fmt"
	"testing"
)

func UserTest(username string) error {
	// create user
	access_key, err := insertUser(username)
	if err != nil {
		return err
	}

	// get user
	user, err := getUser(username)
	if err != nil {
		return err
	}
	if user.Name != username {
		return errors.New(fmt.Sprintf("name is not equal: %s != %s", user.Name, username))
	}
	if user.AccessKey != access_key {
		return errors.New(fmt.Sprintf("access key is not equal: %s != %s", user.AccessKey, access_key))
	}

	// delete user
	err = deleteUser(username)
	if err != nil {
		return err
	}

	// confirm that user deleted
	user, err = getUser(username)
	if err == nil {
		return errors.New(fmt.Sprintf("User %s shouldn't be there", username))
	}

	return nil
}

func TestUserTest(t *testing.T) {
	err := UserTest("test")
	if err != nil {
		t.Error(err)
	}
}

func TestUserSQLinject(t *testing.T) {
	err := UserTest("'asdfasdf\"")
	if err != nil {
		t.Error(err)
	}
}
