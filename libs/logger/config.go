package logger

var logger_config *Logger

// Return CONFIG Logger.
func GetConfig() *Logger {
	if logger_config == nil {
		logger_config = NewLogger("CONFIG", ConfigSuccesss, ConfigErrors)
	}
	return logger_config
}

var ConfigSuccesss map[int]string = map[int]string{
	10: "Successful init",
	20: "Get INT value",
	30: "Get INT64 value",
	40: "Get FLOAT64 value",
	50: "Get BOOl value",
	60: "Get STRING value",
}

var ConfigErrors map[int]string = map[int]string{}
