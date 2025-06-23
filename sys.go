package rui3

import (
	"fmt"
	"strings"
	"time"
)

func (r *RUI3) AT() (string, error) {
	err := r.SendRawCommand("AT")
	if err != nil {
		return "", fmt.Errorf("failed to send AT command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to receive AT response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "AT") {
			return line, nil
		}
	}

	return response, nil
}
