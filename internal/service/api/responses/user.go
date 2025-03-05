package responses

import (
	"github.com/recovery-flow/users-storage/internal/service/domain/models"
	"github.com/recovery-flow/users-storage/resources"
)

func User(user models.User) resources.User {
	update := resources.BaseUserStorage + resources.UserStorageEndpoints.Base.Private + user.ID.String()
	res := resources.User{
		Data: resources.UserData{
			Id:   user.ID.String(),
			Type: resources.UserType,
			Attributes: resources.UserAttributes{
				Username:   user.Username,
				Avatar:     user.Avatar,
				Verified:   user.Verified,
				TitleName:  user.TitleName,
				Speciality: user.Speciality,
				Position:   user.Position,
				Country:    user.Country,
				City:       user.City,

				CreatedAt: user.CreatedAt.Time().UTC(),
			},
			Links: resources.LinksSelf{
				Self:   resources.BaseUserStorage + resources.UserStorageEndpoints.Base.Public + user.ID.String(),
				Update: &update,
			},
		},
	}
	if user.DateOfBirth != nil {
		dob := user.DateOfBirth.Time().UTC()
		res.Data.Attributes.DateOfBirth = &dob
	}
	if user.UpdatedAt != nil {
		upAt := user.UpdatedAt.Time().UTC()
		res.Data.Attributes.UpdatedAt = &upAt
	}
	return res

}
