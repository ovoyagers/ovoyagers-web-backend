package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/controllers/petcontroller"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
)

type PetRoute struct {
	petController *petcontroller.PetController
}

func NewPetRoute(globalCfg *config.GlobalConfig)PetRoute {
	pc := petcontroller.NewPetController(globalCfg)
	return PetRoute{petController: pc}
}

func (pr PetRoute) SetupPetRoute(rg *gin.RouterGroup) {
	router := rg.Group("/pet")
	router.Use(middlewares.DeserializeUser())

	router.POST("/insert-pet", pr.petController.InsertNewPetWithImage)
	router.GET("/primary-pet", pr.petController.GetPrimaryPet)
	router.GET("/list-pets", pr.petController.ListPets)
	router.PUT("/update-pet/:id", pr.petController.UpdatePet)
	router.PUT("/update-pet-image/:id", pr.petController.UpdatePetImage)
	router.DELETE("/delete-pet/:petId", pr.petController.DeletePet)
	router.DELETE("/delete-pet-image/:petId", pr.petController.DeletePetImage)
}
