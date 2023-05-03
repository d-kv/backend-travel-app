package controllerv0

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

func (c *Controller) AddAchievement(ctx context.Context, achievement *user.Achievement, userUUID string) error {
	const mName = "Controller.AddAchievement"

	u, err := c.userProvider.User(ctx, userUUID)
	if err != nil {
		log.Warn().
			Err(err).
			Str("method", mName).
			Msg("error from userProvider")

		return err
	}

	u.Achievements = append(u.Achievements, *achievement)

	err = c.userProvider.Update(ctx, userUUID, u)
	if err != nil {
		log.Error().
			Err(err).
			Str("method", mName).
			Msg("error from userProvider")

		return err
	}

	return nil
}
