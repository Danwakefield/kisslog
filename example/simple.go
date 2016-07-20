package main

import "github.com/danwakefield/kisslog"

var log = kisslog.New("app")

func main() {
	log.Info("Requesting an image at %s", "foo/bar.jpg")
	log.Info("I have just completed a task", kisslog.Attrs{
		"foo": 1,
		"bar": "baz",
	})
	log.Debug("I bet you wont see this")
	log.Error("Failed to start, shutting down")
}
