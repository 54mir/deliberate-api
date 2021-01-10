package main

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	_id      uuid.UUID
	name     string
	email    string
	password string
}

type Project struct {
	_id           uuid.UUID
	user_id       uuid.UUID
	name          string
	dateCompleted time.Time
	password      string
}

type Task struct {
	_id           uuid.UUID
	project_id    uuid.UUID
	name          string
	dateCompleted time.Time
	lastAssigned  time.Time
}

type WorkLog struct {
	_id       uuid.UUID
	task_id   uuid.UUID
	startTime time.Time
	endTime   time.Time
}

type UserStore interface {
	User(_id uuid.UUID) (User, error)
	CreateUser(t *User) error
	UpdateUser(t *User) error
	DeleteUser(t *User) error
}

type ProjectStore interface {
	Project(_id uuid.UUID) (Project, error)
	ProjectsByUser(user_id uuid.UUID) ([]Project, error)
	CreateProject(t *Project) error
	UpdateProject(t *Project) error
	DeleteProject(t *Project) error
}

type TaskStore interface {
	Task(_id uuid.UUID) (Task, error)
	TasksByProject(project_id uuid.UUID) ([]Task, error)
	CreateTask(t *Task) error
	UpdateTask(t *Task) error
	DeleteTask(t *Task) error
}

type WorkLogStore interface {
	WorkLog(_id uuid.UUID) (WorkLog, error)
	WorkLogsByTask(task_id uuid.UUID) ([]WorkLog, error)
	CreateWorkLog(t *WorkLog) error
}
