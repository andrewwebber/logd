package logd

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	Configure("logd_test", "logd")
	log.Println("Simple Log Message")
	Log(PriWarning, map[string]string{"UNIT_TEST_NAME": "TestLog"}, "Simple Journal Message")
}
