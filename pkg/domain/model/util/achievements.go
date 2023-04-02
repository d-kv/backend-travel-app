package util

type Achievements struct {
	string
}

func NewAchievements(ach string) *Achievements {
	return &Achievements{ach}
}
