package service

import "osm/api/models"

type AuthService interface {
	SrvLogin(*models.Login_body) (*models.Login_response, error)
	SrvAdminLogin(*models.Login_body) (*models.Admin_login_response, error)
}
