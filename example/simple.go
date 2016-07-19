package main

import "github.com/danwakefield/kisslog"

var log = logger.New("app")

func main() {
	log.Info("Starting at %d", 9088)

	log.Info("Requesting an image at foo/bar.jpg")
	log.Debug("I bet you wont see this")

	log.Error("Failed to start, shutting down...")
}
