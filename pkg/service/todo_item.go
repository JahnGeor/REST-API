package service

import (
	todo "todo-app"
	"todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userID int, listID int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		//листа с данным id у данного пользователя не существует
		return 0, err
	}
	return s.repo.Create(listID, item)
}

func (s *TodoItemService) GetAll(userID, listID int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userID, listID)
}

func (s *TodoItemService) GetByID(userID int, itemID int) (todo.TodoItem, error) {
	return s.repo.GetByID(userID, itemID)
}
