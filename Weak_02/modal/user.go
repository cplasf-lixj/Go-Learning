package modal

import (
	"database/sql"
	"fmt"
)

func queryUserName(id int) (string, error) {
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id = ?", id).Scan(&username)
	if nil != err {
		if err == sql.ErrNoRows {
			fmt.Printf("SQL ERROR: No user with user_id " + id)
			return "", nil
		}
		return "", err
	}
	return username, nil
}
