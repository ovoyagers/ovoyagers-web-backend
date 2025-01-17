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
