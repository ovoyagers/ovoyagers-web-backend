package authdao

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/client"
	"github.com/petmeds24/backend/pkg/rest/src/models/authmodel"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
	log "github.com/sirupsen/logrus"
)

type AuthDao struct {
	DB          neo4j.DriverWithContext
	ctx         context.Context
	neo4jClient *client.Neo4jClient
	dbClient    *client.DbClient
}

func NewAuthDao(globalCfg *config.GlobalConfig) *AuthDao {
	dbClient := client.NewDbClient()
	driver := dbClient.GetDriver()
	ctx := globalCfg.GetContext()
	neo4jClient := client.NewNeo4jClient(driver, ctx)
	return &AuthDao{DB: driver, ctx: ctx, neo4jClient: neo4jClient, dbClient: dbClient}
}

func (authDao *AuthDao) RegisterUser(register *authmodel.RegisterRequest, otp string) (map[string]interface{}, error) {
	user := &authmodel.User{}
	now := time.Now().UTC()
	hashedPassword, err := utils.HashPassword(register.Password)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	picJson, err := json.Marshal(user.ProfilePicture)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	params := map[string]interface{}{
		"id":             uuid.New().String(),
		"fullname":       register.Fullname,
		"phone":          register.Phone,
		"countryCode":    register.CountryCode,
		"email":          register.Email,
		"password":       hashedPassword,
		"dob":            user.DOB,
		"age":            user.Age,
		"gender":         user.Gender,
		"profilePicture": string(picJson),
		"isVerified":     user.IsVerified,
		"active":         user.Active,
		"coins":          user.Coins,
		"otp":            otp,
		"created_at":     now,
		"updated_at":     now,
	}
	query := "CREATE (u:User {id: $id, fullname: $fullname, phone: $phone, countryCode: $countryCode, email: $email, password: $password, dob: $dob, age: $age, gender: $gender, profilePicture: $profilePicture, isVerified: $isVerified, active: $active, coins: $coins, otp: $otp, created_at: $created_at, updated_at: $updated_at}) RETURN u"
	result, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)

	if err != nil {
		return nil, err
	}

	if len(result.Records) == 0 {
		return nil, err
	}

	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	properties := make(map[string]interface{})
	returnValues := []string{"email", "id", "isVerified"}
	for key, value := range node.Props {
		if slices.Contains(returnValues, key) {
			properties[key] = value
		}
	}
	return properties, nil
}

func (authDao *AuthDao) GetUserByPhone(phone string) (map[string]interface{}, int, error) {
	res, err := authDao.neo4jClient.GetNodeByFields(map[string]interface{}{"phone": phone}, "User")
	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}
	if res == nil {
		log.Error(utils.ErrNotFound)
		return nil, http.StatusNotFound, utils.ErrNotFound
	}
	log.Info(res)
	userMap, err := authDao.neo4jClient.ConvertNodeToMap(res)
	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	return userMap, http.StatusOK, nil
}

// GetUserByEmail retrieves a user from the database using their email
// It returns a map of properties of the user and an error if any
func (authDao *AuthDao) GetUserByEmail(email string) (map[string]interface{}, error) {
	// Get user from database using email as field
	query := "MATCH (u:User {email: $email}) RETURN u"
	params := map[string]interface{}{"email": email}
	result, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if len(result.Records) == 0 {
		log.Error(utils.ErrNotFound)
		return nil, utils.ErrNotFound
	}

	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}
	return properties, nil
}

func (authDao *AuthDao) UpdateOTP(email string, otp string) error {
	query := "MATCH (u:User {email: $email}) SET u.otp = $otp RETURN u"
	params := map[string]interface{}{"email": email, "otp": otp}
	_, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (authDao *AuthDao) GetOTP(email string) (map[string]interface{}, error) {
	query := "MATCH (u:User {email: $email}) RETURN u"
	params := map[string]interface{}{"email": email}
	result, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if len(result.Records) == 0 {
		log.Error(utils.ErrNotFound)
		return nil, utils.ErrNotFound
	}

	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}
	return properties, nil
}

func (authDao *AuthDao) VerifyOTP(email string) (map[string]interface{}, error) {
	// Update the user's isVerified field to true
	query := "MATCH (u:User {email: $email}) SET u.isVerified = true RETURN u"
	params := map[string]interface{}{"email": email}
	if _, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer); err != nil {
		log.Error(err)
		return nil, err
	}
	// Remove the otp field from the user
	query = "MATCH (u:User {email: $email}) SET u.otp = null RETURN u"
	result, err := neo4j.ExecuteQuery(authDao.ctx, authDao.DB, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if len(result.Records) == 0 {
		log.Error(utils.ErrNotFound)
		return nil, utils.ErrNotFound
	}

	node, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "u")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	properties := make(map[string]interface{})
	for key, value := range node.Props {
		properties[key] = value
	}

	return properties, nil
}

func (authdao *AuthDao) ForgetPassword(email string) (map[string]interface{}, error) {
	return nil, nil
}
