package repository

import (
	"bytes"
	"net/http"
	"osm/api/models"
)

type AuthRepository interface {
	RepGetPwd(*bytes.Buffer) (*http.Request, error)
	RepFindAdmin() (*models.Login_Find_result, error)
}
