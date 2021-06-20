package log_test

import (
	"testing"

	"github.com/thehungry-dev/log"
)

func TestLog(t *testing.T) {
	log.
		Tags("data", "trace").
		Data(
			log.String("SomeKey", "some value"),
			log.Int("SomeOtherKey", 123),
			log.Private.String("SomePrivateKey", "a value"),
		).
		Trace("trace msg")
	log.
		Tags("data", "debug").
		Data(
			log.String("SomeKey", "some value"),
			log.Int("SomeOtherKey", 123),
			log.Bool("SomeBool", true),
		).
		Debugd()
	log.Info("info message")
	log.
		JSON.
		Tags("data", "warn").
		Data(
			log.String("SomeKey", "some value"),
			log.Int("SomeOtherKey", 123),
			log.Bool("SomeBool", true),
		).
		Warn("warn message")
	log.Errorf("error %s", "message")
	log.
		Tags("fatal").
		Fatal("fatal msg")
}
