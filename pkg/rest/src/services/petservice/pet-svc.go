package petservice

import (
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/handlers/petdao"
	"github.com/petmeds24/backend/pkg/rest/src/services/imageservice"
)

type PetService struct {
	imageKit *imageservice.ImgKit
	petDao   *petdao.PetDao
}

func NewPetService(globalCfg *config.GlobalConfig) *PetService {
	imgKit := imageservice.NewImgKit(globalCfg)
	return &PetService{
		imageKit: imgKit,
		petDao:   petdao.NewPetDao(ctx),
	}
}
