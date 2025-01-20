package dashboarddao

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

func (d *DashboardDao) GetWebAnalytics() (map[string]interface{}, error) {
	// Define query parameters
	params := map[string]interface{}{
		"category1": "contact",
		"category2": "hotel",
		"category3": "flights",
	}

	// Corrected query
	query := `
		MATCH (f:Form)
		WITH count(f) AS totalCount
		OPTIONAL MATCH (c:Form {category: $category1})
		WITH totalCount, count(c) AS contactCount
		OPTIONAL MATCH (h:Form {category: $category2})
		WITH totalCount, contactCount, count(h) AS hotelCount
		OPTIONAL MATCH (fl:Form {category: $category3})
		RETURN totalCount, contactCount, hotelCount, count(fl) AS flightsCount;
	`

	// Execute query
	result, err := neo4j.ExecuteQuery(d.ctx, d.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	// Check if result has records
	if len(result.Records) == 0 {
		return nil, nil
	}

	// Parse results
	analytics := make(map[string]interface{})
	record := result.Records[0] // Assuming one record is returned
	if len(record.Values) >= 4 {
		if totalCount, ok := record.Values[0].(int64); ok {
			analytics["totalCount"] = totalCount
		}
		if contactCount, ok := record.Values[1].(int64); ok {
			analytics["contactCount"] = contactCount
		}
		if hotelCount, ok := record.Values[2].(int64); ok {
			analytics["hotelCount"] = hotelCount
		}
		if flightsCount, ok := record.Values[3].(int64); ok {
			analytics["flightsCount"] = flightsCount
		}
	}

	return analytics, nil
}
