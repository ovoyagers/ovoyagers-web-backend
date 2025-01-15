package petdao

import (
	"encoding/json"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/models/petmodel"
)

func (pd *PetDao) DeletePet(petID string) error {
	params := map[string]interface{}{
		"id": petID,
	}
	query := `
		MATCH (p:Pet {id: $id}) 
		OPTIONAL MATCH (p)-[:HAS_RECORD]->(r:Record)
		MATCH (p)-[:BELONGS_TO_USER]->(u:User)
		DETACH DELETE p, r
	`
	_, err := neo4j.ExecuteQuery(pd.ctx, pd.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return err
	}
	return nil
}

func (pd *PetDao) DeletePetProfilePicture(petID string) error {
	var profilePicture petmodel.ProfilePicture
	picJson, err := json.Marshal(profilePicture)
	if err != nil {
		return err
	}
	params := map[string]interface{}{
		"id":             petID,
		"profilePicture": string(picJson),
	}
	query := `
		MATCH (p:Pet {id: $id})
		SET p.profilePicture = $profilePicture
	`
	_, err = neo4j.ExecuteQuery(pd.ctx, pd.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return err
	}
	return nil
}
