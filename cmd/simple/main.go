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

	attention, err := rui.Attention()
	if err != nil {
		slog.Error("Failed to get attention", "error", err)
		os.Exit(1)
	}
	slog.Info("Attention", "attention", attention)

	hwModel, err := rui.GetHardwareModel()
	if err != nil {
		slog.Error("Failed to get hardware model", "error", err)
		os.Exit(1)
	}
	slog.Info("Hardware model", "model", hwModel)

	sn, err := rui.GetSerialNumber()
	if err != nil {
		slog.Error("Failed to get serial number", "error", err)
		os.Exit(1)
	}
	slog.Info("Serial number", "serial", sn)

	ver, err := rui.GetFirmwareVersion()
	if err != nil {
		slog.Error("Failed to get firmware version", "error", err)
		os.Exit(1)
	}
	slog.Info("Firmware version", "version", ver)

	apiVer, err := rui.GetAPIVersion()
	if err != nil {
		slog.Error("Failed to get API version", "error", err)
		os.Exit(1)
	}
	slog.Info("API version", "version", apiVer)

	devEUI, err := rui.GetDevEUI()
	if err != nil {
		slog.Error("Failed to get DevEUI", "error", err)
		os.Exit(1)
	}
	slog.Info("DevEUI", "devEUI", devEUI)

	appEUI, err := rui.GetAppEUI()
	if err != nil {
		slog.Error("Failed to get AppEUI", "error", err)
		os.Exit(1)
	}
	slog.Info("AppEUI", "appEUI", appEUI)

	appKey, err := rui.GetAppKey()
	if err != nil {
		slog.Error("Failed to get AppKey", "error", err)
		os.Exit(1)
	}
	slog.Info("AppKey", "appKey", appKey)
}
