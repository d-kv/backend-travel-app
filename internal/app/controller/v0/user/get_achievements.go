package iuser_ctrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

func (c *UserController) GetAchievements(ctx context.Context, userUUID string) ([]user.Achievement, error) {
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
