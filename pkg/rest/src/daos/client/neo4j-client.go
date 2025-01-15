package client

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/config"
)

type DbClient struct {
	cfg    *config.Config
	driver neo4j.DriverWithContext
}

// NewDbClient creates a new instance of DbClient. It first loads the configuration from the file specified by the
// ENVIRONMENT variable, then creates a new neo4j driver using the NEO4J_URI, NEO4J_USER, and NEO4J_PASSWORD
// configuration values. It returns an instance of DbClient with the driver and config set. If there is an error
// loading the configuration or creating the driver, it panics.
func NewDbClient() *DbClient {
	var Neo4jUser, Neo4jPassword, Neo4jUri string
	var cfg *config.Config
	var err error

	cfg, err = config.LoadConfig()
	if err != nil {
		panic(err)
	}

	Neo4jUser = cfg.NEO4J_USER
	Neo4jPassword = cfg.NEO4J_PASSWORD
	Neo4jUri = cfg.NEO4J_URI

	driver, err := neo4j.NewDriverWithContext(Neo4jUri, neo4j.BasicAuth(Neo4jUser, Neo4jPassword, ""))
	if err != nil {
		return nil
	}
	return &DbClient{
		cfg:    cfg,
		driver: driver,
	}
}

// GetDriver returns the underlying neo4j driver that is used by the DbClient to
// communicate with the database. This is useful for creating a Neo4jClient
// instance, which provides a higher-level interface for interacting with the
// graph database.
func (db *DbClient) GetDriver() neo4j.DriverWithContext {
	return db.driver
}

// GetClient takes a query and a map of parameters and runs it against the neo4j
// database. It returns a map of properties of the first node returned by the query,
// or an error if one occurred. If the query didn't return any nodes, it returns nil.
func (c *DbClient) ExecuteQuery(ctx context.Context, query string, params map[string]interface{}) (map[string]interface{}, error) {
	result, err := neo4j.ExecuteQuery(ctx, c.driver, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	// If the query didn't return any nodes, return nil.
	if len(result.Records) == 0 {
		return nil, nil
	}
	// Get the first node from the query
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		return nil, err
	}
	// Convert the node's properties to a map
	properties := make(map[string]interface{})

	for key, value := range node.Props {
		properties[key] = value
	}

	return properties, nil
}
