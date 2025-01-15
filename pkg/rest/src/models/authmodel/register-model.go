package authmodel

type RegisterRequest struct {
	Fullname    string `json:"fullname" validate:"required,min=3,max=64" example:"Pecol"`
	Email       string `json:"email" validate:"required,email,emaildomain" example:"pecol35486@ovoyagers.com"`
	Phone       string `json:"phone" validate:"required,min=10,max=10,numeric" example:"6302068026"`
	CountryCode string `json:"countryCode" validate:"required,iso3166_1_alpha2" example:"IN"`
	Password    string `json:"password" validate:"required,min=8,max=64,pswd" example:"password123"`
}

func (r *RegisterRequest) Validate() error {
	return validateStruct(r)
}

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"required,email,emaildomain" example:"pecol35486@ovoyagers.com"`
	Code  string `json:"code" validate:"required,min=6,max=6,numeric" example:"123456"`
}

func (r *VerifyEmailRequest) Validate() error {
	return validateStruct(r)
}

type LoginEmailRequest struct {
	Email        string `json:"email" validate:"required,email,emaildomain" example:"pecol35486@kwalah.com"`
	Password     string `json:"password" validate:"omitempty,min=8,max=64,pswd" example:"password123"`
	Passwordless bool   `json:"passwordless" default:"false" example:"false"`
}

func (r *LoginEmailRequest) Validate() error {
	return validateStruct(r)
}
