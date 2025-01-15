package userdao

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/models/usermodel"
	log "github.com/sirupsen/logrus"
)

func (ud *UserDao) UpdateUser(user usermodel.UpdateUser, id string) (map[string]interface{}, error) {
	query := "MATCH (u:User {id: $id}) SET u.fullname = $fullname, u.age = $age, u.gender = $gender, u.countryCode = $countryCode, u.phone = $phone RETURN u;"
	payload := map[string]interface{}{
		"id":          id,
		"fullname":    user.Fullname,
		"age":         user.Age,
		"gender":      user.Gender,
		"countryCode": user.CountryCode,
		"phone":       user.Phone,
	}
	// execute query
	result, err := neo4j.ExecuteQuery(ud.ctx, ud.DB, query, payload, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// check if user exists
	if len(result.Records) == 0 {
		log.Warn("User not found")
		return nil, nil
	}

	// extract user properties
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// convert properties to map
	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}

	// return user properties
	return properties, nil
}
