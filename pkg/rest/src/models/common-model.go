package models

type Error struct {
	Message    string `json:"message" example:"Error message"`
	Error      string `json:"error" example:"Internal Server Error"`
	Status     string `json:"status" example:"error"`
	StatusCode int    `json:"status_code" example:"500"`
	Data       any    `json:"data"`
}

type Response struct {
	Message    string `json:"message" example:"Data fetched successfully"`
	Data       any    `json:"data"`
	Status     string `json:"status" example:"success"`
	StatusCode int    `json:"status_code" example:"200"`
}

type ImgMetaData struct {
	UserId   string `json:"userId"`
	Avatar   string `json:"avatar"`
	ImageId  string `json:"imageId"`
	Filename string `json:"filename"`
}
