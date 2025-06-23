package rui3

import (
	"fmt"
	"strings"
	"time"
)

func (r *RUI3) JoinNetwork() error {
	err := r.SendRawCommand("AT+JOIN=?")
	if err != nil {
		return fmt.Errorf("failed to send join command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return fmt.Errorf("failed to receive join response: %w", err)
	}

	if strings.Contains(response, "OK") {
		return nil
	}

	return fmt.Errorf("failed to join network: %s", response)
}

func (r *RUI3) JoinNetworkWithParams(join bool, autoJoin bool, retryInterval int, joinAttempts int) error {
	if retryInterval < 7 || retryInterval > 255 {
		return fmt.Errorf("invalid retry interval: %d", retryInterval)
	}

	if joinAttempts < 0 || joinAttempts > 255 {
		return fmt.Errorf("invalid join attempts: %d", joinAttempts)
	}

	joinCmd := "0"
	if join {
		joinCmd = "1"
	}

	autoJoinCmd := "0"
	if autoJoin {
		autoJoinCmd = "1"
	}

	cmd := fmt.Sprintf("AT+JOIN=%s:%s:%d:%d", joinCmd, autoJoinCmd, retryInterval, joinAttempts)

	err := r.SendRawCommand(cmd)
	if err != nil {
		return fmt.Errorf("failed to send join command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return fmt.Errorf("failed to receive join response: %w", err)
	}

	if strings.Contains(response, "OK") {
		return nil
	}

	return fmt.Errorf("failed to set join network settings: %s", response)
}

func (r *RUI3) JoinStatus() (bool, error) {
	err := r.SendRawCommand("AT+NJS=?")
	if err != nil {
		return false, fmt.Errorf("failed to send joinstatus command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return false, fmt.Errorf("failed to receive joinstatus response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "AT+NJS=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]) == "1", nil
			}
		}
	}

	return false, fmt.Errorf("NJS not found in response: %s", response)
}

func (r *RUI3) GetConfirmMode() (bool, error) {
	err := r.SendRawCommand("AT+CFM=?")
	if err != nil {
		return false, fmt.Errorf("failed to send cfm command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return false, fmt.Errorf("failed to receive cfm response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "AT+CFM=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]) == "1", nil
			}
		}
	}

	return false, fmt.Errorf("CFM not found in response: %s", response)
}

func (r *RUI3) SetConfirmMode(confirm bool) error {
	mode := "0"
	if confirm {
		mode = "1"
	}

	err := r.SendRawCommand(fmt.Sprintf("AT+CFM=%s", mode))
	if err != nil {
		return fmt.Errorf("failed to send cfm command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return fmt.Errorf("failed to receive cfm response: %w", err)
	}

	if strings.Contains(response, "OK") {
		return nil
	}

	return fmt.Errorf("failed to set confirm mode: %s", response)
}

type Class int

const (
	ClassA Class = iota
	ClassB
	ClassC
)

func (r *RUI3) SetClass(class Class) error {
	if class < ClassA || class > ClassC {
		return fmt.Errorf("invalid class: %d", class)
	}

	var classCmd string
	switch class {
	case ClassA:
		classCmd = "A"
	case ClassB:
		classCmd = "B"
	case ClassC:
		classCmd = "C"
	}

	err := r.SendRawCommand(fmt.Sprintf("AT+CLASS=%s", classCmd))
	if err != nil {
		return fmt.Errorf("failed to send class command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return fmt.Errorf("failed to receive class response: %w", err)
	}

	if strings.Contains(response, "OK") {
		return nil
	}

	return fmt.Errorf("failed to set class: %s", response)
}

func (r *RUI3) GetClass() (Class, error) {
	err := r.SendRawCommand("AT+CLASS=?")
	if err != nil {
		return ClassA, fmt.Errorf("failed to send class command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return ClassA, fmt.Errorf("failed to receive class response: %w", err)
	}

	lines := strings.SplitSeq(response, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "AT+CLASS=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				classValue := strings.TrimSpace(parts[1])
				if colonIndex := strings.Index(classValue, ":"); colonIndex != -1 {
					classValue = classValue[:colonIndex]
				}
				switch classValue {
				case "A":
					return ClassA, nil
				case "B":
					return ClassB, nil
				case "C":
					return ClassC, nil
				}
				return ClassA, fmt.Errorf("invalid class: %s", classValue)
			}
		}
	}

	return ClassA, fmt.Errorf("CLASS not found in response: %s", response)
}

func (r *RUI3) SetAdaptiveDataRate(enabled bool) error {
	enabledCmd := "0"
	if enabled {
		enabledCmd = "1"
	}

	err := r.SendRawCommand(fmt.Sprintf("AT+ADR=%s", enabledCmd))
	if err != nil {
		return fmt.Errorf("failed to send adr command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return fmt.Errorf("failed to receive adr response: %w", err)
	}

	if strings.Contains(response, "OK") {
		return nil
	}

	return fmt.Errorf("failed to set adaptive data rate: %s", response)
}

type ChannelMask int

const (
	SubBandAll ChannelMask = 0
	SubBand1   ChannelMask = 1
	SubBand2   ChannelMask = 2
	SubBand3   ChannelMask = 3
	SubBand4   ChannelMask = 4
	SubBand5   ChannelMask = 5
	SubBand6   ChannelMask = 6
	SubBand7   ChannelMask = 7
	SubBand8   ChannelMask = 8
	SubBand9   ChannelMask = 9
	SubBand10  ChannelMask = 10
	SubBand11  ChannelMask = 11
	SubBand12  ChannelMask = 12
)

func (r *RUI3) SetChannelMask(mask ChannelMask) error {
	// check if mask is valid
	if mask < SubBandAll || mask > SubBand12 {
		return fmt.Errorf("invalid channel mask: %d", mask)
	}

	maskCmd := "0000"
	switch mask {
	case SubBand1:
		maskCmd = "0001"
	case SubBand2:
		maskCmd = "0002"
	case SubBand3:
		maskCmd = "0004"
	case SubBand4:
		maskCmd = "0008"
	case SubBand5:
		maskCmd = "0010"
	case SubBand6:
		maskCmd = "0020"
	case SubBand7:
		maskCmd = "0040"
	case SubBand8:
		maskCmd = "0080"
	case SubBand9:
		maskCmd = "0100"
	case SubBand10:
		maskCmd = "0200"
	case SubBand11:
		maskCmd = "0400"
	case SubBand12:
		maskCmd = "0800"
	}

	err := r.SendRawCommand(fmt.Sprintf("AT+MASK=%s", maskCmd))
	if err != nil {
		return fmt.Errorf("failed to send cmask command: %w", err)
	}

	response, err := r.RecvResponse(5 * time.Second)
	if err != nil {
		return fmt.Errorf("failed to receive cmask response: %w", err)
	}

	if strings.Contains(response, "OK") {
		return nil
	}

	return fmt.Errorf("failed to set channel mask: %s", response)
}
