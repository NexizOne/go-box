package internal

import "runtime/debug"

// cmd argumets names
const (
	// info
	Name    = "go-box"
	Version = "0.3.0"
)

// vcs info
var Revision = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return ""
}()
