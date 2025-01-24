package reponses

import (
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"github.com/recovery-flow/users-storage/resources"
)

func NewUserResponse(user models.User) resources.User {
	return resources.User{
		Data: resources.UserData{
			Id:   user.ID.String(),
			Type: resources.UserType,
			Attributes: resources.UserDataAttributes{
				Username:  user.Username,
				Avatar:    "",
				Role:      user.Role,
				CreatedAt: user.CreatedAt,
			},
		},
	}
}
