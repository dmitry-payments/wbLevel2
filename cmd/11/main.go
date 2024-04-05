package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Event представляет собой событие в календаре.
type Event struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Date      time.Time `json:"date"`
}

// Calendar представляет собой календарь, содержащий список событий.
type Calendar struct {
	Events map[int]*Event
}

// JSONError представляет ошибку в формате JSON.
type JSONError struct {
	Message string `json:"message"`
}

// Сериализация объектов доменной области в JSON
func (e *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":         e.ID,
		"title":      e.Title,
		"start_time": e.StartTime.Format(time.RFC3339),
		"end_time":   e.EndTime.Format(time.RFC3339),
		"date":       e.EndTime.Format(time.RFC3339),
	})
}

// Парсинг и валидация параметров методов /create_event и /update_event
func parseEventParams(r *http.Request) (*Event, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	// Получаем значения полей из формы
	id, _ := strconv.Atoi(r.Form.Get("id"))
	title := r.Form.Get("title")
	startTime, _ := time.Parse(time.RFC3339, r.Form.Get("start_time"))
	endTime, _ := time.Parse(time.RFC3339, r.Form.Get("end_time"))
	date, _ := time.Parse(time.RFC3339, r.Form.Get("date"))

	// Создаем объект Event
	event := &Event{
		ID:        id,
		Title:     title,
		StartTime: startTime,
		EndTime:   endTime,
		Date:      date,
	}

	fmt.Println(event)

	return event, nil
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	event, err := parseEventParams(r)
	if err != nil {
		http.Error(w, "Failed to parse event parameters", http.StatusBadRequest)
		return
	}

	// Здесь можно добавить логику для создания события в календаре
	calendar.Events[event.ID] = event

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusCreated)
	fmt.Println(calendar)
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	event, err := parseEventParams(r)
	if err != nil {
		http.Error(w, "Failed to parse event parameters", http.StatusBadRequest)
		return
	}

	// Здесь можно добавить логику для создания события в календаре
	if eve, ok := calendar.Events[event.ID]; ok {
		eve.Title = event.Title
		eve.StartTime = event.StartTime
		eve.EndTime = event.EndTime
		eve.Date = event.Date
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println(calendar)
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	event, err := parseEventParams(r)
	if err != nil {
		http.Error(w, "Failed to parse event parameters", http.StatusBadRequest)
		return
	}

	if _, ok := calendar.Events[event.ID]; ok {
		delete(calendar.Events, event.ID)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println(calendar)
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var events []*Event
	//dateStr := r.URL.Query().Get("date")
	date, _ := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	for _, event := range calendar.Events {
		if event.Date == date {
			events = append(events, event)
		}
	}

	jsonData, _ := json.Marshal(events)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed", http.StatusBadRequest)
		return
	}

}

// Метод middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

var calendar = Calendar{
	Events: map[int]*Event{},
}

func main() {
	// Создаем новый календарь
	//calendar := &Calendar{}

	fmt.Println(calendar)

	// Регистрируем HTTP обработчики
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)

	// Добавляем middleware для логирования
	http.Handle("/", loggingMiddleware(http.DefaultServeMux))

	// Запускаем HTTP сервер
	port := ":8080"
	fmt.Println("Server is listening on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
