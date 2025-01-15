package petdao

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/models/petmodel"
	log "github.com/sirupsen/logrus"
)

func (p *PetDao) AddNewPet(pet petmodel.Pet, userID string) (map[string]interface{}, error) {
	now := time.Now().UTC()

	picJson, err := json.Marshal(pet.ProfilePicture)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	params := map[string]interface{}{
		"id":             pet.ID,
		"name":           pet.Name,
		"dob":            pet.Dob,
		"gender":         pet.Gender,
		"kind":           pet.Kind,
		"breed":          pet.Breed,
		"isPrimary":      pet.IsPrimary,
		"profilePicture": string(picJson),
		"createdAt":      now,
		"updatedAt":      now,
		"userId":         userID,
	}

	// Ensure the user exists before creating the pet
	userCheckQuery := "MATCH (u:User {id: $userId}) RETURN u"
	userCheckResult, err := neo4j.ExecuteQuery(p.ctx, p.DB, userCheckQuery, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(userCheckResult.Records) == 0 {
		return nil, errors.New("user not found")
	}

	// Check if the user already has any pets
	petCountQuery := "MATCH (p:Pet)-[:BELONGS_TO_USER]->(u:User {id: $userId}) RETURN COUNT(p) as petCount"
	petCountResult, err := neo4j.ExecuteQuery(p.ctx, p.DB, petCountQuery, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(petCountResult.Records) == 0 {
		return nil, errors.New("failed to check user's pet count")
	}

	// If the user has no pets, make the first pet primary
	petCount := petCountResult.Records[0].Values[0].(int64)
	if petCount == 0 {
		params["isPrimary"] = true
	}

	// If the new pet is marked as primary, update the existing primary pet to false
	if pet.IsPrimary {
		updatePrimaryQuery := `
			MATCH (p:Pet)-[:BELONGS_TO_USER]->(u:User {id: $userId})
			WHERE p.isPrimary = true
			SET p.isPrimary = false
		`
		_, err := neo4j.ExecuteQuery(p.ctx, p.DB, updatePrimaryQuery, params, neo4j.EagerResultTransformer)
		if err != nil {
			return nil, err
		}
	}

	// Create a query to create a new pet and add a relation to the user with the userID
	createQuery := `
		MATCH (u:User {id: $userId})
		CREATE (p:Pet {
			id: $id, name: $name, dob: $dob, gender: $gender, kind: $kind, breed: $breed,
			profilePicture: $profilePicture, isPrimary: $isPrimary, createdAt: $createdAt, 
			updatedAt: $updatedAt
		})-[:BELONGS_TO_USER {since: $createdAt}]->(u)
		RETURN p
	`

	// Execute the creation query
	createResult, err := neo4j.ExecuteQuery(p.ctx, p.DB, createQuery, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(createResult.Records) == 0 {
		return nil, errors.New("failed to create pet")
	}

	// Extract user node properties
	node, _, err := neo4j.GetRecordValue[neo4j.Node](createResult.Records[0], "p")
	if err != nil {
		return nil, err
	}

	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}

	return properties, nil
}

func (p *PetDao) CheckPetExists(pet petmodel.Pet, userID string) (bool, error) {
	params := map[string]interface{}{
		"name":   pet.Name,
		"dob":    pet.Dob,
		"kind":   pet.Kind,
		"breed":  pet.Breed,
		"userId": userID,
	}
	query := `
		MATCH (p:Pet {name: $name, dob: $dob, kind: $kind, breed: $breed}) -[:BELONGS_TO_USER]-> (u:User {id: $userId})
		RETURN p LIMIT 1
	`
	checkResult, err := neo4j.ExecuteQuery(p.ctx, p.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return false, err
	}
	if len(checkResult.Records) > 0 {
		return true, nil
	}
	return false, nil
}
