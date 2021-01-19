package services

import (
	"fmt"
	"go-api-hexagonal/domain"
)

func CreateUser(u *domain.User) {
	fmt.Println(u.Name)
}

// func FindUsers() []*domain.User {

// }

// func FindUser(id int) *domain.User {

// }

// func DeleteUser(id int) *domain.User {

// }
