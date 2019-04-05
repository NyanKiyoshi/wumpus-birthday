package storage

import (
	"fmt"
	"time"
	"wumpus-birthday/pkg/globals"
)

const selectBirthdays string = `
SELECT * FROM birthdays 
WHERE strftime('%d', date) = $1, strftime('%m', date) = $2;
`

type Birthday struct {
	Date   time.Time `db:"date,required"`
	UserID string    `db:"user_id,required"`
}

func Get(day time.Time) ([]Birthday, error) {
	var foundBirthdays []Birthday
	if err := globals.DB.Select(
		&foundBirthdays, selectBirthdays, day.Day(), day.Month()); err != nil {

		return nil, fmt.Errorf("failed to get birthdays: %s", err)
	}
	return foundBirthdays, nil
}
