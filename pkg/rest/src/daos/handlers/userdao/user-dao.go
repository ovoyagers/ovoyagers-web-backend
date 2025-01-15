package userdao

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/client"
	"github.com/petmeds24/backend/pkg/rest/src/models/usermodel"
	log "github.com/sirupsen/logrus"
)

type UserDao struct {
	DB          neo4j.DriverWithContext
	ctx         context.Context
	neo4jClient *client.Neo4jClient
}

func NewAuthDao(globalCfg *config.GlobalConfig) *UserDao {
	ctx := globalCfg.GetContext()
	dbClient := client.NewDbClient()
	driver := dbClient.GetDriver()
	neo4jClient := client.NewNeo4jClient(driver, ctx)
	return &UserDao{DB: driver, ctx: ctx, neo4jClient: neo4jClient}
}

func (ad *UserDao) UpdateAboutUser(user *usermodel.AboutUser, userid string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"id":     userid,
		"name":   user.Name,
		"email":  user.Email,
		"dob":    user.DOB,
		"gender": user.Gender,
	}
	query := "MATCH (u:User {id: $id}) SET u.name = $name, u.email = $email, u.dob = $dob, u.gender = $gender RETURN u"

	// execute query
	result, err := neo4j.ExecuteQuery(ad.ctx, ad.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// check if user exists
	if len(result.Records) == 0 {
		log.Warn("data not found")
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

func (ad *UserDao) UpdateLanguages(lang *usermodel.Languages, userid string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"id":                 userid,
		"preferredLanguages": lang.PreferredLanguages,
		"nativeLanguages":    lang.NativeLanguages,
	}
	query := "MATCH (u:User {id: $id}) SET u.preferredLanguages = $preferredLanguages, u.nativeLanguages = $nativeLanguages RETURN u"

	// execute query
	result, err := neo4j.ExecuteQuery(ad.ctx, ad.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// check if user exists
	if len(result.Records) == 0 {
		log.Warn("data not found")
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

func (ad *UserDao) UpdateUsername(username string, userid string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"id":       userid,
		"username": username,
	}
	query := "MATCH (u:User {id: $id}) SET u.username = $username RETURN u"

	// execute query
	result, err := neo4j.ExecuteQuery(ad.ctx, ad.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// check if user exists
	if len(result.Records) == 0 {
		log.Warn("data not found")
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
