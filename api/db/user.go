package db

import (
	"fmt"
	"log"
	"vijay/project-1/api/model"

	"golang.org/x/crypto/bcrypt"
)

type User interface {
	Get(int) (*model.User, error)
	Insert(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete(int) (int, error)
	Login(string, string) (*model.User, error)
	GetAll() ([]*model.User, error)
}

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (u *UserDao) Insert(data *model.User) (*model.User, error) {
	result := db.Create(data)
	return data, result.Error
}

func (u *UserDao) Get(id int) (*model.User, error) {
	user := &model.User{}
	log.Println("ID is :", id)
	result := db.First(&user, id)
	return user, result.Error
}

func (u *UserDao) Update(data *model.User) (*model.User, error) {
	result := db.Save(&data)
	data.Password = ""
	return data, result.Error
}

func (u *UserDao) Delete(id int) (int, error) {
	user := &model.User{}
	result := db.Delete(&user, id)
	return int(user.Id), result.Error
}

func (u *UserDao) Login(UserName string, Password string) (*model.User, error) {
	data := &model.User{}
	result := db.Where("email = ? or mobile = ?", UserName, UserName).First(&data)
	if result.Error != nil {
		fmt.Println("Given username is not matched", result.Error)
	}

	err := bcrypt.CompareHashAndPassword([]byte(Password), []byte(data.Password))

	if err != nil {
		fmt.Println("Given password is not matched", err)
	}
	data.Password = ""
	return data, nil
}

func (u *UserDao) GetAll() ([]*model.User, error) {
	data := []*model.User{}
	result := db.Find(&data)

	if result.Error != nil {
		for i := 0; i < len(data); i++ {
			data[i].Password = ""
		}
	}
	return data, nil
}
