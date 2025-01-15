package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jClient struct {
	Db  neo4j.DriverWithContext
	ctx context.Context
}

func NewNeo4jClient(db neo4j.DriverWithContext, ctx context.Context) *Neo4jClient {
	return &Neo4jClient{Db: db, ctx: ctx}
}

func (c *Neo4jClient) GetNodeByField(field string, value string, label string) (*neo4j.Node, error) {
	result, err := neo4j.ExecuteQuery(c.ctx, c.Db, "Match (n:"+label+" { "+field+": $"+field+" }) Return n", map[string]interface{}{field: value}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, fmt.Errorf("node not found")
	}
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "n")
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (c *Neo4jClient) GetNodeById(id string, label string) (*neo4j.Node, error) {
	result, err := neo4j.ExecuteQuery(c.ctx, c.Db, "Match (n:"+label+" { id: $id }) Return n Limit 1", map[string]interface{}{"id": id}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, fmt.Errorf("no element found with id: %s", id)
	}
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "n")
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (c *Neo4jClient) GetNodesByLabel(label string) ([]*neo4j.Node, error) {
	result, err := neo4j.ExecuteQuery(c.ctx, c.Db, "Match (n:"+label+" ) Return n", nil, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, fmt.Errorf("no element found")
	}
	nodes := make([]*neo4j.Node, 0)
	for _, record := range result.Records {
		node, _, err := neo4j.GetRecordValue[neo4j.Node](record, "n")
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, &node)
	}
	return nodes, nil
}

func (c *Neo4jClient) GetNodeByFields(params map[string]interface{}, label string) (*neo4j.Node, error) {
	var fieldParams string
	if len(params) == 0 {
		return nil, fmt.Errorf("no fields provided")
	}
	for field := range params {
		fieldParams += fmt.Sprintf("%s: $%s ,", field, field)
	}
	fieldParams = fieldParams[:len(fieldParams)-2]

	result, err := neo4j.ExecuteQuery(c.ctx, c.Db, "Match (n:"+label+" { "+fieldParams+" }) Return n", params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, fmt.Errorf("no element found")
	}
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "n")
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (c *Neo4jClient) CreateNode(params map[string]interface{}, label string) (*neo4j.Node, error) {
	var fieldParams string
	if len(params) == 0 {
		return nil, fmt.Errorf("no fields provided")
	}
	for field := range params {
		fieldParams += fmt.Sprintf("%s: $%s ,", field, field)
	}
	fieldParams = fieldParams[:len(fieldParams)-2]
	result, err := neo4j.ExecuteQuery(c.ctx, c.Db, "Create (n:"+label+" { "+fieldParams+" }) Return n", params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	if len(result.Records) == 0 {
		return nil, fmt.Errorf("unable to create node")
	}
	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "n")
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (c *Neo4jClient) CountNodesByFields(params map[string]interface{}, label string) (int64, error) {
	var fieldParams string
	if len(params) == 0 {
		return 0, fmt.Errorf("no fields provided")
	}

	for field := range params {
		fieldParams += fmt.Sprintf("%s: $%s ,", field, field)
	}
	fieldParams = fieldParams[:len(fieldParams)-2]

	result, err := neo4j.ExecuteQuery(c.ctx, c.Db, "Match (n:"+label+" { "+fieldParams+" }) Return count(n)", params, neo4j.EagerResultTransformer)
	if err != nil {
		return 0, err
	}
	count, _, err := neo4j.GetRecordValue[int64](result.Records[0], "count(n)")
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *Neo4jClient) CountByFieldsWithOrQuery(params map[string]interface{}, label string) (int64, error) {
	// Check if fields are provided
	if len(params) == 0 {
		return 0, fmt.Errorf("no fields provided")
	}

	// Construct the dynamic WHERE clause
	var fieldConditions []string
	for field := range params {
		// Add condition for each field to match against the provided parameters
		fieldConditions = append(fieldConditions, fmt.Sprintf("n.%s = $%s", field, field))
	}

	// Join all conditions with "OR" to create the WHERE clause
	whereClause := strings.Join(fieldConditions, " OR ")

	// Construct the complete Cypher query
	query := fmt.Sprintf("MATCH (n:%s) WHERE %s RETURN count(n) AS count", label, whereClause)

	// Execute the query
	result, err := neo4j.ExecuteQuery(c.ctx, c.Db, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return 0, err
	}

	// Check if any record is found
	if len(result.Records) == 0 {
		return 0, nil // No users found
	}

	// Retrieve and return the count from the result
	count, _, err := neo4j.GetRecordValue[int64](result.Records[0], "count")
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *Neo4jClient) ConvertNodeToMap(node *neo4j.Node) (map[string]interface{}, error) {
	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}
	return properties, nil
}
