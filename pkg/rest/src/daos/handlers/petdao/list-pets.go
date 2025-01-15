package petdao

import (
	"encoding/json"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/models/petmodel"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

func (p *PetDao) ListPets(userID string) ([]map[string]interface{}, error) {
	query := `
		MATCH (p:Pet)-[:BELONGS_TO_USER]->(u:User {id: $userId})
		RETURN p
	`
	params := map[string]interface{}{
		"userId": userID,
	}
	result, err := neo4j.ExecuteQuery(p.ctx, p.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, utils.ErrNoDataFound
	}

	pets := make([]map[string]interface{}, len(result.Records))
	for i, record := range result.Records {
		petNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
		if err != nil {
			return nil, err
		}
		pets[i] = petNode.Props
		for key, value := range petNode.Props {
			if key == "profilePicture" {
				petPicture := make(map[string]interface{})
				err = json.Unmarshal([]byte(value.(string)), &petPicture)
				if err != nil {
					return nil, err
				}
				pets[i][key] = petPicture
			}
		}
	}
	return pets, nil
}
func (p *PetDao) GetPrimaryPet(userID string) (map[string]interface{}, error) {
	query := `
		MATCH (p:Pet)-[:BELONGS_TO_USER]->(u:User {id: $userId})
		WHERE p.isPrimary = true
		RETURN p
	`
	params := map[string]interface{}{
		"userId": userID,
	}
	result, err := neo4j.ExecuteQuery(p.ctx, p.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, utils.ErrNoDataFound
	}

	primaryPeyNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "p")
	if err != nil {
		return nil, err
	}
	primaryPet := primaryPeyNode.Props
	for key, value := range primaryPet {
		if key == "profilePicture" {
			petPicture := make(map[string]interface{})
			err = json.Unmarshal([]byte(value.(string)), &petPicture)
			if err != nil {
				return nil, err
			}
			primaryPet[key] = petPicture
		}
	}
	return primaryPet, nil
}

func (p *PetDao) GetProfilePictureByPetId(petID string) (petmodel.ProfilePicture, error) {
	var profilePicture petmodel.ProfilePicture
	query := `
		MATCH (p:Pet {id: $petID})
		RETURN p.profilePicture LIMIT 1
	`
	params := map[string]interface{}{
		"petID": petID,
	}
	result, err := neo4j.ExecuteQuery(p.ctx, p.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return profilePicture, err
	}
	if len(result.Records) == 0 {
		return profilePicture, utils.ErrNoDataFound
	}
	record := result.Records[0]
	value, _, err := neo4j.GetRecordValue[string](record, "p.profilePicture")
	if err != nil {
		return profilePicture, err
	}
	// convert string to struct
	err = json.Unmarshal([]byte(value), &profilePicture)
	if err != nil {
		return profilePicture, err
	}
	return profilePicture, nil
}
