package server

import (
	"math/rand"
	"net/http"
	"time"
)

func (s *Server) generateStatusCode() int {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100) + 1
	if randomNum <= s.statusCodeRate {
		return s.statusCode
	}
	return http.StatusOK
}

func (s *Server) generateDelay() time.Duration {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100) + 1
	if randomNum > s.delayRate {
		return 0
	}

	delayTime := rand.Intn(s.delayMax - s.delayMin + 1)
	delayTime += s.delayMin
	return time.Duration(delayTime) * time.Millisecond
}
