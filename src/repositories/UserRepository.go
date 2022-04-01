package repositories

import (
	"log"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-graphql-go/graph/model"
	"github.com/jmoiron/sqlx"
)

const tableName = "users"

type UserRepository struct {
	Connection *sqlx.DB
}

func (ur *UserRepository) Save(entity model.User) (*model.User, error) {
	log.Println("Saving entity user...")

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error during generate UUID :", r)
		}
	}()

	entity.ID = uuid.NewString()
	_, err := ur.Connection.Queryx("INSERT INTO "+tableName+" (id, name) VALUES (?, ?);", entity.ID, entity.Name)

	if err != nil {
		return nil, err
	} else {
		return ur.FindById(&entity.ID), nil
	}
}

func (ur *UserRepository) FindById(ID *string) *model.User {
	user := model.User{}
	err := ur.Connection.Get(&user, "SELECT * FROM users WHERE id=?", *ID)

	if err != nil {
		log.Println("Error when find user by id :", err.Error())
		return nil
	} else {
		log.Println("Finding user by id...")
		return &user
	}
}
