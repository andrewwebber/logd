// +build windows

package logd

import "log"

// Log provides the ability to log additional metadata or analytics with a message
func Log(priority Priority, vars map[string]string, v ...interface{}) error {
	log.Println(v)
	return nil
}

func configureLog() {
	setupTraditionalLog()
}
