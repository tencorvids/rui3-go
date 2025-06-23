package rui3

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"go.bug.st/serial"
)

type RUI3 struct {
	port   serial.Port
	reader *bufio.Reader
	writer *bufio.Writer

	lastResponse string
}

func New(portName string) (*RUI3, error) {
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		return nil, fmt.Errorf("failed to open serial port %s: %w", portName, err)
	}

	return &RUI3{
		port:   port,
		reader: bufio.NewReader(port),
		writer: bufio.NewWriter(port),
	}, nil
}

func (r *RUI3) Close() error {
	return r.port.Close()
}

func (r *RUI3) SetReadTimeout(timeout time.Duration) error {
	return r.port.SetReadTimeout(timeout)
}

func (r *RUI3) ResetInputBuffer() error {
	return r.port.ResetInputBuffer()
}

func (r *RUI3) ResetOutputBuffer() error {
	return r.port.ResetOutputBuffer()
}

func (r *RUI3) Drain() error {
	return r.port.Drain()
}

func (r *RUI3) SendRawCommand(cmd string) error {
	r.ResetInputBuffer()

	_, err := r.writer.WriteString(cmd + "\r\n")
	if err != nil {
		return fmt.Errorf("failed to send command: %w", err)
	}

	return r.writer.Flush()
}

func (r *RUI3) RecvResponse(timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var response strings.Builder
	lines := make([]string, 0)

	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("timeout waiting for response after %v", timeout)
		default:
			lineChan := make(chan string, 1)
			errChan := make(chan error, 1)

			go func() {
				line, err := r.reader.ReadString('\n')
				if err != nil {
					errChan <- err
					return
				}
				lineChan <- line
			}()

			select {
			case line := <-lineChan:
				line = strings.TrimSpace(line)
				if line != "" {
					lines = append(lines, line)
					response.WriteString(line)
					response.WriteString("\n")
				}

				if strings.Contains(line, "OK") {
					result := response.String()
					r.lastResponse = result
					return result, nil
				}
				if strings.Contains(line, "+EVT:TX_DONE") || strings.Contains(line, "+EVT:SEND_CONFIRMED_OK") {
					result := response.String()
					r.lastResponse = result
					return result, nil
				}
				if strings.Contains(line, "+EVT:TXP2P DONE") {
					result := response.String()
					r.lastResponse = result
					return result, nil
				}
				if strings.Contains(line, "AT_COMMAND_NOT_FOUND") ||
					strings.Contains(line, "AT_PARAM_ERROR") ||
					strings.Contains(line, "SEND_CONFIRMED_FAILED") ||
					strings.Contains(line, "AT_NO_NETWORK_JOINED") {
					result := response.String()
					r.lastResponse = result
					return result, fmt.Errorf("command error: %s", line)
				}
			case err := <-errChan:
				if err == io.EOF {
					if len(lines) > 0 {
						result := response.String()
						r.lastResponse = result
						return result, nil
					}
					continue
				}
				continue
			case <-time.After(100 * time.Millisecond):
				if len(lines) > 0 {
					result := response.String()
					r.lastResponse = result
					return result, nil
				}
				continue
			}
		}
	}
}

func (r *RUI3) GetLastResponse() string {
	return r.lastResponse
}
