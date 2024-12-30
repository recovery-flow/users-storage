package requests

import (
	"encoding/json"
	"net/http"

	"github.com/cifra-city/users-storage/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func NewUpdateUserFull(r *http.Request) (req resources.UserUpdate, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	errs := validation.Errors{
		"data/type":       validation.Validate(req.Data.Type, validation.Required, validation.In("user_change")),
		"data/attributes": validation.Validate(req.Data.Attributes, validation.Required),
	}
	return req, errs.Filter()
}
