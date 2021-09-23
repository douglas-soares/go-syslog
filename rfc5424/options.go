package rfc5424

import (
	syslog "github.com/douglas-soares/go-syslog/v3"
)

// WithBestEffort enables the best effort mode.
func WithBestEffort() syslog.MachineOption {
	return func(m syslog.Machine) syslog.Machine {
		m.WithBestEffort()
		return m
	}
}
