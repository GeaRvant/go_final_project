package db

import (
	"database/sql"
	"fmt"
)

type Task struct {
	ID      int64  `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	query := fmt.Sprintf(`
	SELECT id, date, title, comment, repeat
	FROM scheduler
	ORDER BY date ASC
	LIMIT %d
`, limit)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE id = ?
	`

	var t Task

	err := db.QueryRow(query, id).Scan(
		&t.ID,
		&t.Date,
		&t.Title,
		&t.Comment,
		&t.Repeat,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Task not found")
		}
		return nil, err
	}

	return &t, nil
}

func UpdateTask(task *Task) error {
	query := `
	UPDATE scheduler
	SET date = ?, title = ?, comment = ?, repeat = ?
	WHERE id = ?
	`

	res, err := db.Exec(query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID,
	)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Incorrect id for updating task")
	}

	return nil
}

func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = ?`

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}

func UpdateDate(next string, id string) error {
	query := `
	UPDATE scheduler
	SET date = ?
	WHERE id = ?
	`

	res, err := db.Exec(query, next, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}
