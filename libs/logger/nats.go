package logger

var logger_nats *Logger

// Return NATS Logger.
func GetNats() *Logger {
	if logger_nats == nil {
		logger_nats = NewLogger("NATS", NatsSuccess, NatsErrors)
	}
	return logger_nats
}

var NatsSuccess map[int]string = map[int]string{
	10: "Success subscribe",
	20: "Success publish",
	30: "Success get message",
}

var NatsErrors map[int]string = map[int]string{
	10: "Error connection",
	20: "Error subscribe",
	30: "Error parse pb",
	40: "Error publish",
}
