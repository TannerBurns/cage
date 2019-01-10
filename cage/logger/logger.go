package logger

import (
	"fmt"
	"io"
	"log"
)

type Logger struct {
	handleName string
	Debug      *log.Logger
	Info       *log.Logger
	Warning    *log.Logger
	Error      *log.Logger
	Fatal      *log.Logger
}

func (logger *Logger) InitLogging(
	handleName string,
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	fatalHandle io.Writer,
) {
	logger.handleName = handleName
	logger.Debug = log.New(debugHandle, fmt.Sprintf("%s - DEBUG: ", logger.handleName),
		log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info = log.New(infoHandle, fmt.Sprintf("%s - INFO: ", logger.handleName),
		log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warning = log.New(warningHandle, fmt.Sprintf("%s - WARNING: ", logger.handleName),
		log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = log.New(errorHandle, fmt.Sprintf("%s - ERROR: ", logger.handleName),
		log.Ldate|log.Ltime|log.Lshortfile)
	logger.Fatal = log.New(fatalHandle, fmt.Sprintf("%s - FATAL: ", logger.handleName),
		log.Ldate|log.Ltime|log.Lshortfile)
}
