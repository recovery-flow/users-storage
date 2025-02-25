package callbacks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/recovery-flow/users-storage/internal/service"
	"github.com/recovery-flow/users-storage/internal/service/domain/models"
	"github.com/recovery-flow/users-storage/internal/service/infra/events/rabbit/evebody"
)

func AccountUpdateRole(ctx context.Context, svc *service.Service, body []byte) error {
	if svc == nil || svc.Domain == nil {
		return fmt.Errorf("service or domain layer is nil")
	}

	var event evebody.RoleUpdated
	err := json.Unmarshal(body, &event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event body: %w", err)
	}

	userID, err := uuid.Parse(event.AccountID)
	if err != nil {
		return fmt.Errorf("failed to parse account ID: %w", err)
	}

	_, err = svc.Domain.UpdateUser(ctx, models.RequestQuery{
		Filters: map[string]models.QueryFilter{"_id": {Type: "strict", Method: "eq", Value: userID}},
	}, map[string]interface{}{"role": event.Role})
	if err != nil {
		return fmt.Errorf("failed to update role: %w", err)
	}

	return nil
}
