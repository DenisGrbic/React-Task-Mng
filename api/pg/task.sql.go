// Code generated by sqlc. DO NOT EDIT.
// source: task.sql

package pg

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO task (task_group_id, created_at, name, position)
  VALUES($1, $2, $3, $4) RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete
`

type CreateTaskParams struct {
	TaskGroupID uuid.UUID `json:"task_group_id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Position    float64   `json:"position"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.TaskGroupID,
		arg.CreatedAt,
		arg.Name,
		arg.Position,
	)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
	)
	return i, err
}

const deleteTaskByID = `-- name: DeleteTaskByID :exec
DELETE FROM task WHERE task_id = $1
`

func (q *Queries) DeleteTaskByID(ctx context.Context, taskID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTaskByID, taskID)
	return err
}

const deleteTasksByTaskGroupID = `-- name: DeleteTasksByTaskGroupID :execrows
DELETE FROM task where task_group_id = $1
`

func (q *Queries) DeleteTasksByTaskGroupID(ctx context.Context, taskGroupID uuid.UUID) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteTasksByTaskGroupID, taskGroupID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getAllTasks = `-- name: GetAllTasks :many
SELECT task_id, task_group_id, created_at, name, position, description, due_date, complete FROM task
`

func (q *Queries) GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.TaskID,
			&i.TaskGroupID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
			&i.Description,
			&i.DueDate,
			&i.Complete,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT task_id, task_group_id, created_at, name, position, description, due_date, complete FROM task WHERE task_id = $1
`

func (q *Queries) GetTaskByID(ctx context.Context, taskID uuid.UUID) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, taskID)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
	)
	return i, err
}

const getTasksForTaskGroupID = `-- name: GetTasksForTaskGroupID :many
SELECT task_id, task_group_id, created_at, name, position, description, due_date, complete FROM task WHERE task_group_id = $1
`

func (q *Queries) GetTasksForTaskGroupID(ctx context.Context, taskGroupID uuid.UUID) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTasksForTaskGroupID, taskGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.TaskID,
			&i.TaskGroupID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
			&i.Description,
			&i.DueDate,
			&i.Complete,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setTaskComplete = `-- name: SetTaskComplete :one
UPDATE task SET complete = $2 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete
`

type SetTaskCompleteParams struct {
	TaskID   uuid.UUID `json:"task_id"`
	Complete bool      `json:"complete"`
}

func (q *Queries) SetTaskComplete(ctx context.Context, arg SetTaskCompleteParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, setTaskComplete, arg.TaskID, arg.Complete)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
	)
	return i, err
}

const updateTaskDescription = `-- name: UpdateTaskDescription :one
UPDATE task SET description = $2 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete
`

type UpdateTaskDescriptionParams struct {
	TaskID      uuid.UUID      `json:"task_id"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) UpdateTaskDescription(ctx context.Context, arg UpdateTaskDescriptionParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskDescription, arg.TaskID, arg.Description)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
	)
	return i, err
}

const updateTaskDueDate = `-- name: UpdateTaskDueDate :one
UPDATE task SET due_date = $2 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete
`

type UpdateTaskDueDateParams struct {
	TaskID  uuid.UUID    `json:"task_id"`
	DueDate sql.NullTime `json:"due_date"`
}

func (q *Queries) UpdateTaskDueDate(ctx context.Context, arg UpdateTaskDueDateParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskDueDate, arg.TaskID, arg.DueDate)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
	)
	return i, err
}

const updateTaskLocation = `-- name: UpdateTaskLocation :one
UPDATE task SET task_group_id = $2, position = $3 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete
`

type UpdateTaskLocationParams struct {
	TaskID      uuid.UUID `json:"task_id"`
	TaskGroupID uuid.UUID `json:"task_group_id"`
	Position    float64   `json:"position"`
}

func (q *Queries) UpdateTaskLocation(ctx context.Context, arg UpdateTaskLocationParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskLocation, arg.TaskID, arg.TaskGroupID, arg.Position)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
	)
	return i, err
}

const updateTaskName = `-- name: UpdateTaskName :one
UPDATE task SET name = $2 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete
`

type UpdateTaskNameParams struct {
	TaskID uuid.UUID `json:"task_id"`
	Name   string    `json:"name"`
}

func (q *Queries) UpdateTaskName(ctx context.Context, arg UpdateTaskNameParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskName, arg.TaskID, arg.Name)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
	)
	return i, err
}
