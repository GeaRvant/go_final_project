package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/GeaRvant/go_final_project/pkg/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(dateLayout)
	}

	t, err := time.Parse(dateLayout, task.Date)
	if err != nil {
		return err
	}

	var next string

	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if !afterNow(t, now) && t.Format(dateLayout) != now.Format(dateLayout) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(dateLayout)
		} else {
			task.Date = next
		}
	}
	return nil
}

func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeError(w, "Task title is required", http.StatusBadRequest)
		return
	}

	if err := checkDate(&task); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJson(w, map[string]string{
		"id": strconv.FormatInt(id, 10),
	})
}
