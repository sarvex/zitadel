package execution

import (
	"time"

	"github.com/zitadel/zitadel/internal/logstore"
)

var _ logstore.LogRecord = (*Record)(nil)

type Record struct {
	Timestamp      time.Time
	InstanceID     string
	OrganizationID string
	ActionID       string
	RunID          string
	Message        string
	Level          string
	FileDescriptor FileDescriptor
}

type FileDescriptor string

const (
	StdOut FileDescriptor = "stdout"
	StdErr FileDescriptor = "stderr"
)

func (e *Record) RedactSecrets() logstore.LogRecord {
	// TODO implement
	return e
}