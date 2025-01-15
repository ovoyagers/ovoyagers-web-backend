package authmodel

type LoginRequest struct {
	CountryCode string `json:"countryCode" validate:"required,iso3166_1_alpha2"`
	Phone       string `json:"phone" validate:"required,e164"`
}

func (r *LoginRequest) Validate() error {
	return validateStruct(r)
}
