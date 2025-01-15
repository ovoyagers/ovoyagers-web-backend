package followdao

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/daos/client"
	log "github.com/sirupsen/logrus"
)

type FollowDao struct {
	DB          neo4j.DriverWithContext
	ctx         context.Context
	neo4jClient *client.Neo4jClient
}

func NewFolloWDao(ctx context.Context) *FollowDao {
	dbClient := client.NewDbClient()
	driver := dbClient.GetDriver()
	neo4jClient := client.NewNeo4jClient(driver, ctx)
	return &FollowDao{DB: driver, ctx: ctx, neo4jClient: neo4jClient}
}

func (fd *FollowDao) CreateFollowRequest(followUsername string, userid string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"username": followUsername,
		"id":       userid,
	}
	query := "MATCH (u:User {id: $id}) MATCH (f:User {username: $username}) CREATE (u)-[r:FOLLOWS {status: 'PENDING', since: localdatetime()}]->(f) RETURN r"

	// execute query
	result, err := neo4j.ExecuteQuery(fd.ctx, fd.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// check if user exists
	if len(result.Records) == 0 {
		return nil, nil
	}

	// extract user properties
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "r")
	if err != nil {
		return nil, err
	}

	// convert properties to map
	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}
	return properties, nil
}

func (fd *FollowDao) CancelFollowRequest(followerUsername string, userid string) error {
	params := map[string]interface{}{
		"id":       userid,
		"username": followerUsername,
	}
	query := "MATCH (u:User {username: $username})-[r:FOLLOWS]->(f:User {id: $id}) DELETE r return r"
	res, err := neo4j.ExecuteQuery(fd.ctx, fd.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return err
	}

	// check if record deleted
	if len(res.Records) != 0 {
		return fmt.Errorf("unable to delete record! something went wrong")
	}
	return nil
}

func (fd *FollowDao) AcceptFollowRequest(followUsername string, userid string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"username": followUsername,
		"id":       userid,
	}
	query := "MATCH (u:User {username: $username})-[r:FOLLOWS]->(f:User {id: $id}) SET r.status = 'ACCEPTED' r.since = localdatetime() RETURN r"

	// execute query
	result, err := neo4j.ExecuteQuery(fd.ctx, fd.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// check if user exists
	if len(result.Records) == 0 {
		return nil, nil
	}

	// extract user properties
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "r")
	if err != nil {
		return nil, err
	}
	// convert properties to map
	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}
	return properties, nil
}
