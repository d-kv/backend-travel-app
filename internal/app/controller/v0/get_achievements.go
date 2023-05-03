package controllerv0

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

func (c *Controller) GetAchievements(ctx context.Context, userUUID string) ([]user.Achievement, error) {
	const mName = "Controller.GetAchievements"

	u, err := c.userProvider.User(ctx, userUUID)
	if err != nil {
		log.Warn().
			Err(err).
			Str("method", mName).
			Msg("error from userProvider")

		return nil, err
	}

	return u.Achievements, nil
}
