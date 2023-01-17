package server

import (
	"fmt"
	"io"
	"net/http"
)

func (s *Server) relayRequests() []string {
	relayResponses := []string{}

	for _, relay := range s.relays {
		resp, err := http.Get(relay)
		if err != nil {
			relayResponses = append(relayResponses, fmt.Sprintf("request error: %v", err))
		} else {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				relayResponses = append(relayResponses, fmt.Sprintf("read error: %v", err))
			} else {
				relayResponses = append(relayResponses, string(body))
			}
		}
	}

	return relayResponses
}
