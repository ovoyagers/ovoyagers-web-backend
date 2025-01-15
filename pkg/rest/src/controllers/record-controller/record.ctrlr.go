package recordcontroller

import (
	"github.com/petmeds24/backend/config"
	recordservices "github.com/petmeds24/backend/pkg/rest/src/services/record-services"
)

type RecordController struct {
	recordSvc recordservices.RecordService
}

func NewRecordController(globalCfg *config.GlobalConfig) *RecordController {
	return &RecordController{
		recordSvc: *recordservices.NewRecordService(globalCfg),
	}
}
