package service

import (
	"fmt"
	"log"
	"strings"
)

// signedUsersRequests is a method that processes requests from signed in users.
func (serv Service) signedUsersRequests(username, request string) (Response, error) {
	var err error
	if strings.Trim(request, " ") == "" {
		return Response{}, fmt.Errorf("empty request, write /help if you need more information")
	}
	log.Print("signed ", username, "  ", request)
	reqCom := strings.Fields(request)
	var resp = Response{
		ResponseTo: username,
	}
	if strings.Contains(reqCom[0], "@") {
		resp.Response, err = serv.sendMessage(username, request[1:len(reqCom[0])], strings.Join(reqCom[1:], " "))
		if err != nil {
			resp.Response = err.Error()
			return resp, err
		}
		return resp, nil
	}
	if strings.Contains(reqCom[0], "/") {
		if reqCom[0] == "/users_online" {
			res := serv.getUsersOnline()
			resp.Response = res
			return resp, nil
		}

		if reqCom[0] == "/help" {
			res := serv.help()
			resp.Response = res
			return resp, nil
		}
	}
	resp.Response = fmt.Sprintf("unknown command %s, write /help if you need more information \n", request)
	return resp, nil
}

// SendMessage is a method that sends message to online user
// or returns error if this user not online.
func (serv Service) sendMessage(fromName, toName, message string) (string, error) {
	m, ok := serv.UsersList[toName]
	log.Print(ok, m.NameUser)
	if !ok {
		return "", fmt.Errorf("there is no %s user online ", toName)
	}
	if strings.Trim(message, " ") == "" {
		return "", fmt.Errorf("empty message cannot be sent")
	}
	_, err := m.Conn.Write([]byte(fmt.Sprintf("\"%s\": %s\n", fromName, message)))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("your message has been successfully delivered  to %s \n", toName), nil
}

// getUsersOnline is a method that returns lists of online users from memory.
func (serv Service) getUsersOnline() string {
	var usersList string
	var userNum int
	for m := range serv.UsersList {
		userNum++
		if userNum == 1 {
			usersList = m
		} else {
			str := fmt.Sprintf("%s\n%s", usersList, m)
			usersList = str
		}
	}
	response := fmt.Sprintf("Users online: %d\n%s\n", userNum, usersList)
	return response
}
func (serv Service) help() string {
	return "use the @[name] to write to the user who is online, or /users_online to see the list of online users\n"
}
