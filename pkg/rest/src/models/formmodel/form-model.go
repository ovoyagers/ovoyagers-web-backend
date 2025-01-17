package formmodel

type Form struct {
	Id        string `json:"id,omitempty"`
	Fullname  string `json:"fullname,omitempty" example:"John Doe"`
	Email     string `json:"email" validate:"required,email" example:"pecol35486@ovoyagers.com"`
	Mobile    string `json:"mobile" validate:"required,e164" example:"+918765432100"`
	Message   string `json:"message" validate:"required" example:"This is a test message"`
	Category  string `json:"category" validate:"required,oneof=contact hotel flights" example:"contact"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func (f *Form) Validate() error {
	return validateStruct(f)
}
