package gochill

/*
For all Level* descriptions, please refer to
https://en.wikipedia.org/wiki/Syslog#Severity_level
*/

const (

	//EnvServiceKeyName used as key to get service name from os environment variable
	EnvServiceKeyName string = "SERVICE_NAME"

	//EnvOutputJSONFormat used as key to set output to json or text
	EnvOutputJSONFormat string = "JSON_FORMAT"

	//Version indicate version of the payload
	Version int = 1

	//LevelAlert indicate log level of "alert"
	LevelAlert int = 1

	//LevelCritical indicate log level of "critical"
	LevelCritical int = 2

	//LevelError indicate log level of "error"
	LevelError int = 3

	//LevelWarning indicate log level of "warning"
	LevelWarning int = 4

	//LevelNotice indicate log level of "notice"
	LevelNotice int = 5

	//LevelInfo indicate log level of "info"
	LevelInfo int = 6

	//LevelDebug indicate log level of "debug"
	LevelDebug int = 7
)
