package main
import "log"

type LogLevel int
type LogExtended struct{
	logger
}

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

func () SetLogLevel(logLvl LogLevel) {

}
func () Infoln(msg string) {

}
func () Warnln(msg string)
{}
func () Errorln(msg string)
{}
func () NewLogExtended() LogExtended{

}