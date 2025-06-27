package rui3

import (
	"fmt"
	"strings"
	"time"
)

func (r *RUI3) ResetMCU() error {
	err := r.SendRawCommand("ATZ")
	if err != nil {
		return fmt.Errorf("failed to send ATZ command: %w", err)
	}

	time.Sleep(5 * time.Second)

	return nil
}

func (r *RUI3) ResetFactoryDefaults() error {
	err := r.SendRawCommand("ATR")
	if err != nil {
		return fmt.Errorf("failed to send ATR command: %w", err)
	}

	time.Sleep(5 * time.Second)

	return nil
}

func (r *RUI3) Reset() error {
	err := r.ResetMCU()
	if err != nil {
		return fmt.Errorf("failed to reset MCU: %w", err)
	}

	err = r.ResetFactoryDefaults()
	if err != nil {
		return fmt.Errorf("failed to reset factory defaults: %w", err)
	}

	return nil
}

func (r *RUI3) Attention() (bool, error) {
	err := r.SendRawCommand("AT")
	if err != nil {
		return false, fmt.Errorf("failed to send AT command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return false, fmt.Errorf("failed to receive AT response: %w", err)
	}

	if strings.Contains(response, "OK") {
		return true, nil
	}

	return false, nil
}

func (r *RUI3) GetSerialNumber() (string, error) {
	err := r.SendRawCommand("AT+SN=?")
	if err != nil {
		return "", fmt.Errorf("failed to send AT+SN? command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to receive AT+SN response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "AT+SN=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("failed to receive AT+SN response")
}

func (r *RUI3) GetFirmwareVersion() (string, error) {
	err := r.SendRawCommand("AT+VER=?")
	if err != nil {
		return "", fmt.Errorf("failed to send AT+VER? command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to receive AT+VER response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "AT+VER=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("failed to receive AT+VER response")
}

func (r *RUI3) GetAPIVersion() (string, error) {
	err := r.SendRawCommand("AT+APIVER=?")
	if err != nil {
		return "", fmt.Errorf("failed to send AT+APIVER? command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to receive AT+APIVER response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "AT+APIVER=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("failed to receive AT+APIVER response")
}

func (r *RUI3) GetHardwareModel() (string, error) {
	err := r.SendRawCommand("AT+HWMODEL=?")
	if err != nil {
		return "", fmt.Errorf("failed to send AT+HWMODEL? command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to receive AT+HWMODEL response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "AT+HWMODEL=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("failed to receive AT+HWMODEL response")
}

func (r *RUI3) GetBootloaderVersion() (string, error) {
	err := r.SendRawCommand("AT+BOOTVER=?")
	if err != nil {
		return "", fmt.Errorf("failed to send AT+BLVER? command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to receive AT+BLVER response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "AT+BOOTVER=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("failed to receive AT+BLVER response")
}
