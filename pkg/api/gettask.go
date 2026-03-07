package api

import (
	"net/http"
	"strconv"

	"github.com/GeaRvant/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	type taskDTO struct {
		ID      string `json:"id"`
		Date    string `json:"date"`
		Title   string `json:"title"`
		Comment string `json:"comment"`
		Repeat  string `json:"repeat"`
	}

	resp := []taskDTO{}

	for _, t := range tasks {
		resp = append(resp, taskDTO{
			ID:      strconv.FormatInt(t.ID, 10),
			Date:    t.Date,
			Title:   t.Title,
			Comment: t.Comment,
			Repeat:  t.Repeat,
		})
	}

	writeJson(w, map[string]any{
		"tasks": resp,
	})
}
