package main

import (
	"errors"
	"net"
	"os"
	"os/exec"
	"time"
)

// Commander interface
type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfo, error)
}

// PingResult holds the result of a ping command
type PingResult struct {
	Successful bool
	Time       time.Duration
}

// SystemInfo holds system information
type SystemInfo struct {
	Hostname  string
	IPAddress string
}

// commander is the concrete implementation of Commander
type commander struct{}

func NewCommander() Commander {
	return &commander{}
}

func (c *commander) Ping(host string) (PingResult, error) {
	start := time.Now()
	cmd := exec.Command("ping", "-c", "1", host)
	if err := cmd.Run(); err != nil {
		return PingResult{Successful: false, Time: 0}, err
	}
	return PingResult{Successful: true, Time: time.Since(start)}, nil
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return SystemInfo{}, err
	}

	var ipAddress string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ipAddress = ipNet.IP.String()
			break
		}
	}

	if ipAddress == "" {
		return SystemInfo{}, errors.New("could not determine IP address")
	}

	return SystemInfo{Hostname: hostname, IPAddress: ipAddress}, nil
}
