package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/Shahboz4131/7-project/genproto"
)

type taskRepo struct {
	db *sqlx.DB
}

// NewTaskRepo ...
func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task pb.Task) (pb.Task, error) {
	var id string
	err := r.db.QueryRow(`
        INSERT INTO tasks(id, title, summary, deadline, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6) returning id`, task.Id, task.Title, task.Summary, task.Deadline, time.Now().UTC(), time.Now().UTC()).Scan(&id)
	if err != nil {
		return pb.Task{}, err
	}

	task, err = r.Get(id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Get(id string) (pb.Task, error) {
	var task pb.Task

	err := r.db.QueryRow(`
        SELECT id::varchar, title, summary, deadline, created_at, updated_at FROM tasks
        WHERE id=$1 and deleted_at is  null`, id).Scan(&task.Id, &task.Title, &task.Summary, &task.Deadline, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Update(task pb.Task) (pb.Task, error) {
	result, err := r.db.Exec(`UPDATE tasks SET , title=$2, summary=$3, deadline=$4, updated_at=$5 WHERE id=$1 and deleted_at is  null`,
		task.Id, task.Title, task.Summary, task.Deadline, time.Now().UTC())
	if err != nil {
		return pb.Task{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Task{}, sql.ErrNoRows
	}

	task, err = r.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Delete(id string) error {
	result, err := r.db.Exec(`UPDATE tasks SET deleted_at = $2 where id = $1 and deleted_at is null `, id, time.Now().UTC())
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
