package responses

import (
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"github.com/recovery-flow/users-storage/resources"
)

func User(user models.User) resources.User {
	upAt := user.UpdatedAt.Time().UTC()
	return resources.User{
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

				UpdatedAt: &upAt,
				CreatedAt: user.CreatedAt.Time().UTC(),
			},
		},
	}
}
