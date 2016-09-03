package gochill

/*
This library should be take responsible just to build log data structure.

Expected usages :

  log.Println(gochill.Info("short message"))

Each of method should return json string with valid structured data.
Each of method should initiate new instance of Message struct, we need to
do this to initiate host, service and timestamp value.
*/

import (
	"os"
	"time"
)

//Message struct used to build all required fields
type Message struct {
	Version      int    `json:"version"`
	Host         string `json:"host"`
	Service      string `json:"service"`
	ShortMessage string `json:"short_message"`
	Timestamp    int64  `json:"timestamp"`
	Level        int    `json:"level"`
	osEngine     OSHostEngine
}

//ToMap used to convert all exported fields into map string interface
func (m *Message) ToMap() map[string]interface{} {
	var maps map[string]interface{}
	maps = make(map[string]interface{})

	maps["version"] = m.Version
	maps["host"] = m.Host
	maps["service"] = m.Service
	maps["short_message"] = m.ShortMessage
	maps["timestamp"] = m.Timestamp
	maps["level"] = m.Level

	return maps
}

//ReplaceOSEngine used to change native golang os hostname implementation
func (m *Message) ReplaceOSEngine(oe OSHostEngine) {
	m.osEngine = oe
	m.Host = _parseHostName(m.osEngine)
}

//NewMessage used to create new instance of `message` struct
func NewMessage(level int) *Message {
	m := Message{}
	m.Version = Version
	m.Service = os.Getenv(EnvServiceKeyName)
	m.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	m.Level = level
	m.osEngine = os.Hostname
	m.Host = _parseHostName(m.osEngine)
	return &m
}

//parseHostName used to parse os hostname, but to simplify our library
//i just create this function, so we doesn't need to care about multiple
//return value for caller, is something goes wrong just return an empty string
func _parseHostName(os OSHostEngine) string {
	host, err := os()
	if err != nil {
		return ""
	}

	return host
}
