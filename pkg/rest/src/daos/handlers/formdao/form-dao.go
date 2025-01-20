package formdao

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/client"
	"github.com/petmeds24/backend/pkg/rest/src/models/formmodel"
)

type FormDao struct {
	DB          neo4j.DriverWithContext
	ctx         context.Context
	neo4jClient *client.Neo4jClient
	dbClient    *client.DbClient
}

func NewFormDao(globalCfg *config.GlobalConfig) *FormDao {
	dbClient := client.NewDbClient()
	driver := dbClient.GetDriver()
	ctx := globalCfg.GetContext()
	neo4jClient := client.NewNeo4jClient(driver, ctx)
	return &FormDao{DB: driver, ctx: ctx, neo4jClient: neo4jClient, dbClient: dbClient}
}

func (formDao *FormDao) CreateForm(form *formmodel.Form) (string, error) {
	now := time.Now().UTC()
	params := map[string]interface{}{
		"id":        uuid.New().String(),
		"fullname":  form.Fullname,
		"email":     form.Email,
		"mobile":    form.Mobile,
		"message":   form.Message,
		"createdAt": now,
		"updatedAt": now,
		"category":  form.Category,
	}

	// Combined query to check and create or identify existing form
	query := `
        MERGE (f:Form {email: $email, mobile: $mobile, category: $category, message: $message})
        ON CREATE SET 
            f.id = $id, 
            f.fullname = $fullname, 
            f.createdAt = $createdAt, 
            f.updatedAt = $updatedAt
        RETURN CASE 
            WHEN f.createdAt = $createdAt THEN 'created'
            ELSE 'exists'
        END AS status
    `

	// Execute the query
	records, err := neo4j.ExecuteQuery(formDao.ctx, formDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return "", fmt.Errorf("error executing query: %w", err)
	}

	// Check if records are returned
	if len(records.Records) > 0 {
		record := records.Records[0]
		status, _, err := neo4j.GetRecordValue[string](record, "status")
		if err != nil {
			return "", fmt.Errorf("error getting record value: %w", err)
		}
		if status == "exists" {
			return "", fmt.Errorf("form already exists")
		}
		return "", nil
	}

	return "", fmt.Errorf("no records found in query result")
}

func (formDao *FormDao) GetForms(limit, page, offset int) ([]map[string]interface{}, int, error) {
	params := map[string]interface{}{
		"limit":  limit,
		"page":   page,
		"offset": offset,
	}
	query := `
        MATCH (f:Form)
		WITH f
		ORDER BY f.createdAt DESC
		SKIP $offset
		LIMIT $limit
		MATCH (a:Form)
		RETURN f, count(a) AS totalCount
    `
	result, err := neo4j.ExecuteQuery(formDao.ctx, formDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, 0, fmt.Errorf("error executing query: %w", err)
	}
	count := len(result.Records)
	if count == 0 {
		return nil, count, fmt.Errorf("no forms found")
	}
	forms := make([]map[string]interface{}, len(result.Records))
	for i, record := range result.Records {
		formNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "f")
		if err != nil {
			return nil, 0, fmt.Errorf("error getting record value: %w", err)
		}
		forms[i] = formNode.Props
		countNode, _, err := neo4j.GetRecordValue[int64](record, "totalCount")
		if err != nil {
			return nil, 0, fmt.Errorf("error getting record value: %w", err)
		}
		count = int(countNode)
	}
	return forms, count, nil
}

func (formDao *FormDao) GetFormsByCategory(category string, limit, page, offset int) ([]map[string]interface{}, int, error) {
	params := map[string]interface{}{
		"category": category,
		"limit":    limit,
		"page":     page,
		"offset":   offset,
	}
	// Query to get forms by category
	query := `
        MATCH (f:Form {category: $category})
		WITH f
		ORDER BY f.createdAt DESC
		SKIP $offset
		LIMIT $limit
		MATCH (a:Form {category: $category})
		RETURN f, count(a) AS totalCount
    `
	result, err := neo4j.ExecuteQuery(formDao.ctx, formDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, 0, fmt.Errorf("error executing query: %w", err)
	}
	count := len(result.Records)
	if count == 0 {
		return nil, count, fmt.Errorf("no forms found")
	}
	forms := make([]map[string]interface{}, len(result.Records))
	for i, record := range result.Records {
		formNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "f")
		if err != nil {
			return nil, 0, fmt.Errorf("error getting record value: %w", err)
		}
		forms[i] = formNode.Props
		countNode, _, err := neo4j.GetRecordValue[int64](record, "totalCount")
		if err != nil {
			return nil, 0, fmt.Errorf("error getting record value: %w", err)
		}
		count = int(countNode)
	}
	return forms, count, nil
}
