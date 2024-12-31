package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// InitLogger menginisialisasi Logrus logger
func InitLogger(logFilePath string) {
	// Set format log
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Output ke file
	// file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatalf("Failed to open log file: %v", err)
	// }

	// log.SetOutput(file)
	log.SetLevel(logrus.InfoLevel)
	log.Info("Logger initialized")
}

// Info mencatat pesan informasi
func Info(message string) {
	log.WithFields(logrus.Fields{}).Info(message)
}

// Error mencatat pesan error
func Error(message string) {
	log.WithFields(logrus.Fields{}).Error(message)
}

// Fatal mencatat pesan error dan menghentikan aplikasi
func Fatal(message string) {
	log.WithFields(logrus.Fields{}).Fatal(message)
}
