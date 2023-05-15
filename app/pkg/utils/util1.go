package utils

import "fmt"

type Logger struct {
	// Logger için alanlar burada tanımlanır
	//...
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) LogInfo(message string) {
	// Bilgi mesajını işleme mantığı burada yer alır
	fmt.Println(message)
}

func (l *Logger) LogError(err error) {
	// Hata mesajını işleme mantığı burada yer alır
	fmt.Println(err)
}
