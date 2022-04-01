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

func (ur *UserRepository) Save(entity *model.User) (*model.User, error) {
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

func (ur *UserRepository) Update(userId string, entity *model.User) *model.User {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error during updated user :", err)
		}
	}()

	ur.Connection.Query("UPDATE users SET name=? WHERE id=?", entity.Name, userId)
	return ur.FindById(&userId)
}

func (ur *UserRepository) DeleteById(userId string) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error during delete user :", r)
		}
	}()

	_, err := ur.Connection.Query("DELETE FROM users WHERE id=?", userId)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (ur *UserRepository) FindAll() (users []model.User, err error) {
	err = ur.Connection.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Println("Error during get all data users :", err)
		return nil, err
	} else {
		return users, nil
	}
}
