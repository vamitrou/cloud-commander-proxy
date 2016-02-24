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
func TestEnableDisable(t *testing.T) {
	_, err := insertUser("test")
	if err != nil {
		t.Error(err)
	}
	user, err := getUser("test")
	if err != nil {
		t.Error(err)
	}
	if user.Active != 1 {
		t.Error("User not active after creation")
	}
	err = setActive("test", 0)
	if err != nil {
		t.Error(err)
	}
	user, err = getUser("test")
	if user.Active != 0 {
		t.Error("User active after disabling")
	}
	err = setActive("test", 1)
	if err != nil {
		t.Error(err)
	}
	user, err = getUser("test")
	if user.Active != 1 {
		t.Error("User not active after enabling")
	}
	err = deleteUser("test")
	if err != nil {
		t.Error(err)
	}
}

func TestRenewKey(t *testing.T) {
	old_key, err := insertUser("test")
	if err != nil {
		t.Error(err)
	}
	new_key, err := renewUserKey("test")
	if err != nil {
		t.Error(err)
	}
	if old_key == new_key {
		t.Error("Old and new key are the same")
	}
	user, err := getUser("test")
	if err != nil {
		t.Error(err)
	}
	if user.AccessKey != new_key {
		t.Error("Key is not changed")
	}
}
