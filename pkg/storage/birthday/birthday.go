package birthday

import (
	"fmt"
	"github.com/leekchan/timeutil"
	"time"
	"wumpus-birthday/pkg/globals"
)

const selectDate string = `
SELECT * FROM birthdays 
WHERE strftime('%d-%m', date) = $1;
`

const selectAll string = `
SELECT * FROM birthdays;
`

const insert string = `
INSERT INTO birthdays(user_id, date) VALUES ($1, $2);
`

type Birthday struct {
	Date   time.Time `db:"date,required"`
	UserID string    `db:"user_id,required"`
}

func Today() ([]Birthday, error) {
	today := time.Now()
	return Get(&today)
}

func Get(day *time.Time) ([]Birthday, error) {
	var foundBirthdays []Birthday
	if err := globals.DB.Select(
		&foundBirthdays, selectDate, timeutil.Strftime(&day, "%d-%m")); err != nil {

		return nil, fmt.Errorf("failed to get birthdays: %s", err)
	}
	return foundBirthdays, nil
}

func GetAll() ([]Birthday, error) {
	var foundBirthdays []Birthday
	if err := globals.DB.Select(&foundBirthdays, selectAll); err != nil {
		return nil, fmt.Errorf("failed to get birthdays: %s", err)
	}
	return foundBirthdays, nil
}

func Add(userID string, date time.Time) error {
	_, err := globals.DB.Exec(insert, userID, date)
	return err
}
