package structures

import "github.com/dgrijalva/jwt-go"

// User user object
type User struct {
	Name     string `json:"user_name"`
	Password string `json:"password"`
	ID       int
}

// UserAuthClaim User Auth Claim
type UserAuthClaim struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}

// DBConfig Config for DB
type DBConfig struct {
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Config Configuration
type Config struct {
	Host string   `json:"host"`
	DB   DBConfig `json:"db"`
}
