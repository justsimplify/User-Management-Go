package users

import (
	"database/sql"

	st "himanshu.com/sample/structures"
)

// CreateNewUser Create New User
func CreateNewUser(db *sql.DB, userName string, password string) string {
	_, err := db.Query("INSERT INTO user (user_name, password) VALUES (?, ?)", userName, password)
	if err != nil {
		return err.Error()
	}
	return "Success"
}

// GetExistingUser Get Existing User
func GetExistingUser(db *sql.DB, userName string, password string) (st.User, error) {
	var u st.User
	err := db.QueryRow("SELECT * from user where user_name=? and password=?", userName, password).Scan(&u.ID, &u.Name, &u.Password)
	if err != nil {
		return st.User{}, err
	}
	return u, nil
}
