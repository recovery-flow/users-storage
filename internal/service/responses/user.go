package responses

import (
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"github.com/recovery-flow/users-storage/resources"
)

func User(user models.User) resources.User {
	var projects []string
	for _, project := range user.Projects {
		projects = append(projects, project.String())
	}
	var organizations []string
	for _, organization := range user.Organizations {
		organizations = append(organizations, organization.String())
	}
	var ideas []string
	for _, idea := range user.Ideas {
		ideas = append(ideas, idea.String())
	}
	var reportsSent []string
	for _, report := range user.ReportsSent {
		reportsSent = append(reportsSent, report.String())
	}
	var reportsReceived []string
	for _, report := range user.ReportsReceived {
		reportsReceived = append(reportsReceived, report.String())
	}

	return resources.User{
		Data: resources.UserData{
			Id:   user.ID.String(),
			Type: resources.UserType,
			Attributes: resources.UserDataAttributes{
				Username:        user.Username,
				Avatar:          user.Avatar,
				Role:            string(user.Role),
				Projects:        projects,
				Ideas:           ideas,
				Organizations:   organizations,
				ReportsSent:     reportsSent,
				ReportsReceived: reportsReceived,

				UpdatedAt: &user.UpdatedAt,
				CreatedAt: user.CreatedAt,
			},
		},
	}
}
