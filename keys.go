package rui3

import (
	"fmt"
	"strings"
	"time"
)

func (r *RUI3) GetDevEUI() (string, error) {
	err := r.SendRawCommand("AT+DEVEUI=?")
	if err != nil {
		return "", fmt.Errorf("failed to send deveui command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to receive deveui response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "AT+DEVEUI=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("DEVEUI not found in response: %s", response)
}

func (r *RUI3) GetAppKey() (string, error) {
	err := r.SendRawCommand("AT+APPKEY=?")
	if err != nil {
		return "", fmt.Errorf("failed to send appkey command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to receive appkey response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "AT+APPKEY=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("APPKEY not found in response: %s", response)
}

func (r *RUI3) GetAppEUI() (string, error) {
	err := r.SendRawCommand("AT+APPEUI=?")
	if err != nil {
		return "", fmt.Errorf("failed to send appeui command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to receive appeui response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "AT+APPEUI=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("APPEUI not found in response: %s", response)
}
