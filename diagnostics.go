package logd

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Priority of a journal message
type Priority int

const (
	PriEmerg Priority = iota
	PriAlert
	PriCrit
	PriErr
	PriWarning
	PriNotice
	PriInfo
	PriDebug
)

var (
	MachineIdentifier string
	ModuleName        string
	LogDirectory      string
)

func LogFatal(vars map[string]string, err error) {
	log.Println(err)
	Log(PriErr, vars, err)
	panic(err)
}

// Configure sets up the standard logging infrastructure to write to stdout and a local log file
// The final path of the log file is prefixed with a director path obtained from configuration
func Configure(modulename string, logDirectory string) {
	LogDirectory = logDirectory
	ModuleName = modulename
	machineIdentifier, err := getMachineIdentifier()
	if err != nil {
		log.Fatal(err)
	}

	MachineIdentifier = machineIdentifier

	configureLog()
}

func getMachineIdentifier() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln(err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return os.Hostname()
}

func setupTraditionalLog() {
	if _, err := os.Stat(LogDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(LogDirectory, 0777)
		if err != nil {
			log.Fatalln(err)
		}
	}

	filename := fmt.Sprintf("%s/%s.log", LogDirectory, ModuleName)
	log.Printf("Logging to logpath - %s", filename)

	var w *os.File
	var err error
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		w, err = os.Create(filename)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
	} else {
		w, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
	}

	log.SetOutput(io.MultiWriter(os.Stdout, w))
	log.SetPrefix(MachineIdentifier + " " + ModuleName + " ")

	// TODO get flags from configuration
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
