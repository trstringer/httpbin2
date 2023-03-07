package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Server is the web server.
type Server struct {
	port           int
	statusCode     int
	statusCodeRate int
	delayMin       int
	delayMax       int
	delayRate      int
	message        string
	relays         []string
}

// New creates a new web server.
func New(opts ...func(*Server)) *Server {
	svr := &Server{}
	for _, opt := range opts {
		opt(svr)
	}
	return svr
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithQOS(statusCode, statusCodeRate, delayMin, delayMax, delayRate int) func(*Server) {
	return func(s *Server) {
		s.statusCode = statusCode
		s.statusCodeRate = statusCodeRate
		s.delayMin = delayMin
		s.delayMax = delayMax
		s.delayRate = delayRate
	}
}

func WithMessage(message string) func(*Server) {
	return func(s *Server) {
		s.message = message
	}
}

func WithRelays(relays []string) func(*Server) {
	return func(s *Server) {
		s.relays = relays
	}
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	statusCode := s.generateStatusCode()
	delay := s.generateDelay()
	fmt.Printf("statusCode: %d, delay: %d ms\n", statusCode, delay.Milliseconds())

	relayOutput := s.relayRequests()

	time.Sleep(delay)

	w.WriteHeader(statusCode)
	w.Write([]byte(strings.Join(append([]string{s.message}, relayOutput...), ", ")))
}

// Run starts an instance of the web server.
func (s *Server) Run() {
	fmt.Println("Running server...")
	http.HandleFunc("/", s.handler)
	http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}
