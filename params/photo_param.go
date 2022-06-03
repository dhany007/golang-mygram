package params

type CreateUpdatePhoto struct {
	UserID   uint   `json:"user_id,omitempty"`
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required"`
}
