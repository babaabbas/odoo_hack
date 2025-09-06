package types

import "time"

type CreateUserReq struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
type CreateProjectReq struct {
	Name      string `json:"name" validate:"required,min=2,max=200"`
	CreatedBy int    `json:"owner_id" validate:"required"`
}

type CreateTaskReq struct {
	ProjectID  int    `json:"project_id" validate:"required"`
	Name       string `json:"name" validate:"required,min=2,max=200"`
	Status     string `json:"status" validate:"required,oneof=todo in_progress done"`
	AssigneeID *int   `json:"assignee_id,omitempty"`
}
type User struct {
	ID           int64     `json:"id" db:"id" validate:"required,uuid4"`
	Name         string    `json:"name" db:"name" validate:"required,min=2,max=100"`
	Email        string    `json:"email" db:"email" validate:"required,email"`
	PasswordHash string    `json:"-" db:"password_hash" validate:"required,min=8"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
type Project struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required,min=2,max=200"`
	OwnerID   int       `json:"owner_id" db:"owner_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
type Task struct {
	ID         int       `json:"id" db:"id"`
	ProjectID  int       `json:"project_id" db:"project_id" validate:"required"`
	Name       string    `json:"name" db:"name" validate:"required,min=2,max=200"`
	Status     string    `json:"status" db:"status" validate:"required,oneof=todo in_progress done"`
	AssignedTo *int      `json:"assigned_to,omitempty" db:"assigned_to"` // nullable
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type Message struct {
	ID        string    `json:"id" db:"id" validate:"required,uuid4"`
	ProjectID string    `json:"project_id" db:"project_id" validate:"required,uuid4"`
	SenderID  string    `json:"sender_id" db:"sender_id" validate:"required,uuid4"`
	Text      string    `json:"text" db:"text" validate:"required,min=1,max=2000"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
