package server

import (
	"basic_chat_server/service"
	"log"
	"net"
	"time"
)

// Server is struct that's working with TCP methods
// stores the TCP connection  address and users connection timeout.
type Server struct {
	Addr        string
	IdleTimeout time.Duration
}

// Connection is  structure that's stores the client's connection,
// his name, information about his authorization and connection timeout.
type Connection struct {
	conn        net.Conn
	UserName    string
	IsLogged    bool
	IdleTimeout time.Duration
}

// ListenAndServe listens on the TCP network address addr and then calls
func (srv Server) ListenAndServe() error {
	if srv.Addr == "" {
		srv.Addr = ":8080"
	}
	if srv.IdleTimeout.Seconds() == 0.0 {
		srv.IdleTimeout = (time.Minute * 2)
	}
	log.Printf("starting server on %v\n", srv.Addr)
	listener, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}
	serv := service.InitService()
	defer listener.Close()
	for {
		var c Connection
		c.IdleTimeout = srv.IdleTimeout
		c.conn, err = listener.Accept()
		if err != nil {
			log.Printf("error accepting connection %v", err)
			continue
		}
		go c.serve(serv)
	}
}
