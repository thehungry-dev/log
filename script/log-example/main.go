package main

import "github.com/thehungry-dev/log"

func main() {
	log.
		Tags("someOtherTag", "trace").
		Data(
			log.String("SomeKey", "some value"),
			log.Int("SomeOtherKey", 123),
			log.Private.String("SomePrivateKey", "a value"),
		).
		Trace("trace message")

	log.
		Tags("someTag", "debug").
		Data(
			log.String("SomeKey", "some value"),
			log.Int("SomeOtherKey", 123),
			log.Bool("SomeBool", true),
		).
		Debugd()

	log.Info("info message")

	log.
		JSON.
		Tags("someTag", "warn").
		Data(
			log.String("SomeKey", "some value"),
			log.Int("SomeOtherKey", 123),
			log.Bool("SomeBool", true),
		).
		Warn("warn message")

	log.Errorf("error %s", "message")

	log.
		Tags("fatal").
		Fatal("fatal message")
}
