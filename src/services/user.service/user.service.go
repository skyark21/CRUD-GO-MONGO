package user_service

import (
	m "../../models"
	userRepository "../../repositories/user.repository"
)

func Create(user m.Users) error {
	err := userRepository.Create(user)
	if err != nil {
		panic(err)
	}
	return nil
}

func Read() (m.Users, error) {
	users, err := userRepository.Read(user)
	if err != nil {
		panic(err)
	}
	return users, nil
}

func Update(user m.Users, userId string) error {
	err := userRepository.Update(user, userId)
	if err != nil {
		panic(err)
	}
	return nil
}

func Delete(userId string) error {
	err := userRepository.Delete(userId)
	if err != nil {
		panic(err)
	}
	return nil
}
