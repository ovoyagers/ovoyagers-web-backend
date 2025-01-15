package invitedao

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/client"
)

type InviteDao struct {
	DB          neo4j.DriverWithContext
	ctx         context.Context
	neo4jClient *client.Neo4jClient
	dbClient    *client.DbClient
}

func NewInviteDao(globalCfg *config.GlobalConfig) *InviteDao {
	ctx := globalCfg.GetContext()
	dbClient := client.NewDbClient()
	driver := dbClient.GetDriver()
	neo4jClient := client.NewNeo4jClient(driver, ctx)
	return &InviteDao{
		DB:          driver,
		ctx:         ctx,
		neo4jClient: neo4jClient,
		dbClient:    dbClient,
	}
}
