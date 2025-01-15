package petdao

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/daos/client"
)

type PetDao struct {
	DB          neo4j.DriverWithContext
	ctx         context.Context
	neo4jClient *client.Neo4jClient
	dbClient    *client.DbClient
}

func NewPetDao(ctx context.Context) *PetDao {
	dbClient := client.NewDbClient()
	driver := dbClient.GetDriver()
	neo4jClient := client.NewNeo4jClient(driver, ctx)
	return &PetDao{
		DB:          driver,
		ctx:         ctx,
		neo4jClient: neo4jClient,
		dbClient:    dbClient,
	}
}

func (p *PetDao) TableName() string {
	return "pet"
}
