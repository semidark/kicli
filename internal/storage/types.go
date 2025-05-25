package storage

import "time"

// ShellCommandEntry represents a stored shell command
type ShellCommandEntry struct {
	Command   string
	Timestamp time.Time
}
