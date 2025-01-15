package authmodel

type OTP struct {
	Code  string `json:"code" validate:"required,min=6,max=6"`
	Phone string `json:"phone" validate:"required,e164"`
}

func (o *OTP) Validate() error {
	return validateStruct(o)
}

type ResendOTP struct {
	Phone string `json:"phone" validate:"required,e164"`
}

func (ro *ResendOTP) Validate() error {
	return validateStruct(ro)
}
