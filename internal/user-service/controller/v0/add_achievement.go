package ictrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/pkg/user-service/model"
)

func (c *UserController) AddAchievement(ctx context.Context, achievement *model.Achievement, userUUID string) error {
	const mName = "UserController.AddAchievement"

	u, err := c.userProvider.User(ctx, userUUID)
	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from userProvider")

		return err
	}

	u.Achievements = append(u.Achievements, *achievement)

	err = c.userProvider.Update(ctx, userUUID, u)
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from userProvider")

		return err
	}

	return nil
}
