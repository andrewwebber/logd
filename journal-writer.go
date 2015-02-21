// +build !windows

package logd

import (
	"github.com/coreos/go-systemd/journal"
)

// JournalWriter writes log messages to systemd journal
type JournalWriter struct {
	MachineIdentifier string
	ModuleName        string
}

// Writes log message to systemd journal
func (w JournalWriter) Write(p []byte) (int, error) {
	return 0, journal.Send(string(p), journal.PriInfo, map[string]string{"APPLICATION_IP": w.MachineIdentifier, "APPLICATION_MODULE": "true", "APPLICATION_MODULENAME": w.ModuleName})
}
