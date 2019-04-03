package storage

import (
	"fmt"
	"time"
	"wumpus-birthday/pkg/globals"
)

type Birthday struct {
	Date time.Time `db:"date,required"`
	UserID string `db:"user_id,required"`
}

func Get(day time.Time) ([]Birthday, error) {
	var foundBirthdays []Birthday
	if err := globals.DB.Select(&foundBirthdays, "SELECT * FROM birthdays"); err != nil {
		return nil, fmt.Errorf("failed to get birthdays: %s", err)
	}
	return foundBirthdays, nil
}
