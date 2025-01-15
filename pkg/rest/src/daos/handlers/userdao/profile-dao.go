package userdao

import (
	"encoding/json"

	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/models/authmodel"
	log "github.com/sirupsen/logrus"
)

func (ud *UserDao) GetProfileInfo(userid string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"id": userid,
	}
	query := "MATCH (u:User)-[:FOLLOWS]->(:User {id: $id}) WITH COUNT(u) as followers_count, COLLECT(u.username) as followers MATCH (:User {id: $id})-[:FOLLOWS]->(flwng:User) RETURN COUNT(flwng) as following_count,COLLECT(flwng.username), followers_count, followers;"

	// execute query
	result, err := neo4j.ExecuteQuery(ud.ctx, ud.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// check if user exists
	if len(result.Records) == 0 {
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

func (ud *UserDao) EditProfile(profile map[string]interface{}) (map[string]interface{}, error) {
	query := "MATCH (u:User {id: $id}) SET u += $profile RETURN u;"

	// execute query
	result, err := neo4j.ExecuteQuery(ud.ctx, ud.DB, query, profile, neo4j.EagerResultTransformer)
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

func (ud *UserDao) UpdateProfilePicture(imageInfo *uploader.UploadResponse, userid string) (map[string]interface{}, error) {
	query := "MATCH (u:User {id: $id}) SET u.profilePicture = $profilePicture RETURN u;"
	profilePicture := authmodel.ProfilePicture{
		FileId:       imageInfo.Data.FileId,
		Url:          imageInfo.Data.Url,
		Name:         imageInfo.Data.Name,
		ThumbnailUrl: imageInfo.Data.ThumbnailUrl,
	}

	ppJson, err := json.Marshal(profilePicture)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	profile := map[string]interface{}{
		"id":             userid,
		"profilePicture": string(ppJson),
	}

	// execute query
	result, err := neo4j.ExecuteQuery(ud.ctx, ud.DB, query, profile, neo4j.EagerResultTransformer)
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

func (ud *UserDao) DeleteProfilePicture(userid string) (map[string]interface{}, error) {
	query := "MATCH (u:User {id: $id}) SET u.profilePicture = $profilePicture RETURN u;"

	var profilePicture authmodel.ProfilePicture

	ppJson, err := json.Marshal(profilePicture)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	profile := map[string]interface{}{
		"id":             userid,
		"profilePicture": string(ppJson),
	}

	// execute query
	result, err := neo4j.ExecuteQuery(ud.ctx, ud.DB, query, profile, neo4j.EagerResultTransformer)
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

func (ud *UserDao) GetUserById(id string) (map[string]interface{}, error) {
	query := "MATCH (u:User {id: $id}) RETURN u;"
	profile := map[string]interface{}{
		"id": id,
	}

	// execute query
	result, err := neo4j.ExecuteQuery(ud.ctx, ud.DB, query, profile, neo4j.EagerResultTransformer)
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
