package petmodel

type PetProfilePicture struct {
	UserId     string `json:"userId"`
	ImageBytes string `json:"imageBytes"`
	PetId      string `json:"petId"`
	FileName   string `json:"fileName"`
}
