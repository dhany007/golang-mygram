package params

type CreateUser struct {
	Age      uint   `json:"age" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Username string `json:"Username" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"Username" validate:"required"`
}
