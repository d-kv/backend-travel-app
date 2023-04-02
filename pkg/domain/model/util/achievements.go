package util

type Achievements struct {
	string //nolint:unused // Using dummy way to store achievements
}

func NewAchievements(ach string) *Achievements {
	return &Achievements{ach}
}
