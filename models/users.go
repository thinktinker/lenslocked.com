package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null; unique_index"`
}

type UserService struct {
	db *gorm.DB
}

//Lookup user by ID provided
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, OtherErr
var (
	ErrNotFound = errors.New("Resource Not Found")
)

// ByID is a method of UserService to query the database by userid
func (usr *UserService) ByID(id uint) (*User, error) {

	var user User
	err := usr.db.Where("id=?", id).First(&user).Error

	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}

// NewUserService is a function that opens a new DB connection to postgres
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{db: db}, nil
}

// Close defers the disconnection to the UserService's DB
func (usr *UserService) Close() error {
	return usr.db.Close()
}

// ResetDB drops the user table and rebuilds it
func (usr *UserService) ResetDB() {
	usr.db.DropTableIfExists(&User{})
	usr.db.AutoMigrate(&User{})
}

// CreateRecord will create the provided user and
// backfill data like ID, CreatedAt and UpdatedAt fields
func (usr *UserService) CreateRecord(user *User) error {
	return usr.db.Create(user).Error
}
