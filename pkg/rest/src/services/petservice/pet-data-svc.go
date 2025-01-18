package petservice

import "github.com/petmeds24/backend/pkg/rest/src/models/petmodel"

func (ps *PetService) AddNewPet(pet petmodel.Pet, userID string) (map[string]interface{}, error) {
	return ps.petDao.AddNewPet(pet, userID)
}

func (ps *PetService) CheckPetExists(pet petmodel.Pet, userID string) (bool, error) {
	return ps.petDao.CheckPetExists(pet, userID)
}

func (ps *PetService) ListPets(userID string) ([]map[string]interface{}, error) {
	return ps.petDao.ListPets(userID)
}

func (ps *PetService) GetPrimaryPet(userID string) (map[string]interface{}, error) {
	return ps.petDao.GetPrimaryPet(userID)
}

func (ps *PetService) UpdatePet(pet petmodel.Pet, petID string) error {
	return ps.petDao.UpdatePet(pet, petID)
}

func (ps *PetService) DeletePet(petID string) error {
	return ps.petDao.DeletePet(petID)
}
