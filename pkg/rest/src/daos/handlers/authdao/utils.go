package authdao

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/models/authmodel"
	log "github.com/sirupsen/logrus"
)

func (authDao *AuthDao) IsUserExists(user authmodel.CheckUser) (bool, error) {
	var query string
	// check if user exists
	if user.Id != "" {
		query = "MATCH (u:User {id: $id}) RETURN u"
	}
	// check if email exists
	if user.Email != "" {
		query = "MATCH (u:User {email: $email}) RETURN u"
	}

	params := map[string]interface{}{"email": user.Email, "id": user.Id}
	result, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return false, err
	}
	if len(result.Records) == 0 {
		return false, nil
	}
	return true, nil
}
