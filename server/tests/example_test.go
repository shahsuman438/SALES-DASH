package tests

import (
	"fmt"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	usr, err := getUser("1")
	if err != nil {
		testFailed(err)
	}
	fmt.Println(usr.Name)
	fmt.Println(usr.Email)
	logTestInformation("test info")
	testPassed("Test passed")
}

// Todo: remove this after some function implementation in core
type user struct {
	Id    string `json:id`
	Name  string `json:"name"`
	Email string `json:email`
}

func getUser(Id string) (*user, error) {
	user := []user{
		{
			Id:    "0",
			Name:  "yxz",
			Email: "yxz@lftechnology.com",
		},
		{
			Id:    "1",
			Name:  "abc",
			Email: "abc@lftechnology.com",
		},
	}
	for _, usr := range user {
		if strings.EqualFold(usr.Id, Id) {
			return &usr, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
