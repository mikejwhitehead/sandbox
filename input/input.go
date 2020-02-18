package input

import (
	"github.com/mikejwhitehead/sandbox/config"
	"github.com/mikejwhitehead/sandbox/input/plugin"
)

// Input struct
type Input struct {
	Plugins []*plugin.Plugin
}

// Session struct
type Session struct {
	Config *config.Config
	Input Input
	Errors chan error
}

// NewSession function creates the ingest session
func NewSession(cfg *config.Config, errors chan error) *Session {
	s := &Session{
		Config: cfg,
		Input: &Input{
			Plugins: plugin.Plugins,
		},
		Errors: errors,
	}

	return s
}