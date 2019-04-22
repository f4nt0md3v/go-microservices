package logger

import "go-microservices/libs/errors_handler"

var logger_storage_service *Logger

// Return STORAGE SERVICE Logger.
func GetStorageService() *Logger {
	if logger_storage_service == nil {

		resolveHttpErrors := map[int]string{}
		for k, v := range StorageServiceErrors {
			resolveHttpErrors[k] = v
		}
		for k, v := range errors_handler.Handler {
			if k >= 1000 && k < 2000 {
				resolveHttpErrors[k] = v
			}
		}

		logger_storage_service = NewLogger("STORAGE-SERVICE", StorageServiceSuccess, resolveHttpErrors)
	}
	return logger_storage_service
}

var StorageServiceSuccess map[int]string = map[int]string{
	10: "Storage service start",
}

var StorageServiceErrors map[int]string = map[int]string{
	10: "Error start storage service",
}
