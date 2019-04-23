// Package implement logger logic.
package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

var isLog bool = true

func SetIsLogOut(f bool) {
	isLog = f
}

// Create new logger.
func NewLogger(prefix string, successHandler map[int]string, errorHandler map[int]string) *Logger {
	log := logrus.New()
	form := prefixed.TextFormatter{
		ForceColors:      true,
		ForceFormatting:  true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	}
	log.Formatter = &form

	newLogger := new(Logger)
	newLogger.Logger = log
	newLogger.Prefix = prefix
	newLogger.SuccessHandler = successHandler
	newLogger.ErrorHandler = errorHandler

	return newLogger
}

// Type that encapsulate work with logger INFO/ERROR message
// with custom prefix and other structure fields.
type Logger struct {
	Logger         *logrus.Logger
	Prefix         string
	SuccessHandler map[int]string
	ErrorHandler   map[int]string
}

// Print INFO LEVEL message.
func (self *Logger) Info(code int, args ...interface{}) {
	if !isLog {
		return
	}

	message, okMessage := self.SuccessHandler[code]
	if !okMessage {
		return
	}

	var fields map[string]interface{}
	if len(args) > 0 {
		fields = args[0].(map[string]interface{})
		self.Logger.WithFields(logrus.Fields(self.makeFieldsWithPrefix(code, fields))).Info(message)
	} else {
		self.Logger.WithFields(logrus.Fields{
			"prefix": self.Prefix,
			"code":   code,
		}).Info(message)
	}
}

// Print ERROR LEVEL message.
func (self *Logger) Error(code int, args ...interface{}) {
	if !isLog {
		return
	}

	message, okMessage := self.ErrorHandler[code]
	if !okMessage {
		return
	}

	var fields map[string]interface{}
	if len(args) > 0 {
		fields = args[0].(map[string]interface{})
		self.Logger.WithFields(logrus.Fields(self.makeFieldsWithPrefix(code, fields))).Error(message)
	} else {
		self.Logger.WithFields(logrus.Fields{
			"prefix": self.Prefix,
			"code":   code,
		}).Info(message)
	}
}

// Get logger with one custom field - prefix.
func (self *Logger) GetPrefixedLogger() *logrus.Logger {
	return self.Logger.WithFields(logrus.Fields(map[string]interface{}{
		"prefix": self.Prefix,
	})).Logger
}

// Get result fields with custom prefix for prepare logger.
func (self *Logger) makeFieldsWithPrefix(code int, fields map[string]interface{}) map[string]interface{} {
	newFields := make(map[string]interface{})
	for k, v := range fields {
		newFields[k] = v
	}
	newFields["prefix"] = self.Prefix
	newFields["code"] = code
	return newFields
}
