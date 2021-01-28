package service

import (
	"net"
)

// Response is structure accepts the user's public key that should receive the response
// and contains the response itself.
type Response struct {
	ResponseTo string
	Response   string
}

// Service is structure contains methods that process
// user authorization and requests and sends responses
type Service struct {
	UsersList map[string]Сache
}

// Сache  is structure that contains the public key of the user
// and his connection to the server.
type Сache struct {
	NameUser string
	Conn     net.Conn
}

// InitService is returns the service structure.
func InitService() Service {
	var serv Service
	serv.UsersList = make(map[string]Сache)
	return serv
}

// Execute is a method that initializes the authorization of authorized users
// or a method that listens for requests from authorized users.
func (serv Service) Execute(username, request string) (Response, error) {
	if !serv.IsUserAutorized(username) {
		return serv.Authorization(request)
	}
	return serv.signedUsersRequests(username, request)
}

// DeleteUser is a method that deletes the user from memory.
func (serv Service) DeleteUser(name string) {
	delete(serv.UsersList, name)
}
