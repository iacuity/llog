package llog_test

import (
	"testing"

	"github.com/llog"
)

func TestInit(t *testing.T) {

	fileName := "testlog.log"

	err := llog.Init(fileName)

	if nil != err {

		t.Fatal("Failed to Initialize Logger")

	}

	llog.SetLogLevel(llog.INFO)
	llog.Debug("Debug")
	llog.Info("Info")
	llog.Warn("Warn")
	llog.Error("Error")
	llog.Fatal("Fatal")
}
