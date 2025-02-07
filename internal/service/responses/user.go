package responses

import (
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
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
				Role:       string(user.Role),
				Type:       string(user.Type),
				Verified:   user.Verified,
				BanStatus:  string(user.BanStatus),
				TitleName:  user.TitleName,
				Speciality: user.Speciality,
				City:       user.City,
				Country:    user.Country,
				Level:      user.Level,
				Points:     user.Points,

				CreatedAt: user.CreatedAt.Time().UTC(),
			},
			Links: resources.LinksSelf{
				Self:   resources.BaseUserStorage + resources.UserStorageEndpoints.Base.Public + user.ID.String(),
				Update: &update,
			},
		},
	}
	if user.UpdatedAt != nil {
		upAt := user.UpdatedAt.Time().UTC()
		res.Data.Attributes.UpdatedAt = &upAt
	}
	return res

}
