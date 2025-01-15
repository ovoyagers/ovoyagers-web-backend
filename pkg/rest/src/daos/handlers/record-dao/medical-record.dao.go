package recorddao

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	log "github.com/sirupsen/logrus"

	recordmodel "github.com/petmeds24/backend/pkg/rest/src/models/record-model"
)

// InsertMedicalRecordDao inserts a new medical record for a pet.
//
// The method uses the id of the pet to find the corresponding node in the graph
// and creates a new node of type Record, with the given properties, and a
// relationship of type HAS_RECORD between the pet and the record.
//
// The properties of the Record node are the ones provided in the input
// medicalRecord parameter, plus the created_at and updated_at timestamps,
// which are set to the current time.
//
// If the query is successful, the method returns a map with the properties of
// the newly created Record node. If an error occurs, it returns that error.
func (rd *RecordDao) InsertMedicalRecordDao(medicalRecord recordmodel.RecordData) (map[string]interface{}, error) {
	now := time.Now().UTC()
	recordsJson, err := json.Marshal(medicalRecord.PetRecords)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	params := map[string]interface{}{
		"id":          uuid.New().String(),
		"user_id":     medicalRecord.UserId,
		"pet_id":      medicalRecord.PetId,
		"pet_records": string(recordsJson),
		"description": medicalRecord.Description,
		"created_at":  now,
		"updated_at":  now,
	}

	query := "MATCH (p:Pet {id: $pet_id}) CREATE (p)-[:HAS_RECORD]->(r:Record {id: $id, pet_records: $pet_records, description: $description, created_at: $created_at, updated_at: $updated_at}) RETURN r"
	result, err := neo4j.ExecuteQuery(rd.ctx, rd.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, nil
	}
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "r")
	if err != nil {
		return nil, err
	}
	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}
	return properties, nil
}

// GetMedicalRecordByPetId retrieves the medical record of a pet given its id.
// Returns nil if the pet or the record is not found.
// The pet records are returned as a slice of PetRecordImage structs.
func (rd *RecordDao) GetMedicalRecordsByPetId(petId string) ([]map[string]interface{}, error) {
	params := map[string]interface{}{
		"pet_id": petId,
	}
	query := "MATCH (p:Pet {id: $pet_id})-[:HAS_RECORD]->(r:Record) RETURN r"

	result, err := neo4j.ExecuteQuery(rd.ctx, rd.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, nil
	}
	records := make([]map[string]interface{}, len(result.Records))
	for i, record := range result.Records {
		recordNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "r")
		if err != nil {
			return nil, err
		}
		records[i] = recordNode.Props
		for key, value := range recordNode.Props {
			if key == "pet_records" {
				petRecords := make([]map[string]interface{}, 0)
				err = json.Unmarshal([]byte(value.(string)), &petRecords)
				if err != nil {
					return nil, err
				}
				records[i]["pet_records"] = petRecords
			}
		}
	}
	return records, nil
}

// DeleteMedicalRecordById deletes a medical record given its id.
// Returns an error if the query execution fails.
func (rd *RecordDao) DeleteMedicalRecordById(id string) error {
	params := map[string]interface{}{
		"id": id,
	}
	query := "MATCH (r:Record {id: $id}) DETACH DELETE r"
	_, err := neo4j.ExecuteQuery(rd.ctx, rd.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return err
	}
	return nil
}

func (rd *RecordDao) GetMedicalRecordByRecordId(recordId string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"record_id": recordId,
	}
	query := "MATCH (r:Record {id: $record_id}) RETURN r"
	result, err := neo4j.ExecuteQuery(rd.ctx, rd.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, nil
	}
	recordNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "r")
	if err != nil {
		return nil, err
	}
	properties := make(map[string]interface{})
	for key, value := range recordNode.Props {
		if key == "pet_records" {
			petRecords := make([]map[string]interface{}, 0)
			err = json.Unmarshal([]byte(value.(string)), &petRecords)
			if err != nil {
				return nil, err
			}
			properties["pet_records"] = petRecords
		} else {
			properties[key] = value
		}
	}
	return properties, nil
}
