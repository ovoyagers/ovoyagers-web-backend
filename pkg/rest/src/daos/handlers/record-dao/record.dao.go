package recorddao

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/daos/client"
)

type RecordDao struct {
	DB          neo4j.DriverWithContext
	ctx         context.Context
	neo4jClient *client.Neo4jClient
	dbClient    *client.DbClient
}

func NewRecordDao(ctx context.Context) *RecordDao {
	dbClient := client.NewDbClient()
	driver := dbClient.GetDriver()
	neo4jClient := client.NewNeo4jClient(driver, ctx)
	return &RecordDao{
		DB:          driver,
		ctx:         ctx,
		neo4jClient: neo4jClient,
		dbClient:    dbClient,
	}
}
