package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Запрашиваем точное время с использованием библиотеки NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Fatalf("Ошибка при получении точного времени: %v", err)
	}

	// Выводим текущее время и точное время
	fmt.Println("Текущее локальное время:", time.Now().Format(time.RFC3339))
	fmt.Println("Точное время с использованием NTP:", ntpTime.Format(time.RFC3339))
}
