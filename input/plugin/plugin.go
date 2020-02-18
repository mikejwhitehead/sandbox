package plugin

import (
	"sync"
)

// Plugin struct
type Plugin struct {
	Name string
	Interface
}

// Interface for Plugins
type Interface interface {
	Input(msg []byte) (byte, error) 
}

// Plugins global variable
var (
	Plugins []*Plugin
	mutex = &sync.Mutex{}
)

// Add function
func Add(name string, iface Interface){
	p := &Plugin{
		Name: name,
		Interface: iface,
	}

	mutex.Lock()
	Plugins = append(Plugins, p)
	mutex.Unlock()
}