package authmodel

type ForgetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (fpr *ForgetPasswordRequest) Validate() (err error) {
	return validateStruct(fpr)
}
