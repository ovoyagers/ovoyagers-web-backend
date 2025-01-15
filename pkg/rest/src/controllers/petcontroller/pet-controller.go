package petcontroller

import (
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/services/petservice"
)

type PetController struct {
	petService *petservice.PetService
}

func NewPetController(globalCfg *config.GlobalConfig) *PetController {
	return &PetController{
		petService: petservice.NewPetService(globalCfg),
	}
}
