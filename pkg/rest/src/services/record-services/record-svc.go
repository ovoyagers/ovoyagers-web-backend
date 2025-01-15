package recordservices

import (
	"github.com/petmeds24/backend/config"
	recorddao "github.com/petmeds24/backend/pkg/rest/src/daos/handlers/record-dao"
	"github.com/petmeds24/backend/pkg/rest/src/services/imageservice"
)

type RecordService struct {
	imageKit  *imageservice.ImgKit
	recordDao *recorddao.RecordDao
}

func NewRecordService(globalCfg *config.GlobalConfig) *RecordService {
	imgKit := imageservice.NewImgKit(globalCfg)
	return &RecordService{
		imageKit:  imgKit,
		recordDao: recorddao.NewRecordDao(ctx),
	}
}
