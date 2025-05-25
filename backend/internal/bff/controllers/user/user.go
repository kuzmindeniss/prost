package user

import "github.com/google/uuid"

type UserResponse struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Email   string    `json:"email"`
	Role    string    `json:"role"`
}
