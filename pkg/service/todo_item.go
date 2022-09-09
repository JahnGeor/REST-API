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

func (s *TodoItemService) Delete(userID int, itemID int) error {
	return s.repo.Delete(userID, itemID)
}

func (s *TodoItemService) Update(userID int, itemID int, input todo.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userID, itemID, input)
}
