package http

import (
	"fmt"
)

func (s *httpServer) Start() error {
	if s.cert != nil {
		return s.echo.StartTLS(s.Address(), s.cert, s.certKey)
	}

	return s.echo.Start(s.Address())
}

func (s *httpServer) Close() error {
	return s.echo.Close()
}

func (s *httpServer) Address() string {
	return fmt.Sprintf("[%s]:%s", s.host, s.port)
}
