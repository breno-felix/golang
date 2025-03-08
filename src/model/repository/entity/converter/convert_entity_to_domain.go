package converter

import (
	"github.com/breno-felix/golang/src/model"
	"github.com/breno-felix/golang/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetId(entity.Id.Hex())
	return domain
}
