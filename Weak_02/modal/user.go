package modal

import (
	"database/sql"
	"fmt"
)

func queryUserName(id int) string {
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id = ?", id).Scan(&username)
	if nil != err {
		if err == sql.ErrNoRows {
			return ""
		}
		fmt.Printf("SQL ERROR: No user with user_id " + id)
	}
	return username
}
