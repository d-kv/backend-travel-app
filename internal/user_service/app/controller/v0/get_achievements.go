package iuser_ctrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/user_service/domain/model"
	"github.com/rs/zerolog/log"
)

func (c *UserController) GetAchievements(ctx context.Context, userUUID string) ([]model.Achievement, error) {
	const mName = "UserController.GetAchievements"

	u, err := c.userProvider.User(ctx, userUUID)
	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from userProvider")

		return nil, err
	}

	return u.Achievements, nil
}
