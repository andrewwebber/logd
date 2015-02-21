// +build !windows

package logd

import (
	"fmt"
	"log"

	"github.com/coreos/go-systemd/journal"
)

// Log provides the ability to log additional metadata or analytics with a message
func Log(priority Priority, vars map[string]string, v ...interface{}) error {
	if !journal.Enabled() {
		log.Println(v)
	} else {
		if vars == nil {
			vars = map[string]string{}
		}

		vars["APPLICATION_IP"] = MachineIdentifier
		vars["APPLICATION_MODULE"] = "true"
		vars["APPLICATION_MODULENAME"] = ModuleName

		message := fmt.Sprintln(v...)
		return journal.Send(message, journal.Priority(priority), vars)
	}

	return nil
}

func configureLog() {
	if !journal.Enabled() {
		setupTraditionalLog()
	} else {
		log.SetFlags(0)
		journalWriter := JournalWriter{MachineIdentifier, ModuleName}
		log.SetOutput(journalWriter)
		log.Println("Logging to Systemd-journal")
	}
}
