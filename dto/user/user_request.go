package userdto

type CreateUserRequest struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" from:"name" validate:"required"`
	Email    string `json:"email" from:"email" validate:"required"`
	Password string `json:"password" from:"password" validate:"required"`
	Phone    string `json:"phone" from:"phone" validate:"required"`
	Address  string `json:"address" from:"address" validate:"required"`
}
