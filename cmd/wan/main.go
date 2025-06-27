package main

import (
	"log/slog"
	"os"
	"tencorvids/rui3-go"
	"time"
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

	slog.Info("Resetting chip, this will take a up to 15 seconds...")
	err = rui.Reset()
	if err != nil {
		slog.Error("Failed to reset", "error", err)
		os.Exit(1)
	}
	slog.Info("Chip reset, resuming...")

	attention, err := rui.Attention()
	if err != nil {
		slog.Error("Failed to get attention", "error", err)
		os.Exit(1)
	}
	slog.Info("Attention", "attention", attention)

	regionBand, err := rui.GetRegionBand()
	if err != nil {
		slog.Error("Failed to get region band", "error", err)
		os.Exit(1)
	}
	slog.Info("Region band", "band", regionBand)

	channelMask, err := rui.GetChannelMask()
	if err != nil {
		slog.Error("Failed to get channel mask", "error", err)
		os.Exit(1)
	}
	slog.Info("Channel mask", "mask", channelMask)

	err = rui.JoinNetwork()
	if err != nil {
		slog.Error("Failed to join network", "error", err)
		os.Exit(1)
	}
	slog.Info("Joined network")

	maxAttempts := 10
	attempts := 0
	for {
		if attempts >= maxAttempts {
			slog.Error("Max attempts reached, exiting")
			os.Exit(1)
		}

		networkStatus, err := rui.JoinStatus()
		if err != nil {
			slog.Error("Failed to get join status", "error", err)
			os.Exit(1)
		}
		slog.Info("Join status", "status", networkStatus)

		if networkStatus {
			break
		}

		time.Sleep(4 * time.Second)
		attempts++
	}
	slog.Info("Connected to network")
}
