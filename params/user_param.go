package params

type CreateUser struct {
	Age      uint   `json:"age" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Username string `json:"Username" validate:"required"`
}
