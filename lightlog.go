// Package lightlog предоставляет лёгкую обёртку над стандартным log.Logger
// с фильтрацией сообщений по уровню логирования (Error, Warning, Info).
package lightlog

import (
	"log"
	"os"
)

// LogLevel задаёт уровень логирования. Чем выше значение, тем больше
// сообщений выводится: LogLevelError < LogLevelWarning < LogLevelInfo.
type LogLevel int

// LogLevelPrefix — текстовый префикс, который печатается перед сообщением
// соответствующего уровня.
type LogLevelPrefix string

// Logger оборачивает log.Logger и выводит сообщение только если его уровень
// не выше текущего уровня логирования, заданного через SetLogLevel.
type Logger struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
}

// Уровни логирования, упорядоченные от наименее к наиболее подробному.
const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

// Префиксы, печатаемые перед сообщениями соответствующих уровней.
const (
	LogLevelErrorPrefix   LogLevelPrefix = "ERR "
	LogLevelWarningPrefix LogLevelPrefix = "WARN"
	LogLevelInfoPrefix    LogLevelPrefix = "INFO"
)

// IsValid сообщает, является ли l одним из известных уровней логирования.
func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

// SetLogLevel задаёт текущий уровень логирования. Некорректное значение
// игнорируется, и предыдущий уровень сохраняется.
func (l *Logger) SetLogLevel(logLvl LogLevel) {
	if logLvl.IsValid() {
		l.logLevel = logLvl
	}
}

// Infoln печатает сообщение уровня Info (выводится только при LogLevelInfo).
func (l *Logger) Infoln(msg string) {
	l.println(LogLevelInfo, LogLevelInfoPrefix, msg)
}

// Warnln печатает сообщение уровня Warning.
func (l *Logger) Warnln(msg string) {
	l.println(LogLevelWarning, LogLevelWarningPrefix, msg)
}

// Errorln печатает сообщение уровня Error (выводится при любом валидном уровне).
func (l *Logger) Errorln(msg string) {
	l.println(LogLevelError, LogLevelErrorPrefix, msg)
}

// New создаёт Logger с выводом в os.Stderr, стандартными флагами времени
// (log.LstdFlags) и уровнем по умолчанию LogLevelError.
func New() *Logger {
	return &Logger{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}

// println выводит msg с префиксом, только если уровень сообщения srcLogLvl
// не выше текущего уровня логгера. Общий хелпер для Infoln/Warnln/Errorln.
func (l *Logger) println(srcLogLvl LogLevel, prefix LogLevelPrefix, msg string) {
	if l.logLevel >= srcLogLvl {
		l.Logger.Println(prefix, msg)
	}
}
