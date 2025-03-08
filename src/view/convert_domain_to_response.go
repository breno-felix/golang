package view

import (
	"github.com/breno-felix/golang/src/controller/model/response"
	"github.com/breno-felix/golang/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		Id:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
