package utils

import "testing"

func TestCreateUserStorageSpace(t *testing.T) {
	for _, s := range []string{"1", "2", "3", "4"}{
		CreateUserStorageSpace(s)
	}
}

func TestClearUserStorageSpace(t *testing.T) {
	for _, s := range []string{"1", "2", "3", "4"}{
		ClearUserStorageSpace(s)
	}
}
