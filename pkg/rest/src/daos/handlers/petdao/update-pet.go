package petdao

import (
	"encoding/json"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/models/petmodel"
)

// UpdatePet takes a petmodel.Pet and the ID of the pet to update, and
// updates the pet in the database. The updated_at field is set to the
// current time.
func (p *PetDao) UpdatePet(pet petmodel.Pet, petID string) error {
	now := time.Now().UTC()

	// Create a map of parameters to pass to the Neo4j query. This
	// includes the ID of the pet to update, and the fields to update.
	// The updated_at field is set to the current time.
	params := map[string]interface{}{
		"id":         petID,
		"name":       pet.Name,
		"dob":        pet.Dob,
		"gender":     pet.Gender,
		"weight":     pet.Weight,
		"kind":       pet.Kind,
		"breed":      pet.Breed,
		"updated_at": now,
	}

	// Create a Cypher query to execute against the database. This query
	// finds the pet with the given ID, and sets the fields to the values
	// in the params map.
	query := `MATCH (p:Pet) WHERE p.id = $id SET p.name = $name, p.dob = $dob, p.gender = $gender, p.weight = $weight , p.kind = $kind, p.breed = $breed, p.updatedAt = $updated_at`

	// Execute the query against the database. If there is an error,
	// return it.
	_, err := neo4j.ExecuteQuery(p.ctx, p.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return err
	}

	// If the query was executed successfully, return nil.
	return nil
}

// UpdateProfilePicture updates the profile picture for the pet with the given ID.
// It takes the pet ID and a ProfilePicture struct, marshals the ProfilePicture
// to JSON, and updates the profilePicture field of the Pet node in the database.
// If there is an error during the update, it is returned.
func (p *PetDao) UpdateProfilePicture(petID string, profilePicture petmodel.ProfilePicture) error {
	// Marshal the ProfilePicture struct to JSON.
	picJson, err := json.Marshal(profilePicture)
	if err != nil {
		return err
	}
	params := map[string]interface{}{
		"id":             petID,
		"profilePicture": string(picJson),
	}
	// Create a Cypher query to execute against the database. This query
	// finds the pet with the given ID, and sets the profilePicture field to
	// the marshaled ProfilePicture struct.
	query := `
		MATCH (p:Pet {id: $id})
		SET p.profilePicture = $profilePicture
	`
	// Execute the query against the database. If there is an error,
	// return it.
	if _, err = neo4j.ExecuteQuery(p.ctx, p.DB, query, params, neo4j.EagerResultTransformer); err != nil {
		return err
	}
	// If the query was executed successfully, return nil.
	return nil
}
