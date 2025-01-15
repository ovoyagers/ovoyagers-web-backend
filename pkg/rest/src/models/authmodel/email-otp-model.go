package authmodel

type EmailOTP struct {
	Code  string `json:"code" validate:"required,min=6,max=6" example:"123456"`
	Email string `json:"email" validate:"required,email" example:"pecol35486@kwalah.com"`
}

type ResendEmailOTP struct {
	Email string `json:"email" validate:"required,email" example:"pecol35486@kwalah.com"`
}

func (o *EmailOTP) Validate() error {
	return validateStruct(o)
}

func (ro *ResendEmailOTP) Validate() error {
	return validateStruct(ro)
}
