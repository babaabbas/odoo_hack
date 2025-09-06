package types

import "time"

type User struct {
	ID           string    `json:"id" db:"id" validate:"required,uuid4"`
	Name         string    `json:"name" db:"name" validate:"required,min=2,max=100"`
	Email        string    `json:"email" db:"email" validate:"required,email"`
	PasswordHash string    `json:"-" db:"password_hash" validate:"required,min=8"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type Project struct {
	ID        string    `json:"id" db:"id" validate:"required,uuid4"`
	Name      string    `json:"name" db:"name" validate:"required,min=2,max=200"`
	CreatedBy string    `json:"created_by" db:"created_by" validate:"required,uuid4"`
	Members   []string  `json:"members" db:"members" validate:"dive,uuid4"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Task struct {
	ID          string    `json:"id" db:"id" validate:"required,uuid4"`
	ProjectID   string    `json:"project_id" db:"project_id" validate:"required,uuid4"`
	Title       string    `json:"title" db:"title" validate:"required,min=2,max=200"`
	Description string    `json:"description" db:"description" validate:"max=1000"`
	AssigneeID  string    `json:"assignee_id" db:"assignee_id" validate:"omitempty,uuid4"`
	DueDate     time.Time `json:"due_date" db:"due_date" validate:"required"`
	Status      string    `json:"status" db:"status" validate:"required,oneof=todo in_progress done"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Message struct {
	ID        string    `json:"id" db:"id" validate:"required,uuid4"`
	ProjectID string    `json:"project_id" db:"project_id" validate:"required,uuid4"`
	SenderID  string    `json:"sender_id" db:"sender_id" validate:"required,uuid4"`
	Text      string    `json:"text" db:"text" validate:"required,min=1,max=2000"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
