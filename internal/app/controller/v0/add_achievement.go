package controllerv0

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

func (c *Controller) AddAchievement(ctx context.Context, achievement *user.Achievement, userUUID string) error {
	u, err := c.userProvider.User(ctx, userUUID)
	if err != nil {
		log.Warn().
			Err(err)
		return err
	}

	u.Achievements = append(u.Achievements, *achievement)

	err = c.userProvider.Update(ctx, userUUID, u)
	if err != nil {
		log.Error().
			Err(err)
		return err
	}

	return nil
}
