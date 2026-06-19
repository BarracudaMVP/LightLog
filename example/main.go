package main

import "github.com/barracudamvp/lightlog"

func main() {
	logger := lightlog.New()
	logger.SetLogLevel(lightlog.LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
