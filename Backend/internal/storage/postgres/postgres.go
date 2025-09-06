package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"synergy/internal/config"
	"synergy/internal/types"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func New(cfg *config.Config) (*Postgres, error) {
	connStr := cfg.Conn_Str
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping db: %w", err)
	}
	log.Println("Connected to PostgreSQL successfully!")
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);`,
		`CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    owner_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);`,
		`CREATE TABLE IF NOT EXISTS project_members (
    project_id INT REFERENCES projects(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (project_id, user_id)
);`,
		`CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    project_id INT REFERENCES projects(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'todo',
    assigned_to INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);`,
		`CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    project_id INT REFERENCES projects(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);`,
	}
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			return &Postgres{db: db}, fmt.Errorf("failed to execute query: %w", err)
		}
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) CreateUser(u *types.User) error {
	now := time.Now()

	query := `
		INSERT INTO users (username, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := p.db.QueryRow(query, u.Name, u.Email, u.PasswordHash, now, now).Scan(&u.ID)
	if err != nil {
		return fmt.Errorf("could not insert user: %w", err)
	}

	u.CreatedAt = now
	u.UpdatedAt = now

	return nil
}

func (p *Postgres) CheckEmail(email string) (bool, error) {
	var exists bool
	err := p.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		return exists, fmt.Errorf("could not check email existence: %w", err)
	}
	if exists {
		return exists, fmt.Errorf("email already exists")
	}
	return exists, nil
}
func (p *Postgres) CreateProject(proj *types.Project) error {
	now := time.Now()

	query := `
		INSERT INTO projects (name, owner_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := p.db.QueryRow(query, proj.Name, proj.OwnerID, now, now).Scan(&proj.ID)
	if err != nil {
		return fmt.Errorf("could not insert project: %w", err)
	}

	proj.CreatedAt = now
	proj.UpdatedAt = now

	return nil
}

func (p *Postgres) GetProjectByID(id int64) (*types.Project, error) {
	query := `
		SELECT id, name, owner_id, created_at, updated_at
		FROM projects
		WHERE id = $1
	`

	var proj types.Project
	err := p.db.QueryRow(query, id).Scan(
		&proj.ID,
		&proj.Name,
		&proj.OwnerID,
		&proj.CreatedAt,
		&proj.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("project with id %d not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("could not fetch project: %w", err)
	}

	return &proj, nil
}
