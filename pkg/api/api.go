package api

import (
	"fmt"
	"net/http"
	"time"
)

const dateLayout = "20060102"

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", getTasksHandler)
	http.HandleFunc("/api/task/done", taskDoneHandler)
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeat := r.FormValue("repeat")

	if dateStr == "" {
		http.Error(w, "Parameter 'date' is required", http.StatusBadRequest)
		return
	}
	if repeat == "" {
		http.Error(w, "Parameter 'repeat' is required", http.StatusBadRequest)
		return
	}

	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateLayout, nowStr)
		if err != nil {
			http.Error(w, "Invalid 'now' date format, expected YYYYMMDD", http.StatusBadRequest)
			return
		}
	}

	nextDate, err := NextDate(now, dateStr, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, nextDate)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
