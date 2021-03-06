package birthday

import (
	"fmt"
	"github.com/leekchan/timeutil"
	"time"
	"wumpus-birthday/pkg/globals"
)

const selectDate string = `
SELECT * FROM birthdays 
WHERE strftime('%d-%m', date) = $1 AND server_id = $2
ORDER BY date;
`

const selectDateAllServers string = `
SELECT * FROM birthdays 
WHERE strftime('%d-%m', date) = $1
ORDER BY date;
`

const selectAll string = `
SELECT * FROM birthdays
WHERE server_id = $1
ORDER BY date;
`

const insert string = `
INSERT OR REPLACE INTO birthdays(server_id, user_id, date) VALUES ($1, $2, $3);
`

const remove string = `
DELETE FROM birthdays WHERE server_id = $2 AND user_id = $1;
`

type Birthday struct {
	Date     time.Time `db:"date,required"`
	UserID   string    `db:"user_id,required"`
	ServerID string    `db:"server_id"`
}

func Today(serverID string) ([]Birthday, error) {
	today := time.Now()
	return Get(serverID, &today)
}

func TodayAllServers() ([]Birthday, error) {
	today := time.Now()
	return Get("", &today)
}

func Get(serverID string, day *time.Time) ([]Birthday, error) {
	var foundBirthdays []Birthday
	var err error
	formattedDate := timeutil.Strftime(day, "%d-%m")

	if serverID == "" {
		err = globals.DB.Select(&foundBirthdays, selectDateAllServers, formattedDate)
	} else {
		err = globals.DB.Select(&foundBirthdays, selectDate, formattedDate, serverID)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get birthdays: %s", err)
	}
	return foundBirthdays, nil
}

func GetAll(serverID string) ([]Birthday, error) {
	var foundBirthdays []Birthday
	if err := globals.DB.Select(&foundBirthdays, selectAll, serverID); err != nil {
		return nil, fmt.Errorf("failed to get birthdays: %s", err)
	}
	return foundBirthdays, nil
}

func Add(serverID string, userID string, date time.Time) error {
	_, err := globals.DB.Exec(insert, serverID, userID, date)
	return err
}

func Remove(serverID string, userID string) error {
	_, err := globals.DB.Exec(remove, serverID, userID)
	return err
}
