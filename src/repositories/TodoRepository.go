package repositories

import (
	"log"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-graphql-go/graph/model"
	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
	Connection *sqlx.DB
}

func (td *TodoRepository) Save(entity *model.Todo) (*model.Todo, error) {
	query := "INSERT INTO todo (id, todo_text, user_id) VALUES (?, ?, ?)"
	entity.ID = uuid.NewString()
	_, err := td.Connection.Query(query, entity.ID, entity.Text, entity.User.ID)

	if err != nil {
		log.Println("Error during save new Todo :", err)
		return nil, err
	} else {
		return td.FindById(&entity.ID), nil
	}
}

func (td *TodoRepository) FindById(id *string) *model.Todo {
	todo := model.Todo{}
	query := "SELECT * FROM todo WHERE id=?"
	err := td.Connection.Get(&todo, query, *id)

	if err != nil {
		return nil
	} else {
		return &todo
	}
}

func (td *TodoRepository) Update(ID string, entity *model.Todo) *model.Todo {
	query := "UPDATE todo SET todo.todo_text=?, todo.done=? where id=?"
	_, err := td.Connection.Query(query, entity.Text, entity.Done, ID)
	if err != nil {
		return nil
	} else {
		return td.FindById(&ID)
	}
}

func (td *TodoRepository) DeleteById(ID string) error {
	query := "DELETE FROM todo WHERE id=?"
	_, err := td.Connection.Query(query, ID)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (td *TodoRepository) FindAll() (todos []model.Todo, err error) {
	err = td.Connection.Select(&todos, "SELECT * FROM todo")

	if err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}
