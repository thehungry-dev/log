// Package log provides structured and unstructured logging functions
package log

import (
	"github.com/thehungry-dev/cog"
	"github.com/thehungry-dev/cog/handler"
	"github.com/thehungry-dev/cog/message/field"
	"github.com/thehungry-dev/cog/settings"
)

var Private field.Builder
var fields field.Builder

var Default cog.Writer
var JSON cog.Writer

func String(name string, value string) field.Field         { return fields.String(name, value) }
func Int64(name string, value int64) field.Field           { return fields.Int64(name, value) }
func Int(name string, value int) field.Field               { return fields.Int(name, value) }
func Int32(name string, value int32) field.Field           { return fields.Int32(name, value) }
func Int16(name string, value int16) field.Field           { return fields.Int16(name, value) }
func Int8(name string, value int8) field.Field             { return fields.Int8(name, value) }
func Bool(name string, value bool) field.Field             { return fields.Bool(name, value) }
func Complex128(name string, value complex128) field.Field { return fields.Complex128(name, value) }
func Complex64(name string, value complex64) field.Field   { return fields.Complex64(name, value) }
func Float64(name string, value float64) field.Field       { return fields.Float64(name, value) }
func Float32(name string, value float32) field.Field       { return fields.Float32(name, value) }
func Nil(name string, value interface{}) field.Field       { return fields.Nil(name, value) }
func Object(name string, value interface{}) field.Field    { return fields.Object(name, value) }
func Array(name string, value interface{}) field.Field     { return fields.Array(name, value) }

func Tags(tags ...string) cog.Writer        { return Default.Tags(tags...) }
func Data(fields ...field.Field) cog.Writer { return Default.Data(fields...) }

func Trace(body string)                 { Default.Trace(body) }
func Tracef(f string, a ...interface{}) { Default.Tracef(f, a...) }
func Debug(body string)                 { Default.Debug(body) }
func Debugf(f string, a ...interface{}) { Default.Debugf(f, a...) }
func Info(body string)                  { Default.Info(body) }
func Infof(f string, a ...interface{})  { Default.Infof(f, a...) }
func Warn(body string)                  { Default.Warn(body) }
func Warnf(f string, a ...interface{})  { Default.Warnf(f, a...) }
func Error(body string)                 { Default.Error(body) }
func Errorf(f string, a ...interface{}) { Default.Errorf(f, a...) }
func Fatal(body string)                 { Default.Fatal(body) }
func Fatalf(f string, a ...interface{}) { Default.Fatalf(f, a...) }

func init() {
	Private = field.Builder{Private: true}

	set := settings.Getenv()
	device := set.Device()

	Default = cog.With(
		handler.BuildPipe(
			set.LevelFilter(),
			set.TagFilter(),
			set.OutputHandler(),
		),
	)

	JSON = cog.With(
		handler.BuildPipe(
			set.LevelFilter(),
			set.TagFilter(),
			handler.BuildOutput(device, handler.OutputJSON),
		),
	)
}
