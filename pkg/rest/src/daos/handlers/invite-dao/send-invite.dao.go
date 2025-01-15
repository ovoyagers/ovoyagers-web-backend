package invitedao

import (
	"errors"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/pkg/rest/src/models/invitemodel"
)

func (id *InviteDao) SendInvite(invite invitemodel.InviteUser, userId string) error {
	now := time.Now().UTC().Unix()
	// add 30 days to the current time
	expiresAt := time.Now().UTC().Add(30 * 24 * time.Hour).Unix()
	params := map[string]interface{}{
		"id":        invite.Id,
		"userId":    userId,
		"email":     invite.Email,
		"status":    invite.Status,
		"expiresAt": expiresAt,
		"createdAt": now,
		"updatedAt": now,
	}
	// check if email exists else create invite
	exists, err := id.CheckInviteExists(invite.Email)
	if err != nil {
		return err
	}

	// return error if invite already exists
	if exists {
		return errors.New("invite already exists")
	} 

	query := `MATCH (u:User {id: $userId}) CREATE (i:Invite {id: $id, email: $email, status: $status, expiresAt: $expiresAt})-[:INVITE {createdAt: $createdAt, updatedAt: $updatedAt}]->(u)`
	if _, err := neo4j.ExecuteQuery(id.ctx, id.DB, query, params, neo4j.EagerResultTransformer); err != nil {
		return err
	}

	return nil
}

func (id *InviteDao) CheckInviteExists(email string) (bool, error) {
	params := map[string]interface{}{
		"email": email,
	}
	query := `MATCH (i:Invite {email: $email}) RETURN i`
	if _, err := neo4j.ExecuteQuery(id.ctx, id.DB, query, params, neo4j.EagerResultTransformer); err != nil {
		return false, err
	}

	return true, nil
}