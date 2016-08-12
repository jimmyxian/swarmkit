package orchestrator

import (
	"errors"
)

var (
	errStarted    = errors.New("already started")
	errNotStarted = errors.New("not started")
)
