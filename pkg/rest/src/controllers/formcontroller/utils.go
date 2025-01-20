package formcontroller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Constants for defaults and limits
const (
	defaultLimit = 10
	defaultPage  = 1
	maxLimit     = 100
)

// parsePaginationParams extracts and validates pagination parameters from the request context.
func parsePaginationParams(c *gin.Context) (limit, page, offset int, err error) {
	limit, page = defaultLimit, defaultPage

	// Parse limit
	if limitStr := c.Query("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 || limit > maxLimit {
			return 0, 0, 0, fmt.Errorf("invalid limit value: %w", err)
		}
	}

	// Parse page
	if pageStr := c.Query("page"); pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			return 0, 0, 0, fmt.Errorf("invalid page value: %w", err)
		}
	}

	offset = limit * (page - 1)
	return limit, page, offset, nil
}

// validateCategory ensures the category is valid.
func validateCategory(category string) error {
	validCategories := map[string]bool{
		"contact": true,
		"hotel":   true,
		"flights": true,
	}
	if !validCategories[category] {
		return fmt.Errorf("category must be one of contact, hotel, flights, but got: %s", category)
	}
	return nil
}

// prepareResponse prepares a consistent paginated response.
func prepareResponse(forms interface{}, totalCount, limit, page int) map[string]interface{} {
	totalPages := (totalCount + limit - 1) / limit
	offset := limit * (page - 1)

	return map[string]interface{}{
		"forms":         forms,
		"total_count":   totalCount,
		"current_page":  page,
		"limit":         limit,
		"total_pages":   totalPages,
		"has_next_page": offset+limit < totalCount,
		"has_prev_page": offset > 0,
	}
}
