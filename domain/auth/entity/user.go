package entity

import (
	"time"

	"github.com/arvinpaundra/go-boilerplate/model"
	"github.com/guregu/null/v6"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64      `json:"id" redis:"id"`
	Email     string     `json:"email" redis:"email"`
	Password  *string    `json:"password" redis:"-"`
	Fullname  string     `json:"fullname" redis:"-"`
	Image     *string    `json:"image" redis:"-"`
	DeletedAt *time.Time `json:"deleted_at" redis:"-"`
}

func (e *User) GeneratePassword(password string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	hashed := string(b)

	e.Password = &hashed

	return nil
}

func (e *User) ComparePassword(password string) bool {
	if e.Password != nil {
		err := bcrypt.CompareHashAndPassword([]byte(*e.Password), []byte(password))

		return err == nil
	}

	return false
}

func (e *User) IsEmpty() bool {
	return *e == (User{})
}

func (e *User) ToModel() model.User {
	return model.User{
		ID:       e.ID,
		Email:    e.Email,
		Password: null.StringFromPtr(e.Password),
		Fullname: e.Fullname,
		Image:    null.StringFromPtr(e.Image),
	}
}
