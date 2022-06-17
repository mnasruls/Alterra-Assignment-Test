package auth

import "alterra/entities"

type AuthServiceInterface interface {
	Login(AuthReq entities.AuthRequest) (entities.UserAuthResponse, error)
	Me(ID int, token interface{}) (entities.UserAuthResponse, error)
}
