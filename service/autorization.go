package service

import (
	"fmt"
	"net"
	"strings"
)

// Authorization executing command
func (serv Service) Authorization(userName string) (Response, error) {
	err := nameChecker(userName)
	if err != nil {
		return Response{}, err
	}
	if serv.IsUserAutorized(userName) {
		return Response{}, fmt.Errorf("user with name %s already  logged in", userName)
	}
	var resp = Response{
		ResponseTo: userName,
		Response:   fmt.Sprintf("Welcom %s\n Write /help for information\n", userName),
	}
	return resp, nil
}

// AddUser  is a method that saves the user in memory as a user online.
func (serv Service) AddUser(name string, conn net.Conn) bool {
	if strings.Trim(name, " ") == "" {
		return false
	}
	serv.UsersList[name] = Сache{
		NameUser: name,
		Conn:     conn,
	}
	return true
}

// IsUserAutorized is a method that checks if the user is authorized.
func (serv Service) IsUserAutorized(name string) bool {
	_, ok := serv.UsersList[name]
	return ok
}

// The nameChecker function checks if the name contains forbidden characters.
func nameChecker(name string) error {
	forbiddenChars := "'\" ! # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \\ ] ^ _` { | }"
	if strings.ContainsAny(name, forbiddenChars) {
		return fmt.Errorf("name shouldn't containt any of: %s", forbiddenChars)
	}
	if len(name) == 0 {
		return fmt.Errorf("please write your name")
	}
	if len(name) > 32 {
		return fmt.Errorf("name shouldn't containt more of 32 characters")
	}
	return nil
}
