//nolint:revive, stylecheck, gochecknoglobals // Using SNAKE_CASE & global maps for enums
package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type AchievementE int32

func (a AchievementE) string() string {
	return achievementEName[a]
}

// Don't use iota because it is easier to keep enum in sync with api using
// explicit numbers.
const (
	ACH_UNSPECIFIED        AchievementE = 0
	ACH_GOOD_REST          AchievementE = 1
	ACH_INSPECTOR_MICHELIN AchievementE = 2
	ACH_GUIDE              AchievementE = 3
	ACH_TRUE_GREEN         AchievementE = 4
	ACH_BAD_USER           AchievementE = 5
	ACH_HR                 AchievementE = 6
	ACH_PREMIUM            AchievementE = 7
	ACH_DIFFICULT_CHOICE   AchievementE = 8
	ACH_TESTER             AchievementE = 9
	ACH_DEUS_VULT          AchievementE = 10
)

// Enum value maps for achievement.
var (
	achievementEName = map[AchievementE]string{
		ACH_UNSPECIFIED:        "ACH_UNSPECIFIED",
		ACH_GOOD_REST:          "ACH_GOOD_REST",
		ACH_INSPECTOR_MICHELIN: "ACH_INSPECTOR_MICHELIN",
		ACH_GUIDE:              "ACH_GUIDE",
		ACH_TRUE_GREEN:         "ACH_TRUE_GREEN",
		ACH_BAD_USER:           "ACH_BAD_USER",
		ACH_HR:                 "ACH_HR",
		ACH_PREMIUM:            "ACH_PREMIUM",
		ACH_DIFFICULT_CHOICE:   "ACH_DIFFICULT_CHOICE",
		ACH_TESTER:             "ACH_TESTER",
		ACH_DEUS_VULT:          "ACH_DEUS_VULT",
	}
	achievementEValue = map[string]AchievementE{
		"ACH_UNSPECIFIED":        ACH_UNSPECIFIED,
		"ACH_GOOD_REST":          ACH_GOOD_REST,
		"ACH_INSPECTOR_MICHELIN": ACH_INSPECTOR_MICHELIN,
		"ACH_GUIDE":              ACH_GUIDE,
		"ACH_TRUE_GREEN":         ACH_TRUE_GREEN,
		"ACH_BAD_USER":           ACH_BAD_USER,
		"ACH_HR":                 ACH_HR,
		"ACH_PREMIUM":            ACH_PREMIUM,
		"ACH_DIFFICULT_CHOICE":   ACH_DIFFICULT_CHOICE,
		"ACH_TESTER":             ACH_TESTER,
		"ACH_DEUS_VULT":          ACH_DEUS_VULT,
	}
)

func (a Achievement) MarshalBSON() ([]byte, error) {
	marshalStruct := bsonAchievementStruct{
		Achievement: a.Achievement.string(),
		ReachedAt:   a.ReachedAt,
	}

	return bson.Marshal(marshalStruct)
}

func (c *Achievement) UnmarshalBSON(data []byte) error {
	var ach bsonAchievementStruct

	if err := bson.Unmarshal(data, &ach); err != nil {
		return err
	}

	*c = *NewAchievement(achievementEValue[ach.Achievement], ach.ReachedAt)
	return nil
}

type bsonAchievementStruct struct {
	Achievement string    `bson:"achievement"`
	ReachedAt   time.Time `bson:"reached_at"`
}

type Achievement struct {
	Achievement AchievementE
	ReachedAt   time.Time
}

func NewAchievement(ach AchievementE, time time.Time) *Achievement {
	return &Achievement{
		Achievement: ach,
		ReachedAt:   time,
	}
}
