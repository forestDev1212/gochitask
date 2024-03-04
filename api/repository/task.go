package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Kbgjtn/notethingness-api.git/api/model"
	"github.com/Kbgjtn/notethingness-api.git/types"
	"github.com/lib/pq"
)

type TaskRepository struct {
	store *sql.DB
}

func NewTaskRepo(store *sql.DB) *TaskRepository {
	return &TaskRepository{store}
}

func (r TaskRepository) List(
	ctx context.Context,
	args *types.Pageable,
) (model.Tasks, error) {
	query := `SELECT *, COUNT(*) OVER() AS total FROM "tasks" ORDER BY "id" LIMIT $1 OFFSET $2`
	rows, err := r.store.QueryContext(ctx, query, args.Limit, args.Offset)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var tasks model.Tasks

	for rows.Next() {
		var task model.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Priority,
			&task.Date,
			&task.CreatedAt, &task.UpdatedAt,
			&args.Total,
		)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		args.Calc()
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r TaskRepository) Get(
	ctx context.Context,
	args model.TaskURLParams,
) (model.Task, error) {
	var task model.Task

	query := `SELECT * FROM "tasks" WHERE "id" = $1 LIMIT 1`

	row, err := r.store.QueryContext(ctx, query, args.ID)
	if err != nil {
		return task, err
	}

	if row.Next() {
		if err = row.Scan(&task.ID, &task.Title, &task.Priority, &task.Date, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return task, err
		}
	}

	return task, nil
}

func (r TaskRepository) Create(c context.Context, title string, priority int, date time.Time) (model.Task, error) {
	query := `INSERT INTO "tasks" ("title", "priority", "date") VALUES ($1, $2, $3) RETURNING *`
	row := r.store.QueryRowContext(c, query, title, priority, date)

	var task model.Task

	if err := row.Scan(&task.ID, &task.Title, &task.Priority, &task.Date, &task.CreatedAt, &task.UpdatedAt); err != nil {
		pqErr, ok := err.(*pq.Error)

		if ok && pqErr.Constraint == "tasks_title_key" {
			return model.Task{}, fmt.Errorf(
				"error: task with title %s already exists", title,
			)
		}

		return model.Task{}, err
	}

	return task, nil
}

func (r TaskRepository) Delete(c context.Context, args model.TaskURLParams) error {
	query := `DELETE FROM "tasks" WHERE "id" = $1`
	_, err := r.store.ExecContext(c, query, args.ID)
	if err != nil {
		return err
	}

	return nil
}

func (db TaskRepository) Update(
	c context.Context, args model.TaskURLParams, payload model.Task,
) (model.Task, error) {
	query := `UPDATE "tasks" SET "title" = $1, "priority" = $2, "date" = $3 WHERE "id" = $4 RETURNING *`

	row := db.store.QueryRowContext(
		c, query,
		payload.Title, payload.Priority, payload.Date, payload.ID,
	)
	var task model.Task

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Priority,
		&task.Date,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		log.Println(row)
		return task, fmt.Errorf("error: task with \"id\" %d not found", payload.ID)
	}
	return task, nil

}
