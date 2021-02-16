package server

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	code := m.Run()
	os.Exit(code)
}
