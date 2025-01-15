package invitemodel

type InviteUser struct {
	Id        string `json:"id"`
	Email     string `json:"email" validate:"required,email" example:"test@example.com"`	
	Status    string `json:"status" oneof:"pending accepted declined" default:"pending" example:"pending"`
	ExpiresAt string `json:"expires_at,omitempty" example:"2021-01-01T00:00:00Z"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (i *InviteUser) Validate() error {
	return validateStruct(i)
}
