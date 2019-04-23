package logger

var logger_cockroach *Logger

// Return COCKROACH Logger.
func GetCockroach() *Logger {
	if logger_cockroach == nil {
		logger_cockroach = NewLogger("COCKROACH", CockroachSuccess, CockroachErrors)
	}
	return logger_cockroach
}

var CockroachSuccess map[int]string = map[int]string{
	10: "Success connect",
}

var CockroachErrors map[int]string = map[int]string{
	10: "Error connection",
	11: "Error query in db",
}
