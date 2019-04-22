package logger

import "go-microservices/libs/errors_handler"

var logger_client_service *Logger

// Return CLIENT SERVICE Logger.
func GetClientService() *Logger {
	if logger_client_service == nil {

		resolveHttpErrors := map[int]string{}
		for k, v := range ClientServiceErrors {
			resolveHttpErrors[k] = v
		}
		for k, v := range errors_handler.Handler {
			if k >= 1000 && k < 2000 {
				resolveHttpErrors[k] = v
			}
		}

		logger_client_service = NewLogger("CLIENT-SERVICE", ClientServiceSuccess, resolveHttpErrors)
	}
	return logger_client_service
}

var ClientServiceSuccess map[int]string = map[int]string{
	10: "Client service start",
}

var ClientServiceErrors map[int]string = map[int]string{
	10: "Error start client service",
}
