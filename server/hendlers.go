package server

import (
	"basic_chat_server/service"
	"bufio"
	"fmt"
	"log"
	"time"
)

// The serve method accepts connection, accepts requests, sends responses.
func (c Connection) serve(service service.Service) {
	defer func() {
		service.DeleteUser(c.UserName)
		log.Printf("closing connection from %v", c.conn.RemoteAddr())
		c.conn.Close()
	}()
	log.Printf("accepted connection from %v", c.conn.RemoteAddr())
	_, err := c.conn.Write([]byte("just write your name to continue\n"))
	if err != nil {
		log.Fatal(err)
	}
	err = c.conn.SetDeadline(time.Now().Add(c.IdleTimeout))
	if err != nil {
		log.Fatal(err)
	}
	scanr := bufio.NewScanner(bufio.NewReader(c.conn))
	for {
		err := c.updateDeadline()
		if err != nil {
			fmt.Print(err)
		}
		scanned := scanr.Scan()
		if !scanned {
			if err := scanr.Err(); err != nil {
				log.Printf("%v(%v)", err, c.conn.RemoteAddr())
				return
			}
			break
		}
		response, err := service.Execute(c.UserName, (scanr.Text()))
		if err != nil {
			log.Printf("%v(%v)", err, c.conn.RemoteAddr())
			c.conn.Write([]byte(err.Error() + "\n"))
			continue
		}
		if !c.IsLogged {
			c.UserName = response.ResponseTo
			c.IsLogged = service.AddUser(c.UserName, c.conn)
		}

		_, err = c.conn.Write([]byte(response.Response))
		if err != nil {
			log.Print(err)
		}
	}

}

func (c Connection) updateDeadline() error {
	idleDeadline := time.Now().Add(c.IdleTimeout)
	err := c.conn.SetDeadline(idleDeadline)
	if err != nil {
		return err
	}
	return nil
}
