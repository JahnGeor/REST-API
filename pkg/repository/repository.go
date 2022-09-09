package repository

import (
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetByID(userID int, id int) (todo.TodoList, error)
	Delete(userID int, id int) error
	Update(userID int, id int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listID int, item todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetByID(userID, itemID int) (todo.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID int, itemID int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
