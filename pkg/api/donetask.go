package api

import (
	"net/http"
	"time"

	"github.com/GeaRvant/go_final_project/pkg/db"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "ID is required", http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJson(w, map[string]any{})
		return
	}

	next, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateDate(next, id)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJson(w, map[string]any{})
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJson(w, map[string]any{})
}
