package authdao

import (
	"time"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/models/authmodel"
)

func (authDao *AuthDao) FindUserByPhone(phone string) (map[string]interface{}, error) {
	// neo4j query to find user by phone
	params := map[string]interface{}{"phone": phone}
	query := "MATCH (u:User { phone: $phone }) RETURN u"

	// execute query
	result, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	// check if user exists
	if len(result.Records) == 0 {
		return nil, nil
	}

	// extract pet properties
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		return nil, err
	}

	// convert properties to map
	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}

	// return pet properties
	return properties, nil
}

func (authDao *AuthDao) FindUserByID(id string) (map[string]interface{}, error) {
	// neo4j query to find user by id
	params := map[string]interface{}{"id": id}
	query := "MATCH (u:User { id: $id }) RETURN u"

	// execute query
	result, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	// check if user exists
	if len(result.Records) == 0 {
		return nil, nil
	}

	// extract user properties
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		return nil, err
	}

	// convert properties to map
	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}

	// return pet properties
	return properties, nil
}

func (authDao *AuthDao) InsertUser(pet authmodel.User) (map[string]interface{}, error) {
	now := time.Now().UTC()
	params := map[string]interface{}{
		"id":             uuid.New().String(),
		"fullname":       pet.Fullname,
		"countryCode":    pet.CountryCode,
		"phone":          pet.Phone,
		"email":          pet.Email,
		"dob":            pet.DOB,
		"age":            pet.Age,
		"gender":         pet.Gender,
		"profilePicture": pet.ProfilePicture,
		"isVerified":     pet.IsVerified,
		"active":         pet.Active,
		"coins":          pet.Coins,
		"createdAt":      now,
		"updatedAt":      now,
	}
	query := "CREATE (u:User { id: $id, fullname: $fullname, countryCode: $countryCode, phone: $phone, email: $email, dob: $dob, age: $age, gender: $gender, profilePicture: $profilePicture, isVerified: $isVerified, active: $active, coins: $coins, createdAt: $createdAt, updatedAt: $updatedAt }) RETURN u"

	result, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, nil
	}
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		return nil, err
	}

	properties := make(map[string]interface{})

	for key, value := range node.Props {
		properties[key] = value
	}
	return properties, nil
}

func (authDao *AuthDao) UpdateVerifiedUser(userid string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"id":         userid,
		"isVerified": true,
	}
	query := "MATCH (u:User { id: $id }) SET u.isVerified = $isVerified RETURN u"
	result, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, nil
	}
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		return nil, err
	}
	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}
	return properties, nil
}
