package storage

import "synergy/internal/types"

type Storage interface {
	CreateUser(u *types.User) error
	CheckEmail(email string) (bool, error)
	// GetUserByEmail(email string) (*types.User, error)
	// GetUserByID(id int64) (*types.User, error)
	// // Project CRUD

	CreateProject(p *types.Project) error
	// GetProjectByID(id int64) (*types.Project, error)
	// ListProjects() ([]*types.Project, error)
	// UpdateProject(p *types.Project) error
	// DeleteProject(id int64) error

	// // Project members
	// AddProjectMember(projectID, userID int64) error
	// RemoveProjectMember(projectID, userID int64) error
	// ListProjectMembers(projectID int64) ([]*types.User, error)

	// // Task CRUD
	// CreateTask(t *types.Task) error
	// GetTaskByID(id int64) (*types.Task, error)
	// ListTasks(projectID int64) ([]*types.Task, error)
	// UpdateTask(t *types.Task) error
	// DeleteTask(id int64) error

	// // Task status
	// UpdateTaskStatus(taskID int64, status string) error
	// // Project chat
	// AddMessage(m *types.Message) error
	// ListMessages(projectID int64) ([]*types.Message, error)
}
