package log

import "github.com/sirupsen/logrus"

const (
	encodingTypePlain string = "PLAIN"
	encodingTypeJSON  string = "JSON"
)

const (
	LevelDebugString = "DEBUG"
	LevelInfoString  = "INFO"
	LevelWarnString  = "WARN"
	LevelErrorString = "ERROR"
)

// DEBUG INFO WARN ERROR
var envLevelsMapping = map[string]logrus.Level{
	LevelDebugString: logrus.DebugLevel,
	LevelInfoString:  logrus.InfoLevel,
	LevelWarnString:  logrus.WarnLevel,
	LevelErrorString: logrus.ErrorLevel,
}
