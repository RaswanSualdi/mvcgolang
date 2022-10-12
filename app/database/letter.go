package database

import (
	"errors"
	"mvcgolang/app/model"
)

func GetLetter(uuid string) (model.Letter, error) {
	var letters model.Letter
	err := DB().Where("uuid = ?", uuid).Find(&letters).Error
	if err != nil {
		return letters, err
	}
	if letters.Uuid == "" {
		return letters, errors.New("No Letter Found")
	}

	return letters, nil

}

func SaveLetter(letter model.Letter) (model.Letter, error) {
	err := DB().Create(&letter).Error
	if err != nil {
		return letter, err
	}

	return letter, nil
}



