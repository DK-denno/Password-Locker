package users

import (
	"errors"
	"fmt"
)

type Users struct {
	Username string
	Email string
	Password string
}

var UsersDb map[string]Users

func Init()  {
	UsersDb = make(map[string]Users)
}

func AddUserAccount(name string, Email string, Password string) (Users, error) {
	if _, exists := UsersDb[name]; exists {
		return Users{}, errors.New("User With Thats Username already exists")
	}

	user := Users {name, Email, Password}

	if UsersDb == nil {
		UsersDb = make(map[string]Users)
	}
	UsersDb[name] = user
	fmt.Println("[*] New Account Has been created")
	return user, nil
}

func GetUser(name string, Password string) (Users, error) {
	// fmt.Println(UsersDb[name])
	if user, exists := UsersDb[name]; exists {
		if Password == user.Password {
			return UsersDb[name], nil
		}
		return Users{}, fmt.Errorf("Incorrect Password")
	}

	return Users{}, fmt.Errorf("User matching name %s was not found", name)
}

func (u *Users) EditUserDetails(name string, key string, value string,
	oldValue string) (Users, error) {
	if user, exists := UsersDb[name]; exists {
		switch key {
			case "Username":
				user.Username = value
				UsersDb[value] = user
				delete(UsersDb, oldValue)
			case "Email":
				user.Email = value
				UsersDb[name] = user

			case "Password":
				user.Password = value
				UsersDb[name] = user
		}
		return user, nil
	}

	return Users{}, fmt.Errorf("User matching name %s was not found", name)
}

func (u *Users) ViewUser(name string) error {
	if user, exists := UsersDb[name]; exists {
		fmt.Println("Here are your profile details :- ")
		fmt.Sprintln("Username :- %s", user.Username)
		fmt.Sprintln("Email :- %s", user.Email)

		return nil
	}
	return fmt.Errorf("User Not Found")
}