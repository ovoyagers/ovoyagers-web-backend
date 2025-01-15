package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	recordcontroller "github.com/petmeds24/backend/pkg/rest/src/controllers/record-controller"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
)

type RecordRoute struct {
	recordController *recordcontroller.RecordController
}

func NewRecordRoute(globalCfg *config.GlobalConfig) RecordRoute {
	return RecordRoute{
		recordController: recordcontroller.NewRecordController(globalCfg),
	}
}

func (rr *RecordRoute) SetupRecordRoute(rg *gin.RouterGroup) {
	router := rg.Group("/record")
	router.Use(middlewares.DeserializeUser())

	router.POST("/insert-medical-records", rr.recordController.InsertMedicalRecords)
	router.GET("/get-medical-records/:petId", rr.recordController.GetMedicalRecordsByPetId)
	router.GET("/get-medical-record/:medicalRecordId", rr.recordController.GetMedicalRecordById)
	router.POST("/delete-medical-record/:recordId", rr.recordController.DeleteMedicalRecordById)
}
