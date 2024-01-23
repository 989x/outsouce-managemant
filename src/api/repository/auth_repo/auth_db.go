package repository

import (
	"bytes"
	"net/http"
	"osm/api/models"

	"go.mongodb.org/mongo-driver/bson"
)

type authRepository struct {
	mgConn *models.MongoInstance
}

func NewAuthRepository(mgConn *models.MongoInstance) AuthRepository {
	return authRepository{mgConn: mgConn}
}
func (r authRepository) RepGetPwd(buffer *bytes.Buffer) (*http.Request, error) {
	request, err := http.NewRequest("POST", "https://one.th/api/oauth/getpwd", buffer)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (r authRepository) RepFindAdmin() (*models.Login_Find_result, error) {
	collection := r.mgConn.Db.Collection("users")

	var admin_result models.Login_Find_result

	query := bson.D{{Key: "account_id", Value: "0000000000"}}
	if err := collection.FindOne(r.mgConn.Ctx, query).Decode(&admin_result); err != nil {
		return nil, err
	}

	return &admin_result, nil
}
