package main

import (
	"flag"
	"github.com/hsmade/gogallery/pkg/server"
	"github.com/sirupsen/logrus"
)

var (
	listenPort int
	rootPath   string
	verbose    bool
)

func main() {
	flag.IntVar(&listenPort, "listen-port", 80, "the port to listen on")
	flag.StringVar(&rootPath, "root-path", ".", "the root path to serve pictures from")
	flag.BoolVar(&verbose, "verbose", false, "enable debug logging")
	flag.Parse()

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("enabled debug level")
	}

	s := server.New(listenPort, rootPath)
	logrus.Fatal(s.Run())
}
