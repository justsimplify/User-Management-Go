package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"himanshu.com/sample/setup"
	"himanshu.com/sample/sqlhelper"
	st "himanshu.com/sample/structures"
	"himanshu.com/sample/users"
)

// UserRequest Object
type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWTToken Token Object
type JWTToken struct {
	Token string `json:"token"`
}

func getDBConnection(w http.ResponseWriter) *sql.DB {
	db, err := sqlhelper.GetDBConnection(w.Header().Get("dbname"), w.Header().Get("dbuser"), w.Header().Get("dbpassword"), w.Header().Get("dbhost"))
	if err != nil {
		panic(err)
	}
	return db
}

func setupDB(w http.ResponseWriter) string {
	db := getDBConnection(w)
	defer db.Close()
	return setup.DatabaseTables(db)
}

func createUser(user UserRequest, w http.ResponseWriter) (string, error) {
	db := getDBConnection(w)
	_, err := getUser(user, w)
	if err == nil {
		return "", errors.New("User already exists")
	}
	r := users.CreateNewUser(db, user.Username, user.Password)
	defer db.Close()
	return r, nil
}

func getUser(user UserRequest, w http.ResponseWriter) (st.User, error) {
	db := getDBConnection(w)
	u, err := users.GetExistingUser(db, user.Username, user.Password)
	defer db.Close()
	if err != nil {
		return st.User{}, err
	}
	return u, nil
}

func getUserToken(user UserRequest, w http.ResponseWriter) (string, error) {
	u, err := getUser(user, w)
	if err != nil {
		return "", err
	}
	token := users.CreateToken(u.Name, u.Password, u.ID)
	return token, nil
}

func verifyUserToken(token JWTToken) (*st.UserAuthClaim, error) {
	userClaim, err := users.VerifyToken(token.Token)
	if err != nil {
		return nil, err
	}
	return userClaim, nil
}

// LoginHandler Handle Login object
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var user UserRequest
	_ = json.NewDecoder(r.Body).Decode(&user)
	token, err := getUserToken(user, w)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

// VerifyTokenHandler Verify JWT Token Handler
func VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var token JWTToken
	_ = json.NewDecoder(r.Body).Decode(&token)
	claim, err := verifyUserToken(token)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		json.NewEncoder(w).Encode(map[string]*st.UserAuthClaim{"response": claim})
	}
}

// SetupHandler Setup Handler
func SetupHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": setupDB(w)})
}

// RegisterHandler User Register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var user UserRequest
	_ = json.NewDecoder(r.Body).Decode(&user)
	cu, err := createUser(user, w)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"response": cu})
	}
}
