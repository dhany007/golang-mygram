package params

type CreateComment struct {
	Message string `json:"message" validate:"required"`
	PhotoID uint   `json:"photo_id" validate:"required,min=1"`
	UserID  uint   `json:"user_id,omitempty"`
}

type UpdateComment struct {
	Message string `json:"message" validate:"required"`
}
