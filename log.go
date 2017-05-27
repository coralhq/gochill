package gochill

import (
	"io"
	"log"
	"os"

	"github.com/Sirupsen/logrus"
)

//Logger used to create new log instance
var (
	Logger       *log.Logger
	CustomOutput io.Writer
)

//NilWriter used as custom output writer when CustomOutput nil
type NilWriter struct{}

func (nw NilWriter) Write(p []byte) (int, error) { return 0, nil }

func init() {
	Logger = log.New(os.Stdout, "", 0)
	if CustomOutput == nil {
		CustomOutput = NilWriter{}
	}
}

//NewCustomOutput used to change how Logger send log message
func NewCustomOutput(output io.Writer) {
	Logger.SetOutput(output)
}

//Alert used to create log message with alert level
func Alert(msg MessageLog, options ...Option) {
	message := buildMessage(LevelAlert, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stderr, CustomOutput))

	if getOutput(os.Getenv(EnvOutputJSONFormat)) {
		Logger.Println(string(message))
	} else {
		logrus.Warning(string(message))
	}
}

//Info used to create log message with info level
func Info(msg MessageLog, options ...Option) {
	message := buildMessage(LevelInfo, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))

	if getOutput(os.Getenv(EnvOutputJSONFormat)) {
		Logger.Println(string(message))
	} else {
		logrus.Info(string(message))
	}
}

//Critical used to create log message with critical level
func Critical(msg MessageLog, options ...Option) {
	message := buildMessage(LevelCritical, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stderr, CustomOutput))

	if getOutput(os.Getenv(EnvOutputJSONFormat)) {
		Logger.Println(string(message))
	} else {
		logrus.Fatal(string(message))
	}
}

//Error used to create log message with error level
func Error(msg MessageLog, options ...Option) {
	message := buildMessage(LevelError, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stderr, CustomOutput))

	if getOutput(os.Getenv(EnvOutputJSONFormat)) {
		Logger.Println(string(message))
	} else {
		logrus.Error(string(message))
	}
}

//Warn used to create log message with warning level
func Warn(msg MessageLog, options ...Option) {
	message := buildMessage(LevelWarning, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))

	if getOutput(os.Getenv(EnvOutputJSONFormat)) {
		Logger.Println(string(message))
	} else {
		logrus.Warn(string(message))
	}
}

//Notice used to create log message with warning level
func Notice(msg MessageLog, options ...Option) {
	message := buildMessage(LevelNotice, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))

	if getOutput(os.Getenv(EnvOutputJSONFormat)) {
		Logger.Println(string(message))
	} else {
		logrus.Warn(string(message))
	}
}

//Debug used to create log message with warning level
func Debug(msg MessageLog, options ...Option) {
	message := buildMessage(LevelDebug, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))

	if getOutput(os.Getenv(EnvOutputJSONFormat)) {
		Logger.Println(string(message))
	} else {
		logrus.Debug(string(message))
	}
}

//buildMessage is a private function used to build message structure
func buildMessage(level int, msg MessageLog, options []Option) []byte {
	m := NewMessage(level)
	m.ShortMessage = msg.Short
	m.FullMessage = msg.Full

	mergedOptions := MergeOptionsToMap(options)

	if !m.OutputJSON {
		texts := messageText(m, mergedOptions)
		return texts
	}

	maps := CombineMaps(m.ToMap(), mergedOptions)

	return MapToJSON(maps)
}

func messageText(m *Message, options map[string]interface{}) []byte {
	maps := CombineMaps(m.TextMap(), options)

	return MapToJSON(maps)
}
