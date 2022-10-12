package database

import (
	"errors"
	"mvcgolang/app/model"

	"golang.org/x/crypto/bcrypt"
)

func SaveUser(user model.User) (model.User, error) {
	err := DB().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUser(users []model.User) ([]model.User, error) {
	err := DB().Model(&model.User{}).Preload("CreditCards").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func Login(email string, password string) (model.User, error) {
	var user model.User
	err := DB().Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil

}

func SaveFile(ID int, fileLocation string) (model.User, error) {
	var user model.User
	err := DB().Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation
	err = DB().Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func FindByID(ID int) (model.User, error) {
	var user model.User
	err := DB().Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByID(ID int) (model.User, error) {
	user, err := FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No User Found with that ID")
	}

	return user, nil
}
