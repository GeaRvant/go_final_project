package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeaRvant/go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "ID is required", http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, "Task not found", http.StatusBadRequest)
		return
	}

	writeJson(w, map[string]string{
		"id":      strconv.FormatInt(task.ID, 10),
		"date":    task.Date,
		"title":   task.Title,
		"comment": task.Comment,
		"repeat":  task.Repeat,
	})
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID      string `json:"id"`
		Date    string `json:"date"`
		Title   string `json:"title"`
		Comment string `json:"comment"`
		Repeat  string `json:"repeat"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if req.ID == "" {
		writeError(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		writeError(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		writeError(w, "Task title is required", http.StatusBadRequest)
		return
	}

	task := db.Task{
		ID:      id,
		Date:    req.Date,
		Title:   req.Title,
		Comment: req.Comment,
		Repeat:  req.Repeat,
	}

	if err := checkDate(&task); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeError(w, "Task not found", http.StatusBadRequest)
		return
	}

	writeJson(w, map[string]any{})
}
