package dashboarddao

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/client"
)

type DashboardDao struct {
	DB          neo4j.DriverWithContext
	ctx         context.Context
	neo4jClient *client.Neo4jClient
	dbClient    *client.DbClient
}

func NewDashboardDao(globalCfg *config.GlobalConfig) *DashboardDao {
	dbClient := client.NewDbClient()
	driver := dbClient.GetDriver()
	ctx := globalCfg.GetContext()
	neo4jClient := client.NewNeo4jClient(driver, ctx)
	return &DashboardDao{DB: driver, ctx: ctx, neo4jClient: neo4jClient, dbClient: dbClient}
}
