package main

import (
	"log/slog"
	"os"
	"tencorvids/rui3-go"
)

func main() {
	portName := "/dev/ttyS0"
	slog.Info("Using port", "port", portName)

	rui, err := rui3.New(portName)
	if err != nil {
		slog.Error("Failed to create RUI3 instance", "error", err)
		os.Exit(1)
	}
	defer rui.Close()

	resp, err := rui.AT()
	if err != nil {
		slog.Error("Failed to send AT command", "error", err)
		os.Exit(1)
	}
	slog.Info("AT response", "response", resp)
}
