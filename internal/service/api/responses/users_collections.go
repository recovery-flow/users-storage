package responses

import (
	"fmt"
	"net/url"

	"github.com/recovery-flow/users-storage/internal/service/domain/models"
	"github.com/recovery-flow/users-storage/resources"
)

func NewUsersCollectionResponse(users []models.User, baseURL string, queryParams url.Values, totalItems, pageSize, pageNumber int64) resources.UsersCollection {

	var data []resources.UserData
	for _, user := range users {
		data = append(data, User(user).Data)
	}

	links := resources.LinksPagination{
		Self:     *generatePaginationLink(baseURL, queryParams, pageNumber, pageSize),
		Previous: generatePaginationLink(baseURL, queryParams, pageNumber-1, pageSize),
		Next:     generatePaginationLink(baseURL, queryParams, pageNumber+1, pageSize),
	}

	if pageNumber <= 1 {
		links.Previous = nil
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	if pageNumber >= totalPages {
		links.Next = nil
	}

	return resources.UsersCollection{
		Data:  data,
		Links: links,
	}
}

func generatePaginationLink(baseURL string, queryParams url.Values, pageNumber, pageSize int64) *string {
	if pageNumber < 1 {
		return nil
	}

	queryParams.Set("page[number]", fmt.Sprintf("%d", pageNumber))
	queryParams.Set("page[size]", fmt.Sprintf("%d", pageSize))
	res := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
	return &res
}
