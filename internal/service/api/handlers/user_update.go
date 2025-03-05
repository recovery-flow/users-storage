package handlers

import (
	"net/http"

	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/service/api/requests"
	"github.com/recovery-flow/users-storage/internal/service/api/responses"
	"github.com/recovery-flow/users-storage/internal/service/domain/models"
)

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	userID, _, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Error("Failed to retrieve account data")
		httpkit.RenderErr(w, problems.Unauthorized(err.Error()))
		return
	}

	req, err := requests.NewUserUpdate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	fields := map[string]any{
		"username":      req.Data.Attributes.Username,
		"verified":      req.Data.Attributes.Verified,
		"title_name":    req.Data.Attributes.TitleName,
		"speciality":    req.Data.Attributes.Speciality,
		"position":      req.Data.Attributes.Position,
		"city":          req.Data.Attributes.City,
		"country":       req.Data.Attributes.Country,
		"date_of_birth": req.Data.Attributes.DateOfBirth,
	}

	stmt := make(map[string]any)
	for key, value := range fields {
		if value != nil {
			stmt[key] = value
		}
	}

	user, err := Domain(r).UpdateUser(r.Context(), models.RequestQuery{
		Filters: map[string]models.QueryFilter{"_id": {Type: "strict", Method: "eq", Value: userID}},
	}, stmt)
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to update username")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.User(*user))
}
